package models

type Stat struct {
	Name    string `json:"name"`
	Assists int    `json:"assists"`
	Goals   int    `json:"goals"`
	Date    string
}

type Sheet struct {
	Date  string `json:"date"`
	Stats []Stat `json:"stats"`
}
