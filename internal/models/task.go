package models

type Task struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Done    bool   `json:"done"`
	User_id int    `json:"userId"`
}
