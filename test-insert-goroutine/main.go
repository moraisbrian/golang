package main

import (
	"fmt"
	"go-test/database"
)

type Reminder struct {
	Title       string
	Description string
	Alias       string
}

func main() {
	reminders := createReminders()
	channel := make(chan Reminder)

	for i := 1; i <= 100; i++ {
		go worker(channel)
	}

	for i := range reminders {
		channel <- reminders[i]
	}
}

func worker(channel chan Reminder) {
	for reminder := range channel {
		func() {
			conn := database.GetConnection()
			defer conn.SqlDb.Close()

			res, err := conn.CreateReminder(reminder.Title, reminder.Description, reminder.Alias)

			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(res)
			}

			fmt.Println(reminder)
		}()
	}
}

func createReminders() []Reminder {
	reminders := make([]Reminder, 1000)

	for i := range reminders {
		reminders[i].Title = fmt.Sprintf("Título %v", i)
		reminders[i].Description = fmt.Sprintf("Descrição %v", i)
		reminders[i].Alias = fmt.Sprintf("Alias %v", i)
	}

	return reminders
}
