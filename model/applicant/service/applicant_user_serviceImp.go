package service

import (
	"Applicant_user/model/applicant/domain"
	"Applicant_user/model/applicant/dto"
	"Applicant_user/model/applicant/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ApplicantServiceImpl struct {
	applicantRepository repository.ApplicantRepository
}

func NewApplicantService(applicantRepository repository.ApplicantRepository) ApplicatService {
	return &ApplicantServiceImpl{applicantRepository: applicantRepository}
}

func (a *ApplicantServiceImpl) GetAll() []domain.ApplicantUser {
	return a.applicantRepository.FindAll()
}

func (a *ApplicantServiceImpl) GetByID(id string) domain.ApplicantUser {
	return a.applicantRepository.FindById(id)
}

func (a *ApplicantServiceImpl) Create(ctx *gin.Context) (*domain.ApplicantUser, error) {
	var input dto.CreateApplicantUserReq

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}
	uniqueID := uuid.New()
	unique := (uniqueID).String()

	// bytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	// if err != nil {
	// 	return nil, err
	// }
	// password := string(bytes)
	applicant := domain.ApplicantUser{
		Id:           unique,
		NamaLengkap:  input.NamaLengkap,
		Email:        input.Email,
		NomorTelepon: input.NomorTelepon,
		Umur:         input.Umur,
		Pekerjaan:    input.Pekerjaan,
	}
	result, err := a.applicantRepository.Save(applicant)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *ApplicantServiceImpl) Update(ctx *gin.Context) (*domain.ApplicantUser, error) {
	var input dto.UpdateApplicantReq

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()
	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	id := ctx.Param("id")
	if err != nil {
		return nil, err
	}

	// newPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	// if err != nil {
	// 	return nil, err
	// }
	// newPassword := string(newPasswordBytes)

	updatedApplicant := domain.ApplicantUser{
		Id:           id,
		NamaLengkap:  input.NamaLengkap,
		Email:        input.Email,
		NomorTelepon: input.NomorTelepon,
		Umur:         input.Umur,
		Pekerjaan:    input.Pekerjaan,
	}

	result, err := a.applicantRepository.Update(updatedApplicant)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *ApplicantServiceImpl) Delete(ctx *gin.Context) (*domain.ApplicantUser, error) {
	id := ctx.Param("id")
	applicant := domain.ApplicantUser{
		Id: id,
	}

	result, err := a.applicantRepository.Delete(applicant)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *ApplicantServiceImpl) NextRegHeroes(ctx *gin.Context) (*domain.ApplicantUser, error) {
	var input dto.NextStepRegHeroes

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()
	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	id := ctx.Param("id")
	if err != nil {
		return nil, err
	}

	// newPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	// if err != nil {
	// 	return nil, err
	// }
	// newPassword := string(newPasswordBytes)

	NextStepRegHeroes := domain.ApplicantUser{
		Id:        id,
		Countries: input.Countries,
		Skills:    input.Skills,
		Username:  input.Username,
		Password:  input.Password,
	}

	result, err := a.applicantRepository.NextRegHeroes(NextStepRegHeroes)
	if err != nil {
		return nil, err
	}

	return result, nil
}
