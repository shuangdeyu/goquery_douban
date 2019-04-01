package model

import (
	"fmt"
	"helper_go/comhelper"
	"log"
)

type Movie struct {
	Id            int     `xorm:"int(11) NOT NULL autoincr" json:"id"`
	Film_id       int     `xorm:"int(11) NOT NULL" json:"film_id"`
	Name          string  `xorm:"varchar(50) NOT NULL" json:"name"`
	Tags          string  `xorm:"varchar(50) NOT NULL" json:"tags"`
	Img           string  `xorm:"text NOT NULL" json:"img"`
	Year          string  `xorm:"varchar(10) NOT NULL" json:"year"`
	Rating_num    float64 `xorm:"float(16,1) NOT NULL" json:"rating_num"`
	Rating_people int     `xorm:"int(10) NOT NULL" json:"rating_people"`
	Bigstar       int     `xorm:"int(5) NOT NULL" json:"bigstar"`
	Stars_five    float64 `xorm:"float(16,1) NOT NULL" json:"stars_five"`
	Stars_four    float64 `xorm:"float(16,1) NOT NULL" json:"stars_four"`
	Stars_three   float64 `xorm:"float(16,1) NOT NULL" json:"stars_three"`
	Stars_two     float64 `xorm:"float(16,1) NOT NULL" json:"stars_two"`
	Stars_one     float64 `xorm:"float(16,1) NOT NULL" json:"stars_one"`
	Summary       string  `xorm:"text NOT NULL" json:"summary"`
	Director      string  `xorm:"varchar(50) NOT NULL" json:"director"`
	Actor         string  `xorm:"varchar(50) NOT NULL" json:"actor"`
	Country       string  `xorm:"varchar(50) NOT NULL" json:"country"`
	Release_date  string  `xorm:"varchar(50) NOT NULL" json:"release_date"`
	Film_length   int     `xorm:"int(5) NOT NULL" json:"film_length"`
}

var DefaultMovie = &Movie{}

/**
 * 执行原生sql查询，返回string类型的map
 */
func (m *Movie) Query(args ...interface{}) ([]map[string]string, error) {
	// 基础sql语句
	sql := ""
	switch val := args[0].(type) {
	case string:
		sql = val
	}
	// 映射参数
	params := []interface{}{}
	if len(args) > 1 {
		switch val := args[1].(type) {
		case []interface{}:
			params = val
		}
	}
	// order 语句拼接
	if len(args) > 2 {
		switch val := args[2].(type) {
		case string:
			sql += " " + val
		}
	}
	// limit 语句拼接
	if len(args) > 3 {
		switch val := args[3].(type) {
		case []int:
			sql += " limit " + comhelper.IntToString(val[0]) + "," + comhelper.IntToString(val[1])
		}
	}

	//ret, err := DbInit().SQL(sql, params...).QueryString()
	ret, err := DbInit().QueryString(sql, params...)
	if err != nil {
		log.Print(err.Error())
	}
	return ret, nil
}

/**
 * 执行原生sql，返回定义的结构体类型
 */
func (m *Movie) QueryStructure(args ...interface{}) ([]Movie, error) {
	// 基础sql语句
	sql := ""
	switch val := args[0].(type) {
	case string:
		sql = val
	}
	// 映射参数
	params := []interface{}{}
	if len(args) > 1 {
		switch val := args[1].(type) {
		case []interface{}:
			params = val
		}
	}
	// order 语句拼接
	if len(args) > 2 {
		switch val := args[2].(type) {
		case string:
			sql += " " + val
		}
	}
	// limit 语句拼接
	if len(args) > 3 {
		switch val := args[3].(type) {
		case []int:
			sql += " limit " + comhelper.IntToString(val[0]) + "," + comhelper.IntToString(val[1])
		}
	}

	result := []Movie{}
	err := DbInit().SQL(sql, params...).Find(&result)
	if err != nil {
		log.Print(err.Error())
	}
	return result, nil
}

/**
 * 通过参数构造sql查询，返回string类型的map
 */
func (m *Movie) QueryByMap(args ...interface{}) ([]map[string]string, error) {
	sql := "select * from q_movie where 1=1 "
	// 拼接where语句
	var params []interface{}
	switch val := args[0].(type) {
	case Arr:
		for k, v := range val {
			switch v_type := v.(type) {
			case string:
				if v_type == "" {
					continue
				}
			}
			sql += fmt.Sprintf("and q_movie.%s = ? ", k)
			params = append(params, v)
		}
	}
	// 拼接order语句
	if len(args) > 1 {
		switch val := args[1].(type) {
		case string:
			if val != "" {
				sql += val + " "
			}
		}
	}
	// 拼接limit语句
	if len(args) > 2 {
		switch val := args[2].(type) {
		case []int:
			sql += "limit " + comhelper.IntToString(val[0]) + "," + comhelper.IntToString(val[1])
		}
	}

	//ret, err := DbInit().SQL(sql, params...).QueryString()
	ret, err := DbInit().QueryString(sql, params...)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return ret, nil
}

/**
 * 通过参数构造sql查询，返回定义的结构体类型
 */
func (m *Movie) QueryStructureByMap(args ...interface{}) ([]Movie, error) {
	result := []Movie{}

	sql := "select * from q_movie where 1=1 "
	// 拼接where语句
	var params []interface{}
	switch val := args[0].(type) {
	case Arr:
		for k, v := range val {
			switch v_type := v.(type) {
			case string:
				if v_type == "" {
					continue
				}
			}
			sql += fmt.Sprintf("and q_movie.%s = ? ", k)
			params = append(params, v)
		}
	}
	// 拼接order语句
	if len(args) > 1 {
		switch val := args[1].(type) {
		case string:
			if val != "" {
				sql += val + " "
			}
		}
	}
	// 拼接limit语句
	if len(args) > 2 {
		switch val := args[2].(type) {
		case []int:
			sql += "limit " + comhelper.IntToString(val[0]) + "," + comhelper.IntToString(val[1])
		}
	}

	err := DbInit().SQL(sql, params...).Find(&result)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return result, nil
}

/**
 * 获取count
 */
func (m *Movie) Count() (int64, error) {
	ret, err := DbInit().Count(m)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return ret, nil
}

/**
 * 删除，通过参数构造sql
 */
func (m *Movie) Delete(args Arr) error {
	sql := "delete from q_movie where 1=1 "
	var params []interface{}
	for k, v := range args {
		sql += fmt.Sprintf("and q_movie.%s = ? ", k)
		params = append(params, v)
	}
	//_, err := DbInit().SQL(sql, params...).QueryString()
	_, err := DbInit().QueryString(sql, params...)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

/**
 * 删除，绑定结构体
 */
func (m *Movie) DeleteByStructure(id int) error {
	_, err := DbInit().Id(id).Delete(m)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

/**
 * 更新，通过参数构造sql
 */
func (m *Movie) Update(set, args Arr) error {
	sql := "update q_movie set "
	var params []interface{}
	for k, v := range set {
		sql += fmt.Sprintf("q_movie.%s = ?,", k)
		params = append(params, v)
	}

	sql = sql[:len(sql)-1]
	sql += " where 1=1 "
	for k, v := range args {
		sql += fmt.Sprintf("and q_movie.%s = ? ", k)
		params = append(params, v)
	}

	//_, err := DbInit().SQL(sql, params...).QueryString()
	_, err := DbInit().QueryString(sql, params...)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

/**
 * 更新，绑定结构体
 */
func (m *Movie) UpdateByStructure(args *Movie) error {
	_, err := DbInit().Update(m, args)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

/**
 * 新增，绑定结构体
 */
func (m *Movie) InsertByStructure(args ...string) error {
	_, err := DbInit().Omit(args...).Insert(m)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
