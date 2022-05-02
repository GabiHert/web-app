package model

import (
	"fmt"
	"web-app/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

func GetProducts() []Product {
	db := db.DbConnection()
	fmt.Println("QUERY STARTED")

	selectAllProducts, err := db.Query("select * from products order by id asc")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("QUERY OK")

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next() {

		var id, amount int
		var name_, description string
		var price float64

		err = selectAllProducts.Scan(&id, &name_, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name_
		p.Description = description
		p.Price = price
		p.Amount = amount

		products = append(products, p)
	}

	defer db.Close()
	return products

}

func NewProduct(name string, description string, amount int, price float64) {
	db := db.DbConnection()

	insertData, err := db.Prepare("insert into products (name_, description, price, amount) values ($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}

	insertData.Exec(name, description, amount, price)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.DbConnection()

	deleteProduct, err := db.Prepare("delete from products where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
	defer db.Close()
}

func GetProduct(id string) Product {
	db := db.DbConnection()

	product, err := db.Query("select * from products where id = $1", id)

	if err != nil {
		panic(err.Error())
	}

	resolvedProduct := Product{}

	for product.Next() {

		var id, amount int
		var name_, description string
		var price float64

		err = product.Scan(&id, &name_, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}
		resolvedProduct.Id = id
		resolvedProduct.Amount = amount
		resolvedProduct.Name = name_
		resolvedProduct.Price = price
		resolvedProduct.Description = description

	}
	defer db.Close()
	return resolvedProduct
}

func UpdateProduct(id int, name string, description string, price float64, amount int) {
	db := db.DbConnection()

	updatedProduct, err := db.Prepare("update products set name_ = $1,description = $2,price = $3,amount = $4 where id = $5")

	if err != nil {
		panic(err.Error())
	}
	updatedProduct.Exec(name, description, price, amount, id)
	defer db.Close()
}
