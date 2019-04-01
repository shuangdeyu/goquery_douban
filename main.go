package main

import (
	"fmt"
	"goquery_douban/server"
	"helper_go/comhelper"
	"time"
)

func main() {

	//server.ReadFile("data/movie.json")
	//server.WriteFile("data/movie.json", "[1,2,3]")
	//server.AppendFile("data/movie.json", ",[1,2,3]")

	url := ""
	for i := 25; i < 26; i++ {
		start := i * comhelper.StringToInt(server.PAGELIMIT)
		url = server.BASEURL + server.TAG + server.SORT + "&page_limit=" + server.PAGELIMIT + "&page_start=" + comhelper.IntToString(start)
		err := server.GetMovieList(url)
		if err != nil {
			break
		}
		fmt.Println(url)
		time.Sleep(1 * time.Second)
	}
	fmt.Println(url)
}
