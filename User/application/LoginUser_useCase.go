package application

import (
    "github.com/bchanona/profile_warmheart_backend/User/domain"
    "golang.org/x/crypto/bcrypt"
)

type LoginUserUseCase struct {
    Repo domain.UserRepository
}

func NewLoginUserUseCase(repo domain.UserRepository) *LoginUserUseCase {
    return &LoginUserUseCase{Repo: repo}
}

func (uc *LoginUserUseCase) Login(login domain.LoginRequest) (domain.User, error) {
    user, err := uc.Repo.GetByEmail(login.Email)
    if err != nil {
        return domain.User{}, domain.ErrInvalidCredentials
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
    if err != nil {
        return domain.User{}, domain.ErrInvalidCredentials
    }

    user.Password = ""
    return user, nil
}