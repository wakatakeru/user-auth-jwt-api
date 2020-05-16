package usecase

import "github.com/wakatakeru/user-auth-jwt-api/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func NewUserInteractor(UserRepository UserRepository) UserInteractor {
	UserInteractor := UserInteractor{UserRepository: UserRepository}
	return UserInteractor
}

func (interactor *UserInteractor) Add(u domain.User) (id int, err error) {
	id, err = interactor.UserRepository.Store(u)
	return
}

func (interactor *UserInteractor) Update(u domain.User) (user domain.User, err error) {
	user, err = interactor.UserRepository.Update(u)
	return
}

func (interactor *UserInteractor) UserByName(u string) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindByName(u)
	return
}

func (interactor *UserInteractor) UserByID(uid int) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindByID(uid)
	return
}
