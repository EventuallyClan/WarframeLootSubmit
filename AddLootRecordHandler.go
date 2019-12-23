package Openapi

import "time"
import "database/sql"
import "log"

var (
	connectionName = mustGetenv("CLOUDSQL_CONNECTION_NAME")
	user           = mustGetenv("CLOUDSQL_USER")
	password       = os.Getenv("CLOUDSQL_PASSWORD")
)

func AddLootRecord(ctx echo.Context) error {
	// get date time
	now = time.Now()
	// db connect
	db, err = sql.Open("psql", fmt.Sprintf("%s:%s@cloudsql(%s)/", user, password, connectionName))
	if err ! = nil {
		log.Fatalf("Could not open db: %v", err)
	}
	// make psql insert call
	// return ok
}