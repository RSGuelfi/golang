package main

import "fmt"

func main() {
	x := struct {
		nome      string
		sabor     string
		ondetem   []string
		vaibemcom map[string]string
	}{nome: "Stroopwafel",
		sabor:   "doce",
		ondetem: []string{"na Holanda", "na casa do seu tio rico"},
		vaibemcom: map[string]string{
			"Café da manha": "chá,",
			"Almoço":        "cafezinho,",
			"Janta":         "não vai bem",
		}}
	fmt.Println(x)

}
