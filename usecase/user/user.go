package user

import "seatmap-backend/entities"

type userUsecase struct {
	repo UserRepository
}

func NewService(r UserRepository) *userUsecase {
	return &userUsecase{
		repo: r,
	}
}

func (s *userUsecase) GetUser(username string) (*entities.User, error) {
	return s.repo.Get(username)
}

func (s *userUsecase) ListUsers() (user *[]entities.User,err error) {
	return s.repo.List()
}

func (s *userUsecase) CreateUser(userInput *entities.User) (*entities.User, error) {
	// e, err := entities.NewUser(userInput)
	// if err != nil {
	// 	return nil, err
	// }
	return s.repo.Create(userInput)
}