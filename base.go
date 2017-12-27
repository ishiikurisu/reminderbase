package reminderbase

import (
    "os"
    "fmt"
    "time"
    "encoding/json"
    "io/ioutil"
)

// Definition of database
type Base struct {
    Source string
    Users []int
}

// Creates a new standard database
func NewBase() *Base {
    base := Base {
        Source: "data/reminderbase",
        Users: make([]int, 0),
    }
    return &base
}

// Gets all reminders from the database
func (base *Base) GetAllReminders() []Reminder {
    reminders := make([]Reminder, 0)

    for _, user := range base.Users {
        for _, reminder := range base.GetAllRemindersFromUser(user) {
            reminders = append(reminders, reminder)
        }
    }

    return reminders
}

// Adds a reminder from an user
func (base *Base) AddReminder(user int, reminder Reminder) error {
    userPath := fmt.Sprintf("%s/%d", base.Source, user)

    // Adding user to memory if necessary
    flag := false
    for _, u := range base.Users {
        if u == user {
            flag = true
        }
    }
    if !flag {
        base.Users = append(base.Users, user)
        oops := os.Mkdir(userPath, os.ModeDir)
        if oops != nil {
            return oops
        }
    }

    // Saving reminder to memory
    reminderPath := fmt.Sprintf("%s/%d.json", userPath, time.Now().Unix())
    raw, oops := json.Marshal(reminder)
    if oops != nil {
        return oops
    }
    oops = ioutil.WriteFile(reminderPath, raw, 0644)
    if oops != nil {
        return oops
    }

    return nil
}

// Gets all reminders from an user.
func (base *Base) GetAllRemindersFromUser(user int) []Reminder {
    reminders := make([]Reminder, 0)
    userPath := fmt.Sprintf("%s/%d", base.Source, user)
    files, _ := ioutil.ReadDir(userPath)

    for _, file := range files {
        var reminder Reminder
        raw, _ := ioutil.ReadFile(fmt.Sprintf("%s/%s", userPath, file.Name()))
        json.Unmarshal(raw, &reminder)
        reminders = append(reminders, reminder)
    }

    return reminders
}
