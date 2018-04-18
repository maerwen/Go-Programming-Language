package main

type Girl struct {
	Name  string
	Age   int
	Like  string
	Email string `json:"邮箱,omitempty"`
}
type SimpleGirl struct {
	Name string
}
