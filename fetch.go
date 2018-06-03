package main

import (
	"os"
	"net/http"
	"fmt"
//	"io/ioutil"
	"io"
	"strings"
)

func main(){
	for _, url := range os.Args[1:] {
		b := strings.HasPrefix(url, "https://")
		if b != true{
			url = "https://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil{
			fmt.Fprintf(os.Stderr, "fetch:  reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", os.Stdout)
	}
}

