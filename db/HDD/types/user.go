package types

type User struct {
	ID          int64    `json:"t_id"`
	UserName    string   `json:"user_name"`
	Image       []string `json:"image"`
	Description string   `json:"description"`
	Hobby       []string `json:"hobby"`
	IsValid     bool     `json:"is_valid"`
	Latitude    float64  `json:"latitude"`
	Hardness    float64  `json:"hardness"`
	Location    string   `json:"location"`
}

type AroundUser struct {
	UserName string   `json:"user_name"`
	Image    []string `json:"image"`
	Latitude float64  `json:"latitude"`
	Hardness float64  `json:"hardness"`
}
