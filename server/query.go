package server

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/kataras/iris/core/errors"
	"helper_go/comhelper"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	// https://movie.douban.com/explore#!type=movie&tag=热门&sort=recommend&page_limit=20&page_start=0
	BASEURL = "https://movie.douban.com/j/search_subjects?type=movie"
	//TAG        = "&tag=热门"
	TAG  = "&tag=经典"
	SORT = "&sort=recommend" // 按热度
	//SORT       = "&sort=rank" // 按评价
	PAGELIMIT  = "20"
	CONTENTURL = "https://movie.douban.com/subject/"
)

/**
 * 获取电影列表
 */
func GetMovieList(url string) error {
	// 获取列表
	ret := HttpGet(url)
	if ret != "" {
		var data map[string]interface{}
		json.Unmarshal([]byte(ret), &data)
		if subjects, ok := data["subjects"].([]interface{}); ok && len(subjects) > 0 {
			for _, v := range subjects {
				d := v.(map[string]interface{})
				// 获取详情
				info_url := d["url"].(string)
				param := GetMovieInfo(info_url)
				param.Film_id = comhelper.StringToInt(d["id"].(string))
				param.Name = d["title"].(string)
				param.Rating_num = comhelper.StringToFloat(d["rate"].(string), 64)

				// 插入库
				info := GetMovieByFilmId(param.Film_id)
				if info == nil {
					InsertFilm(param)
				} else {
					UpdateFilm(param)
				}
				time.Sleep(2 * time.Second)
			}
		}
		return nil
	} else {
		return errors.New("Http Get is error, empty content")
	}
}

/**
 * 获取详情
 */
func GetMovieInfo(url string) *MovieParam {
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("status code error: %d %s", res.StatusCode, res.Status)
		return nil
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(err)
		return nil
	}

	param := MovieParam{}
	doc.Find("#content").Each(func(i int, s *goquery.Selection) {
		param.Year = s.Find("h1 .year").Text()
		param.Img, _ = s.Find("#mainpic img").Attr("src")
		param.Summary, _ = s.Find("#link-report span[property]").Html()
		param.Rating_people = comhelper.StringToInt(s.Find(".rating_people span[property]").Text())
		star, _ := s.Find(".bigstar").Attr("class")
		param.Bigstar = comhelper.StringToInt(star[len(star)-2 : len(star)])
		stars_five := s.Find(".stars5+div+span").Text()
		param.Stars_five = comhelper.StringToFloat(stars_five[0:len(stars_five)-1], 64)
		stars_four := s.Find(".stars4+div+span").Text()
		param.Stars_four = comhelper.StringToFloat(stars_four[0:len(stars_four)-1], 64)
		stars_three := s.Find(".stars3+div+span").Text()
		param.Stars_three = comhelper.StringToFloat(stars_three[0:len(stars_three)-1], 64)
		stars_two := s.Find(".stars2+div+span").Text()
		param.Stars_two = comhelper.StringToFloat(stars_two[0:len(stars_two)-1], 64)
		stars_one := s.Find(".stars1+div+span").Text()
		param.Stars_one = comhelper.StringToFloat(stars_one[0:len(stars_one)-1], 64)

		s.Find("#info").Each(func(ii int, ss *goquery.Selection) {
			info, _ := ss.Html()
			param.Director = ss.Find("a[rel*=directedBy]").Text()
			film_length, _ := ss.Find("span[property*=runtime]").Attr("content")
			param.Film_length = comhelper.StringToInt(film_length)
			param.Release_date = ss.Find("span[property*=initialReleaseDate]").Text()

			tags := ""
			ss.Find("span[property*=genre]").Each(func(i int, s *goquery.Selection) {
				if tags == "" {
					tags += s.Text()
				} else {
					tags += "/" + s.Text()
				}
			})
			param.Tags = tags

			actor := ""
			ss.Find("a[rel*=starring]").Each(func(i int, s *goquery.Selection) {
				if actor == "" {
					actor += s.Text()
				} else {
					actor += "/" + s.Text()
				}
			})
			param.Actor = actor

			c_start := strings.Index(info, "<span class=\"pl\">制片国家/地区:</span>")
			c_end := strings.Index(info, "<span class=\"pl\">语言")
			param.Country = comhelper.TrimHtml(info[c_start+44 : c_end])
		})
	})

	return &param
}
