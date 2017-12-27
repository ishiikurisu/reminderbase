package reminderbase

type Base struct {
    Source string
    Users []int
}

func NewBase() *Base {
    base := Base {
        Source: "data/reminderbase",
        Users: make([]int, 0),
    }
    return &base
}

func (base *Base) GetAllReminders() []Reminder {
    return make([]Reminder, 0)
}
