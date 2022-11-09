package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/ramadhanalfarisi/go-graphql-kocak/helpers"
)

type Product struct {
	ProductID       string `json:"productID"`
	ProductName     string `json:"productName" validate:"required"`
	ProductPrice    float64 `json:"productPrice" validate:"required"`
	ProductCategory string `json:"productCategory" validate:"required"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
}

func (product *Product) GetProducts(db *sql.DB) []Product {
	var products []Product
	var productObject Product
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		helpers.Error(err)
	}

	for rows.Next() {
		var (
			ProductID       string
			ProductName     string
			ProductPrice    float64
			ProductCategory string
			CreatedAt       string
			UpdatedAt       string
		)
		if err := rows.Scan(&ProductID,
			&ProductName,
			&ProductPrice,
			&ProductCategory,
			&CreatedAt,
			&UpdatedAt); err != nil {
			helpers.Error(err)
		}
		productObject.ProductID = ProductID
		productObject.ProductName = ProductName
		productObject.ProductPrice = ProductPrice
		productObject.ProductCategory = ProductCategory
		productObject.CreatedAt = CreatedAt
		productObject.UpdatedAt = UpdatedAt
		products = append(products, productObject)
	}
	return products
}

func (product *Product) InsertProducts(db *sql.DB, products []Product) error {
	stmt, err := db.Prepare("INSERT INTO products(product_id,product_name,product_price,product_category,created_at,updated_at) VALUES( ?, ?, ?, ?, ?, NULL)")
	if err != nil {
		helpers.Error(err)
		return err
	}
	defer stmt.Close()

	for _, prod := range products {
		prod_id := uuid.New()
		created_at := time.Now().Format("2006-01-02 15:04:05")
		if _, err := stmt.Exec(prod_id,prod.ProductName,prod.ProductPrice,prod.ProductCategory,created_at); err != nil {
			helpers.Error(err)
			return err
		}
	}
	return nil
}
