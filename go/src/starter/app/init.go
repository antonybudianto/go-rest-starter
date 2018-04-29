package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"go-rest-starter/go/src/starter/routes/user"
	// for driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// App level struct containing its dependencies
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize App dependencies
func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@tcp(db:3306)/%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

// Run app
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	u := user.Handler{Router: a.Router, DB: a.DB}
	u.InitializeRoutes()
}
