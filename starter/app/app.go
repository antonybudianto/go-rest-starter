package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"starter/core"
	"starter/user"
	// for driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// App level struct containing its dependencies
type App struct {
	AppCtx *core.AppContext
}

const dbDriver = "mysql"
const dbName = "rest_api_example"

var dbUsername = os.Getenv("DB_USERNAME")
var dbPassword = os.Getenv("DB_PASSWORD")

// Initialize App dependencies
func (a *App) Initialize() {
	connectionString := fmt.Sprintf("%s:%s@tcp(db:3306)/%s", dbUsername, dbPassword, dbName)

	var err error
	a.AppCtx = &core.AppContext{}

	a.AppCtx.DB, err = sql.Open(dbDriver, connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.AppCtx.MainRouter = mux.NewRouter()
	a.AppCtx.APIRouter = a.AppCtx.MainRouter.PathPrefix("/api").Subrouter()
	a.initializeRoutes()
}

// Run app
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.AppCtx.MainRouter))
}

func (a *App) initializeRoutes() {
	user.New(a.AppCtx.DB, a.AppCtx.APIRouter)
}
