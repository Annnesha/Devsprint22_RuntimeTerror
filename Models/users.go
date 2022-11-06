package Models

type Events struct {
	EventID          string `gorm:"unique"`
	Title            string
	Organizer        string
	Student_Name     string
	Roll             string
	Branch           string
	School_concerned string
	Description      string
	Venue            string
	Room_Allot       string
	Start_date       string
	End_date         string
	Start_time       string
	End_time         string
	Comment          string
	Status           string
}

type Teacher struct {
	First_Name string
	Last_Name  string
	Email      string `gorm:"unique"`
	Password   []byte
	Faculty_ID string
	School     string
	Position   string
	EventID    string `gorm:"foreignKey:EventID"`
	Events     Events
}

type Student struct {
	First_Name string
	Last_Name  string
	Email      string `gorm:"unique"`
	Password   []byte
	Roll       string
	Branch     string
	EventID    string `gorm:"foreignKey:EventID"`
	Events     Events
}
