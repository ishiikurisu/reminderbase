package reminderbase

import (
    "testing"
)

func TestCanAddAndListEverythingFromUser(t *testing.T) {
    base := NewBase()
    allListings := base.GetAllReminders()
    if len(allListings) > 0 {
        t.Error("Clear the database, please?")
        return
    }

    reminder := NewReminder("Do something")
    base.AddReminder(1, *reminder)
    allListings = base.GetAllReminders()
    if len(allListings) != 1 {
        t.Error("Couldn't add this reminder")
    }
    allListings = base.GetAllRemindersFromUser(1)
    if len(allListings) != 1 {
        t.Error("Couldn't load reminders from that user")
    }
}
