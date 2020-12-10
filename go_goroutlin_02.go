package main

import(
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"time"
)

func responseSize(url string){
	fmt.Println(url)
	resp, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer resp.Body.Close()
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(len(body))
}

func main(){
	go responseSize("https://www.inhatc.ac.kr")
	go responseSize("https://www.nate.com")
	go responseSize("https://www.google.com")
	go responseSize("https://www.harvard.edu")
	time.Sleep(time.Second * 5)

	/*
	resp, e := http.Get("https://www.inhatc.ac.kr")
	if e != nil {  //정상인경우 nil 값을 받기 때문에 아닐 경우 에러처리 
		log.Fatal(e)
	}
	defer resp.Body.Close() //main 함수 끝나는 시점에 에러 여부 상관 없이 연결 해제 
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		log.Fatal(e)
	}
	
	fmt.Println(string (body)) //string 안할시 아스키 코드로 출력 
	*/
}
