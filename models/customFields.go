package models

type (
	CustomField struct {
		Id     int
		Name   string
		Values []struct {
			Value string
			Enum  string
		}
		IsSystem bool `json:"is_system"`
	}
)
