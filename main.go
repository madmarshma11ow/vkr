package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/skip2/go-qrcode"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func main() {

	err := qrcode.WriteFile("https://youtu.be/dQw4w9WgXcQ?t=43", qrcode.Medium, 256, "assets/qr.png")
	if err != nil {
		fmt.Printf("Sorry couldn't create qrcode:,%v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)

}
