package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ramadhanalfarisi/go-graphql-kocak/helpers"
	"github.com/ramadhanalfarisi/go-graphql-kocak/models"
)

func (c *Controller) InsertProducts(w http.ResponseWriter, r *http.Request) {
	db := c.DB
	var response map[string]interface{}
	var products []models.Product
	product := models.Product{}
	err := json.NewDecoder(r.Body).Decode(&products)
	if err != nil {
		helpers.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		response = helpers.FailedResponse(http.StatusInternalServerError, err.Error())
	} else {
		var errors []errorValdiate
		for i := range products {
			errors = validateModel(products[i], i, &errors)
		}
		if errors != nil {
			err := product.InsertProducts(db, products)
			if err != nil {
				helpers.Error(err)
				w.WriteHeader(http.StatusInternalServerError)
				response = helpers.FailedResponse(http.StatusInternalServerError, err.Error())
			} else {
				w.WriteHeader(http.StatusCreated)
				response = helpers.SuccessResponse(http.StatusCreated, "Insert Product Successfully", nil, nil)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			response = helpers.InvalidResponse(http.StatusBadRequest, errors)
		}
	}

	json, _ := json.Marshal(response)
	w.Write(json)
}
