package core

import (
	"database/sql"

	// for driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// AppContext contains app dependencies
type AppContext struct {
	MainRouter *mux.Router
	APIRouter  *mux.Router
	DB         *sql.DB
}
