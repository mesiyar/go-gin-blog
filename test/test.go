package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main()  {
	//httptest()
	createFile()
}
func httptest()  {
	defer fmt.Println("httptest")
	resp, _ := http.Get("http://www.baidu.com")
	defer resp.Body.Close()
	body ,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func createFile() {
	defer fmt.Println("end")
	fmt.Println("start")
	data := []byte("this is a file")
	err := ioutil.WriteFile("aa.txt", data, 0664)
	if err != nil {
		log.Fatal(err)
	}
}


