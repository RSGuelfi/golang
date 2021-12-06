package main

import (
	"curso-go-jeff/avancado/servidor_web/manipulador"
	"fmt"
	"net/http"
)

func Funcao(w http.ResponseWriter, r *http.Request) {

}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá Mundo!")
	})
	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	fmt.Println("Server is running...")
	http.ListenAndServe(":7777", nil)
}
