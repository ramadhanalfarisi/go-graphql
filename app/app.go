package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/ramadhanalfarisi/go-graphql-kocak/helpers"
	"github.com/ramadhanalfarisi/go-graphql-kocak/routers"
)

var host, uname, password, port, dbname string

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
	if env := os.Getenv("ENVIRONMMENT"); env == "production" {
		port = os.Getenv("PORT_PRODDUCTION")
		host = os.Getenv("HOST_PRODDUCTION")
		uname = os.Getenv("UNAME_PRODDUCTION")
		password = os.Getenv("PASS_PRODDUCTION")
		dbname = os.Getenv("DBNAME_PRODDUCTION")
	} else if env == "development" {
		port = os.Getenv("PORT_DEVELOPMENT")
		host = os.Getenv("HOST_DEVELOPMENT")
		uname = os.Getenv("UNAME_DEVELOPMENT")
		password = os.Getenv("PASS_DEVELOPMENT")
		dbname = os.Getenv("DBNAME_DEVELOPMENT")
	} else {
		port = os.Getenv("PORT_TESTING")
		host = os.Getenv("HOST_TESTING")
		uname = os.Getenv("UNAME_TESTING")
		password = os.Getenv("PASS_TESTING")
		dbname = os.Getenv("DBNAME_TESTING")
	}
}

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Routes() {
	r := mux.NewRouter()
	graph := r.PathPrefix("/graph").Subrouter()
	a.Router = graph
	a.listRouter()
}

func (a *App) listRouter() {
	router := routers.Router{}
	router.Router = a.Router
	router.DB = a.DB

	router.ProdRouters()
}

func (a *App) Run() {
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	http.ListenAndServe(port, handlers.CORS(headers, methods, origins)(a.Router))
}

func (a *App) ConnectDB() {
	strCon := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", uname, password, host, port, dbname)
	db, err := sql.Open("mysql", strCon)
	if err != nil {
		helpers.Error(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	a.DB = db
}
