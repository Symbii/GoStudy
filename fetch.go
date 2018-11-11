package main

import (
	"os"
	"fmt"
	"net/http"
	"time"
	"io"
	"strings"
)

func main(){
	ch := make(chan string)


	for _, url := range os.Args[1:] {
		b := strings.HasPrefix(url, "http://")
		if b != true{
			url = "http://" + url
		}
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
}

func fetch(url string, ch  chan<- string){
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch<- fmt.Sprintf("while getting %s, %v ", url , err)
		return
	}
	nbytes, err:= io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s,%v", url, err)
		return
	}
	cost := time.Since(start).Seconds()
	ch<- fmt.Sprintf("Time: %v, \t%v, \t%s",cost,nbytes, url)
	return
}
