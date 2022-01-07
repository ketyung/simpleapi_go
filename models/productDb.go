package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

const dbuser = "sample_erp_usr"
const dbpass = "128272bnH627Ya3a"
const dbname = "erp_sample_db"

func GetProducts() []Product {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

	// if there is an error opening the connection, handle it
	if err != nil {

		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM product")

	if err != nil {

		fmt.Println("Err", err.Error())

		return nil

	}

	products := []Product{}

	for results.Next() {

		var prod Product

		err = results.Scan(&prod.Code, &prod.Name, &prod.Qty, &prod.LastUpdated)

		if err != nil {
			panic(err.Error())
		}

		products = append(products, prod)

		//fmt.Println("product.code :", prod.Code+" : "+prod.Name)
	}

	return products

}

func GetProduct(code string) *Product {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

	prod := &Product{}

	if err != nil {

		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM product where code=?", code)

	if err != nil {

		fmt.Println("Err", err.Error())

		return nil
	}

	if results.Next() {

		err = results.Scan(&prod.Code, &prod.Name, &prod.Qty, &prod.LastUpdated)

		if err != nil {
			return nil
		}
	} else {

		return nil
	}

	return prod

}

func AddProduct(product Product) {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

	if err != nil {
		panic(err.Error())
	}

	// defer the close till after this function has finished
	// executing
	defer db.Close()

	insert, err := db.Query(
		"INSERT INTO product (code,name,qty,last_updated) VALUES (?,?,?, now())",
		product.Code, product.Name, product.Qty)

	/*
		// Or use fmt.Sprintf to concatenate SQL statement if prepared statement isn't worth here

		sqlstm :=
			fmt.Sprintf("INSERT INTO product (code,name,qty,last_updated)"+
				" VALUES ('%s','%s',%d, now())",
				product.Code, product.Name, product.Qty)

		insert, err := db.Query(sqlstm)
	*/

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}
