package main

import (
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
	"time"
	"io"
)

func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		//fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		//os.Exit(1)
		ch <- fmt.Sprint(err) // 发送到通道ch
		return
	}

	//b, err:= ioutil.ReadAll(resp.Body)

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	resp.Body.Close()

	if err != nil {
		//fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		//os.Exit(1)

		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return

	}

	secs := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func main() {

	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // 启动一个goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // 从通道ch接收
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}
