package domain

type ApplicantUser struct {
	Id           string `json:"Id" gorm:"primaryKey"`
	NamaLengkap  string `json:"NamaLengkap" gorm:"type:varchar(150);not null"`
	Email        string `json:"Email" gorm:"type:varchar(100);not null"`
	NomorTelepon string `json:"NomorTelepon" gorm:"type:varchar(100);not null"`
	Umur         string `json:"gmail" gorm:"type:varchar(100);not null"`
	Pekerjaan    string `json:"telepon" gorm:"type:varchar(100);not null"`
	Countries    string `json:"Countries" gorm:"type:varchar(200);not null"`
	Skills       string `json:"Skills" gorm:"type:varchar(255);not null"`
	Username     string `json:"Username" gorm:"type:varchar(100);not null"`
	Password     string `json:"Password" gorm:"type:varchar(100);not null"`
}
