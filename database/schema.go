package database

type Schema struct {
	Name   string `json:"name"`
	Tables Tables `json:"tables"`
}
