package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"starter/routes/user"
	// for driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// App level struct containing its dependencies
type App struct {
	MainRouter *mux.Router
	APIRouter  *mux.Router
	DB         *sql.DB
}

const dbDriver = "mysql"
const dbName = "rest_api_example"

var dbUsername = os.Getenv("DB_USERNAME")
var dbPassword = os.Getenv("DB_PASSWORD")

// Initialize App dependencies
func (a *App) Initialize() {
	connectionString := fmt.Sprintf("%s:%s@tcp(db:3306)/%s", dbUsername, dbPassword, dbName)

	var err error
	a.DB, err = sql.Open(dbDriver, connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.MainRouter = mux.NewRouter()
	a.APIRouter = a.MainRouter.PathPrefix("/api").Subrouter()
	a.initializeRoutes()
}

// Run app
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.MainRouter))
}

func (a *App) initializeRoutes() {
	u := user.Handler{APIRouter: a.APIRouter, DB: a.DB}
	u.InitializeRoutes()
}
