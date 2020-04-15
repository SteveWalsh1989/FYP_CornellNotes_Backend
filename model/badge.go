package models

// Badge : Sample struct for Badge to test router
type Badge struct {
	ID          string `db:"id" json:"ID"`
	Title       string `db:"title" json:"Title"`
	Requirement string `db:"requirement" json:"Requirement"`
}

// Badges : badges
var Badges []Badge
