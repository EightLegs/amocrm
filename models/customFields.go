package models

type (
	CustomField struct {
		Id     int `json:"id"`
		Name   string `json:"name"`
        Values []CustomValue `json:"values"`
		IsSystem bool `json:"is_system"`
	}

	CustomValue struct {
	  Value string `json:"value"`
	  Enum string `json:"enum"`
	}
)
