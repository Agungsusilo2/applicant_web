package domain

type NextHeroes struct {
	Id        string `json:"Id" gorm:"primaryKey"`
	Countries string `json:"Countries" gorm:"type:varchar(200);not null"`
	Skills    string `json:"Skills" gorm:"type:varchar(255);not null"`
	Username  string `json:"Username" gorm:"type:varchar(100);not null"`
	Password  string `json:"Password" gorm:"type:varchar(100);not null"`
}
