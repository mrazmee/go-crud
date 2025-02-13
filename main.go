package main

import (
	"net/http"
	"fmt"
	"github.com/mrazmee/go-crud/controllers/pasiencontroller"
)

func main() {
	http.HandleFunc("/", pasiencontroller.Index)
	http.HandleFunc("/pasien", pasiencontroller.Index)
	http.HandleFunc("/pasien/index", pasiencontroller.Index)
	http.HandleFunc("/pasien/add", pasiencontroller.Add)
	http.HandleFunc("/pasien/edit", pasiencontroller.Edit)
	http.HandleFunc("/pasien/delete", pasiencontroller.Delete)
	
	fmt.Println("Server started on: http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}