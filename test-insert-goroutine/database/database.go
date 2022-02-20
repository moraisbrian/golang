package database

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

type Database struct {
	SqlDb *sql.DB
}

var (
	server   = "localhost"
	port     = 1433
	user     = "sa"
	password = "Password!"
	database = "GoTeste"
)

func GetConnection() Database {
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)

	sqlObj, err := sql.Open("mssql", connectionString)
	if err != nil {
		fmt.Println(err.Error())
	}

	data := Database{
		SqlDb: sqlObj,
	}

	return data
}

func (db Database) CreateReminder(title string, description string, alias string) (int64, error) {

	queryText := fmt.Sprintf("INSERT INTO Reminders (Title, Description, Alias) VALUES ('%s', '%s', '%s');", title, description, alias)
	fmt.Println(queryText)

	res, err := db.SqlDb.Exec(queryText)

	if err != nil {
		return 0, nil
	}

	return res.RowsAffected()
}
