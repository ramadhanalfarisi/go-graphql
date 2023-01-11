package schemas

import (

	"github.com/graphql-go/graphql"
	"github.com/ramadhanalfarisi/go-graphql-kocak/models"
)


func (s *Schema) GetProduct() graphql.Schema {

	var productObject = graphql.NewObject(graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"productID": &graphql.Field{
				Type: graphql.String,
			},
			"productName": &graphql.Field{
				Type: graphql.String,
			},
			"productPrice": &graphql.Field{
				Type: graphql.Float,
			},
			"productCategory": &graphql.Field{
				Type: graphql.String,
			},
			"createdAt": &graphql.Field{
				Type: graphql.String,
			},
			"updatedAt": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	var rootProduct = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"products": &graphql.Field{
				Type:        graphql.NewList(productObject),
				Description: "List of products",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					db := s.DB
					product := models.Product{}
					products := product.GetProducts(db)
					return products, nil
				},
			},
		},
	})

	var ProductSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: rootProduct,
	})

	return ProductSchema
}
