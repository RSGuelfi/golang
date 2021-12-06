package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Patterns

// <-chan - canal somente-leitura
// Titulo obtem o título de uma janela html
func Titulo(urls ...string) <-chan string {
	c := make(chan string)

	for _, url := range urls {
		go func(cUrl string) {
			resp, _ := http.Get(cUrl)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			aRetorno := r.FindStringSubmatch(string(html))

			// teste para evitar erro
			if cap(aRetorno) == 0 {
				c <- "Erro ao ler página " + cUrl
				return
			}

			c <- aRetorno[1]
		}(url)
	}

	return c
}

func main() {
	t1 := Titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := Titulo("https://www.amazon.com", "https://www.youtube.com")
	fmt.Println("Primeiros:", <-t1, "|", <-t2)
	fmt.Println("Segundos:", <-t1, "|", <-t2)
}
