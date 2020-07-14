package main

import (
	"net/http"
)

func main() {
	//Rutas
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)
	//Inicia el servidor
	http.ListenAndServe(":3000", nil)
}

//w es de escribir "write"en ingles, r es de leer "read"
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pagina de contacto"))
}
