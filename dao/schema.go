package dao

import "time"

const (
	TableName_User       = "users"
	TableName_Work       = "works"
	TableName_Rest       = "rests"
	TableName_Leaderlist = "leaderlists"
	TableName_Worklist   = "worklists"
)

type User struct {
	Name     string `gorm:"primary_key" json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Work struct {
	ID            uint      `gorm:"primary_key" json:"id"`
	Name          string    `json:"name"`
	Date          time.Time `json:"date"`
	Worktime      uint      `json:"worktime"`
	Leaderlistid  uint      `json:"leaderlistid"`
	Examine       string    `json:"examine"`
	Availabletime uint      `json:"availabletime"`
	Committime    time.Time `json:"committime"`
	Reason        string    `json:"reason"`
}
type Rest struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	Name         string    `json:"name"`
	Date         time.Time `json:"date"`
	Noon         string    `json:"noon"`
	Resttime     uint      `json:"resttime"`
	Worklistid   uint      `json:"worklistid"`
	Leaderlistid uint      `json:"leaderlistid"`
	Examine      string    `json:"examine"`
	Committime   time.Time `json:"committime"`
	Reason       string    `json:"reason"`
}
type Leaderlist struct {
	ID           uint   `gorm:"primary_key" json:"id"`
	Leaderlistid uint   `json:"leaderlistid"`
	Leadername   string `json:"leadername"`
	Examineby    string `json:"examineby"`
}
type Worklist struct {
	ID         uint `gorm:"primary_key" json:"id"`
	Worklistid uint `json:"worklistid"`
	Workid     uint `json:"workid"`
	Utime      uint `json:"utime"`
}
