package db

type Picture struct {
	ID   string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name string `json:"name"`
}
