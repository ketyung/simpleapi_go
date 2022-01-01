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
			panic(err.Error()) // proper error handling instead of panic in your app
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
