package schemas

import (
	"github.com/graphql-go/graphql"
	"github.com/ramadhanalfarisi/go-graphql-kocak/helpers"
	"github.com/ramadhanalfarisi/go-graphql-kocak/models"
)

type Get struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}

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
					page, isOK := p.Args["page"].(int)
					limit, isOK2 := p.Args["limit"].(int)
					if !isOK {
						page = 1
					}
					if !isOK2 {
						limit = 10
					}
					db := s.DB
					product := models.Product{}

					param, err := helpers.GetMetaParam(page, limit)
					if err != nil {
						helpers.Error(err)
					}
					pagination := helpers.Pagination{}
					get_pagin := pagination.CreatePagination(param)
					products := product.GetProducts(db)
					get := Get{}
					get.Data = products
					get.Meta = get_pagin
					return get, nil
				},
			},
		},
	})

	var ProductSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: rootProduct,
	})

	return ProductSchema
}
