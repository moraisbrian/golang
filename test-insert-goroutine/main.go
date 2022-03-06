package main

import (
	"fmt"
	"go-test/database"
	"go-test/models"
	"sync"
)

var wg sync.WaitGroup

func main() {
	createReminders()
	// readReminders()
	// deleteReminder()
	// updateReminder()
}

func updateReminder() {
	reminder := models.Reminder{
		ID:          955,
		Title:       "Novo título",
		Description: "Nova descrição",
		Alias:       "Novo alias",
	}

	res, err := database.UpdateReminder(reminder)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}
}

func deleteReminder() {
	res, err := database.DeleteReminder(956)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}
}

func createReminders() {
	reminders := seedReminders()
	channel := make(chan models.Reminder)

	for i := 1; i <= 50; i++ {
		wg.Add(1)
		go worker(channel)
	}

	for i := range reminders {
		channel <- reminders[i]
	}

	close(channel)
	wg.Wait()
}

func readReminders() {
	res, err := database.ReadReminders()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		for i := range res {
			fmt.Printf("%v, %v, %v, %v\n", res[i].ID, res[i].Title, res[i].Description, res[i].Alias)
		}
	}
}

func worker(channel chan models.Reminder) {
	defer wg.Done()
	for reminder := range channel {
		res, err := database.CreateReminder(reminder.Title, reminder.Description, reminder.Alias)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(res)
		}

		fmt.Println(reminder)
	}
}

func seedReminders() []models.Reminder {
	reminders := make([]models.Reminder, 1000)

	for i := range reminders {
		reminders[i].Title = fmt.Sprintf("Título %v", i)
		reminders[i].Description = fmt.Sprintf("Descrição %v", i)
		reminders[i].Alias = fmt.Sprintf("Alias %v", i)
	}

	return reminders
}
