package database

type Table struct {
	Name    string  `json:"name"`
	Columns Columns `json:"columns"`
}

type Tables []Table
