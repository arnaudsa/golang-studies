package main

import (
	"golang-studies/alura/routes"
	"net/http"
)

func main() {
	routes.CarregarRotas()
	http.ListenAndServe(":8000", nil)
}



