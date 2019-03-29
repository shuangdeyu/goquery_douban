package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

/**
 * 读取文件
 */
func ReadFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

/**
 * 写文件
 */
func WriteFile(path, content string) {
	// 可写方式打开文件
	file, err := os.OpenFile(
		path,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// 写字节到文件中
	byteSlice := []byte(content)
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)
}

/**
 * 追加写入文件
 */
func AppendFile(path, content string) {
	// 可写方式打开文件
	file, err := os.OpenFile(
		path,
		os.O_WRONLY|os.O_APPEND|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// 写字节到文件中
	byteSlice := []byte(content)
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)
}

/**
 * 获取get请求
 */
func HttpGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("httpGet url : %v , error : %v\n", url, err)
		return ""
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("httpGet url : %v , error : %v\n", url, err)
		}
	}()
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("httpGet url : %v , error : %v\n", url, err)
		return ""
	}
	return string(body)
}
