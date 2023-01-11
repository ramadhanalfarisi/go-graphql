package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/ramadhanalfarisi/go-graphql-kocak/helpers"
)

type Product struct {
	ProductID       string         `json:"productID"`
	ProductName     string         `json:"productName" validate:"required"`
	ProductPrice    float64        `json:"productPrice" validate:"required"`
	ProductCategory string         `json:"productCategory" validate:"required"`
	CreatedAt       string         `json:"createdAt"`
	UpdatedAt       sql.NullString `json:"updatedAt"`
}

func (product *Product) GetProducts(db *sql.DB) []Product {
	var products []Product
	rows, err := db.Query("SELECT id, name, price, category, created_at, updated_at FROM products")
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
			UpdatedAt       sql.NullString
		)
		if err := rows.Scan(&ProductID,
			&ProductName,
			&ProductPrice,
			&ProductCategory,
			&CreatedAt,
			&UpdatedAt); err != nil {
			helpers.Error(err)
		}
		var productObject Product
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
	stmt, err := db.Prepare("INSERT INTO products( id, name, price, category, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, NULL);")
	if err != nil {
		helpers.Error(err)
		return err
	}
	defer stmt.Close()

	for _, prod := range products {
		prod_id := uuid.New()
		created_at := time.Now().Format("2006-01-02 15:04:05")
		if _, err := stmt.Exec(prod_id, prod.ProductName, prod.ProductPrice, prod.ProductCategory, created_at); err != nil {
			helpers.Error(err)
			return err
		}
	}
	return nil
}
