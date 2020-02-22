package models

import (
	"time"
)

// Folder : struct for folder
type Folder struct {
	ID          string       `db:"id" json:"ID"`
	Name        string       `db:"name" json:"Name"`
	Status      string       `db:"status" json:"Status"`
	DateCreated time.Time    `db:"date_created" json:"DateCreated"`
	DateEdited  time.Time    `db:"date_edited" json:"DateEdited"`
	Items       []FolderItem `json:"FolderItems"`
}

// FolderItem : stores item in folders - notes, cornell notes, pdf
type FolderItem struct {
	ID   string `db:"id" json:"ID"`
	Name string `db:"name" json:"Name"`
}

// Folders : list of folders
var Folders []Folder
