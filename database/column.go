package database

type Column struct {
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	Length      string      `json:"length"`
	Default     interface{} `json:"default"`
	IsNull      bool        `json:"is_null"`
	IsPrimary   bool        `json:"in_primary"`
	IsUnique    bool        `json:"is_unique"`
	IsReference bool        `json:"is_reference"`
	References  Reference   `json:"references"`
}

type Columns []Column
