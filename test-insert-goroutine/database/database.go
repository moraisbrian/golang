package database

import (
	"database/sql"
	"fmt"
	"go-test/models"

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

func getConnection() Database {
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

func CreateReminder(title string, description string, alias string) (int64, error) {

	conn := getConnection()
	defer conn.SqlDb.Close()

	queryText := fmt.Sprintf("INSERT INTO Reminders (Title, Description, Alias) VALUES ('%s', '%s', '%s');", title, description, alias)
	fmt.Println(queryText)

	res, err := conn.SqlDb.Exec(queryText)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func ReadReminders() ([]models.Reminder, error) {
	conn := getConnection()
	defer conn.SqlDb.Close()

	queryText := "SELECT * FROM Reminders"

	rows, err := conn.SqlDb.Query(queryText)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	reminders := make([]models.Reminder, 0)

	for rows.Next() {
		var reminder models.Reminder

		err := rows.Scan(&reminder.ID, &reminder.Title, &reminder.Description, &reminder.Alias)

		if err != nil {
			return nil, err
		}

		reminders = append(reminders, reminder)
	}

	return reminders, nil
}

func DeleteReminder(id int) (int64, error) {
	conn := getConnection()
	defer conn.SqlDb.Close()

	queryText := fmt.Sprintf("DELETE FROM Reminders WHERE ID = '%v'", id)

	res, err := conn.SqlDb.Exec(queryText)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func UpdateReminder(reminder models.Reminder) (int64, error) {
	conn := getConnection()
	defer conn.SqlDb.Close()

	queryText := fmt.Sprintf(`
				UPDATE Reminders 
				SET Title = '%s', 
				Description = '%s',
				Alias = '%s'
				WHERE ID = '%v'`,
		reminder.Title, reminder.Description, reminder.Alias, reminder.ID)

	res, err := conn.SqlDb.Exec(queryText)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
