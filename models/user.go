package models

type User struct {
	ID         int32  `json:"id" gorm:"primary_key"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Age        int    `json:"age"`
	Occupation string `json:"occupation"`
}

type Response struct {
	Message string      `json:"message,omitempty"`
	Body    interface{} `json:"body,omitempty"`
}
