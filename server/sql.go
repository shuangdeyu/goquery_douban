package server

import (
	"goquery_douban/model"
)

type MovieParam struct {
	Film_id       int
	Name          string
	Tags          string
	Img           string
	Year          string
	Rating_num    float64
	Rating_people int
	Bigstar       int
	Stars_five    float64
	Stars_four    float64
	Stars_three   float64
	Stars_two     float64
	Stars_one     float64
	Summary       string
	Director      string
	Actor         string
	Country       string
	Release_date  string
	Film_length   int
}

/**
 * 通过film_id获取
 */
func GetMovieByFilmId(film_id int) map[string]string {
	info, _ := model.DefaultMovie.QueryByMap(model.Arr{"film_id": film_id})
	if len(info) > 0 {
		return info[0]
	} else {
		return nil
	}
}

/**
 * 插入
 */
func InsertFilm(param *MovieParam) error {
	// 插入用户
	insert := &model.Movie{
		Film_id:       param.Film_id,
		Name:          param.Name,
		Tags:          param.Tags,
		Img:           param.Img,
		Year:          param.Year,
		Rating_num:    param.Rating_num,
		Rating_people: param.Rating_people,
		Bigstar:       param.Bigstar,
		Stars_five:    param.Stars_five,
		Stars_four:    param.Stars_four,
		Stars_three:   param.Stars_three,
		Stars_two:     param.Stars_two,
		Stars_one:     param.Stars_one,
		Summary:       param.Summary,
		Director:      param.Director,
		Actor:         param.Actor,
		Country:       param.Country,
		Release_date:  param.Release_date,
		Film_length:   param.Film_length,
	}
	err := insert.InsertByStructure()
	return err
}

/**
 * 更新
 */
func UpdateFilm(param *MovieParam) error {
	err := model.DefaultMovie.Update(model.Arr{
		"name":          param.Name,
		"tags":          param.Tags,
		"img":           param.Img,
		"year":          param.Year,
		"rating_num":    param.Rating_num,
		"rating_people": param.Rating_people,
		"bigstar":       param.Bigstar,
		"stars_five":    param.Stars_five,
		"stars_four":    param.Stars_four,
		"stars_three":   param.Stars_three,
		"stars_two":     param.Stars_two,
		"stars_one":     param.Stars_one,
		"summary":       param.Summary,
		"director":      param.Director,
		"actor":         param.Actor,
		"country":       param.Country,
		"release_date":  param.Release_date,
		"film_length":   param.Film_length,
	}, model.Arr{"film_id": param.Film_id})
	return err
}
