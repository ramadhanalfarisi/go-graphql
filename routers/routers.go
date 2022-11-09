package routers

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type Router struct {
	Router *mux.Router
	DB *sql.DB
}