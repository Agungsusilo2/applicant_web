package dto

type UpdateApplicantReq struct {
	NamaLengkap  string `json:"NamaLengkap" validate:"required,max=50,min=1"`
	Email        string `json:"Email" validate:"email,required,max=50,min=1"`
	NomorTelepon string `json:"NomorTelepon" validate:"numeric,required,max=25,min=1"`
	Umur         string `json:"Umur" validate:"numeric,required,min=1"`
	Age          string `json:"Age" validate:"required,gte=16,lte=100,numeric"`
	Pekerjaan    string `json:"Pekerjaan" validate:"required,min=1,max=25"`
}
