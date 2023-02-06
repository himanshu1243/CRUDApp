package sqlconnection

//hello world
import (
	"database/sql"
	"log"
)

var db *sql.DB

func GetMySQLDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/studentinfo")
	if err != nil {
		log.Fatal("No database found")

	}
	return db
}
