package controllers

import (
	"database/sql"

	"github.com/ramadhanalfarisi/go-graphql-kocak/helpers"
)

type Controller struct {
	DB *sql.DB
}

type errorValdiate struct {
	Index   int
	Message []string
}

func validateModel(model interface{},i int, messages *[]errorValdiate) []errorValdiate {
	var errors []errorValdiate
	validate, valid := helpers.Validate(model)
	if !valid {
		err := errorValdiate{}
		err.Index = i
		err.Message = validate
		errors = append(errors, err)
	}
	*messages = append(*messages, errors...)
	return *messages
}
