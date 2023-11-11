package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	f, err := os.OpenFile("url.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch, w)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	f.WriteString(fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds()))
	w.Flush()
}

func fetch(url string, ch chan<- string, w *bufio.Writer) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
	w.WriteString(fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url))
}
