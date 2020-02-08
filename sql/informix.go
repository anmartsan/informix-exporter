package informix

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/alexbrainman/odbc"
)

func OpenDatabase(informixServer string) *sql.DB {

	fmt.Printf("Conectando a %s:\n", informixServer)
	db, err := sql.Open("odbc", informixServer)

	if err != nil {

		log.Println("Error in Open Database: ", err)
	}
	return db
}

func CloseDatabase(db *sql.DB) {

	err := db.Close()

	if err != nil {

		log.Println("Error in Close Database: ", err)
	}

}

func QueryDatabase(db *sql.DB, query string) (*sql.Rows, error) {

	rows, error := db.Query(query)
	if error != nil {
		fmt.Println(error)
	}
	return rows, error

}
