package main

import (
	"log"
	"net/http"
	"project-pertama/config"
	categorycontroller "project-pertama/controllers/categorycontrollers"
	homecontroller "project-pertama/controllers/homexontroller"
	"project-pertama/controllers/productcontroller"
)

func main() {
	config.ConnectDB()

	// Routes
	// 1.Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// 2. Category
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	// 3. Products
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("Server running in port 8000")
	http.ListenAndServe(":8000", nil)
}
