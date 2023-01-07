package routers

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type Router struct {
	Router *mux.Router
	RouterGraph *mux.Router
	DB *sql.DB
}