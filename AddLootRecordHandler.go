package Openapi

import "time"
import "database/sql"
import "log"

var (
	connectionName = mustGetenv("CLOUDSQL_CONNECTION_NAME")
	user           = mustGetenv("CLOUDSQL_USER")
	password       = os.Getenv("CLOUDSQL_PASSWORD")
	tableName	   = mustGetenv("TABLE_NAME")
)

func AddLootRecord(ctx echo.Context) error {
	// get date time
	now = time.Now()
	// db connect
	db, err = sql.Open("psql", fmt.Sprintf("%s:%s@cloudsql(%s)/", user, password, connectionName))
	if err ! = nil {
		log.Fatalf("Could not open db: %v", err)
	}
	//get datas from ctx
	//TODO
	// make psql insert call
	//username submit_timestamp loot[id][amt] mission_id
	sqlStatement := `
	INSERT INTO $1 (username, submit_timestamp, loot, mission_id)
	VALUES ($2, $3, $4, $5)`
	_, err = db.Exec(sqlStatement, tableName, "mary", '2016-06-22 19:10:25-07', "{{1,2}}", "1")
	if err != nil {
		panic(err)
	}
	// return ok
	return ctx.JSON(http.StatusOK, "HOKAYDOKAY")
}