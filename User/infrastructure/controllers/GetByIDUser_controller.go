package controllers

import "github.com/bchanona/profile_warmheart_backend/User/application"

type GetByIDUserController struct {
	UseCase *application.GetByIDUserUseCase
}

func NewGetByIDUserController(useCase *application.GetByIDUserUseCase) *GetByIDUserController{
	return &GetByIDUserController{UseCase: useCase}
}
