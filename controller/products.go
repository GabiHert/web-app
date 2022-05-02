package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"web-app/model"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := model.GetProducts()
	temp.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)

}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		amount := r.FormValue("amount")
		description := r.FormValue("description")
		price := r.FormValue("price")

		priceConverted, err := strconv.ParseFloat(price, 64)

		if err != nil {
			log.Println("Error converting price to float")
		}

		amountConverted, err := strconv.Atoi(amount)

		if err != nil {
			log.Println("Error converting amount to int")
		}

		model.NewProduct(name, description, amountConverted, priceConverted)

	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	model.DeleteProduct(productId)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	product := model.GetProduct(productId)

	temp.ExecuteTemplate(w, "Edit", product)
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.PostFormValue("id"))
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		amount := r.FormValue("amount")
		price := r.FormValue("price")

		idParsed, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error parsing id to int")
		}

		amountParsed, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Error parsing amount to int")
		}

		priceParsed, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error parsing price to float")
		}

		model.UpdateProduct(idParsed, name, description, priceParsed, amountParsed)
		http.Redirect(w, r, "/", 301)

	}
}
