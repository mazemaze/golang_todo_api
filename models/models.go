package models

type Task struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Text   string `json:"text"`
	IsDone bool   `json:"isdone"`
}

var Tasks = []Task{
	{ID: "1", Name: "John", Text: "Get a job", IsDone: false},
	{ID: "2", Name: "Loki", Text: "Treat people nicely", IsDone: false},
	{ID: "3", Name: "Sevasthean", Text: "Clean my room", IsDone: false},
}
