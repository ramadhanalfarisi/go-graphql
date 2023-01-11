package app

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/ramadhanalfarisi/go-graphql-kocak/helpers"
	"github.com/ramadhanalfarisi/go-graphql-kocak/routers"
	"net/http"
	"os"
	"time"
)

var host, uname, password, port, dbname string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		helpers.Error(err)
	}
	if env := os.Getenv("ENVIRONMMENT"); env == "production" {
		port = os.Getenv("DB_PORT")
		host = os.Getenv("DB_HOST")
		uname = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname = os.Getenv("DB_NAME")
	} else if env == "development" {
		port = os.Getenv("DB_PORT_DEV")
		host = os.Getenv("DB_HOST_DEV")
		uname = os.Getenv("DB_USER_DEV")
		password = os.Getenv("DB_PASSWORD_DEV")
		dbname = os.Getenv("DB_NAME_DEV")
	} else {
		port = os.Getenv("DB_PORT_TEST")
		host = os.Getenv("DB_HOST_TEST")
		uname = os.Getenv("DB_USER_TEST")
		password = os.Getenv("DB_PASSWORD_TEST")
		dbname = os.Getenv("DB_NAME_TEST")
	}
}

type App struct {
	Router      *mux.Router
	RouterGraph *mux.Router
	DB          *sql.DB
}

func (a *App) Routes() {
	r := mux.NewRouter()
	graph := r.PathPrefix("/graph").Subrouter()
	a.Router = r
	a.RouterGraph = graph
	a.listRouter()
}

func (a *App) listRouter() {
	router := routers.Router{}
	router.Router = a.Router
	router.RouterGraph = a.RouterGraph
	router.DB = a.DB

	router.ProdRouters()
}

func (a *App) Run() {
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(a.Router))
}

func (a *App) Migrate() {
	driver, err := postgres.WithInstance(a.DB, &postgres.Config{})
	if err != nil {
		helpers.Error(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres", driver)
	if err != nil {
		helpers.Error(err)
	}
	err2 := m.Up()
	if err2 != nil {
		helpers.Error(err2)
	}
}

func (a *App) ConnectDB() {
	strCon := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, uname, password, dbname)
	db, err := sql.Open("postgres", strCon)
	if err != nil {
		helpers.Error(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	a.DB = db
	a.Migrate()
}
