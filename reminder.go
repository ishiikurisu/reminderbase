package reminderbase

type Reminder struct {
    Content string
}

func NewReminder(content string) *Reminder {
    reminder := Reminder {
        Content: content,
    }
    return &reminder
}
