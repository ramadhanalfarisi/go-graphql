package testing

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/ramadhanalfarisi/go-graphql-kocak/app"
)

var a app.App

func TestMain(m *testing.M) {
	a.ConnectDB()
	a.Routes()

	code := m.Run()
	clearTable()
	os.Exit(code)
}

func clearTable() {
	a.DB.Exec("DELETE FROM products;")
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func insertProducts(i int) {
	for j := 0; j < i; j++ {
		prod_id := uuid.New()
		prod_name := "product_" + strconv.Itoa(j)
		prod_price := j
		prod_cat := "category_" + strconv.Itoa(j)
		created_at := time.Now().Format("2006-01-02 15:04:05")
		a.DB.Exec("INSERT INTO products VALUES(?, ?, ?, ?, ?, NULL);", prod_id, prod_name, prod_price, prod_cat, created_at)
	}
}

func TestGetProducts(t *testing.T) {
	clearTable()
	insertProducts(5)
	data := []byte(`query {
		products {
			productID
			productName
		}
	}`)
	req, _ := http.NewRequest("POST", "/graph/products", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}
