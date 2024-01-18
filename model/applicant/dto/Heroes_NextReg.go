package dto

type NextStepRegHeroes struct {
	Countries string `json:"Countries" validate:"required`
	Skills    string `json:"Skills" validate:"required"`
	Username  string `json:"Username" validate:"required"`
	Password  string `json:"Password" validate:"required"`
}
