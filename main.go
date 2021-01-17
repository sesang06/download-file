package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func DownloadFile(urls ...string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)

		for _, url := range urls {
			func(url string) {

				tokens := strings.Split(url, "/")
				fileName := tokens[len(tokens)-1]
				fmt.Println("downloading", url, "to", fileName)

				output, err := os.Create(fileName)

				if err != nil {
					log.Fatal("Error while createing", fileName, "-", err)
					return
				}
				defer output.Close()

				res, err := http.Get(url)
				if err != nil {
					log.Fatal("http get error", err)
					return
				}
				defer res.Body.Close()
				_, err = io.Copy(output, res.Body)
				if err != nil {
					log.Fatal("Error while downloading", url, "-", err)
					return
				}
				fmt.Println("Downloaded", fileName)
				out <- fileName
			}(url)
		}
	}()
	return out
}

func main() {

	for file := range DownloadFile("https://i.picsum.photos/id/862/200/300.jpg") {
		fmt.Println("file", file)
	}

}

