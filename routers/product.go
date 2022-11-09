package routers

import (
	"github.com/graphql-go/handler"
	"github.com/ramadhanalfarisi/go-graphql-kocak/controllers"
	"github.com/ramadhanalfarisi/go-graphql-kocak/schemas"
)

func (router *Router) ProdRouters() {
	controller := controllers.Controller{}
	schema := schemas.Schema{}
	controller.DB = router.DB
	schema.DB = router.DB

	getProducts := schema.GetProduct()
	handleProducts := handler.New(&handler.Config{
		Schema:   &getProducts,
		Pretty:   true,
		GraphiQL: false,
	})

	router.Router.Handle("/products",handleProducts).Methods("POST")
	router.Router.HandleFunc("/products", controller.InsertProducts).Methods("POST")
}
