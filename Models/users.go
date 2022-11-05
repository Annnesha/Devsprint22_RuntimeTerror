package models

type Events struct{
  EventID string  `gorm:"unique"`
  Title   string
  Organizer  string
  Student_Name  string
  Roll  int
  Branch  string
  School_concerned  string
  Description string
  Venue  string
  Request_time  time.Now()
  Start_date  string
  End_date  string
  Start_time  string
  End_time  string
  Status string
}

type Teacher struct{
    First_Name string
    Last_Name string
    Email   string  `gorm:"unique"`
    Password []byte
    Faculty_ID string
    School  string
    Position  string
    EventID  string  `gorm:"foreignKey:EventID"`
    Events   Events
}

type Student struct{
  First_Name string
  Last_Name string
  Email   string  `gorm:"unique"`
  Password []byte
  Roll int
  Branch  string
  EventID  string  `gorm:"foreignKey:EventID"`
  Events   Events
}