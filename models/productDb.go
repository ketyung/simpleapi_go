package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetProducts() []Product {

	dbuser := "sample_erp_usr"
	dbpass := "128272bnH627Ya3a"
	dbname := "erp_sample_db"

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM product LIMIT 0,10")

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
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
