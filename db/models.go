package db

type Picture struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Url  string `json:"url"`
}
