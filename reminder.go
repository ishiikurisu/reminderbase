package reminderbase

// Reminder definition
type Reminder struct {
    Content string
}

// Creates a reminder from a simple text
func NewReminder(content string) *Reminder {
    reminder := Reminder {
        Content: content,
    }
    return &reminder
}
