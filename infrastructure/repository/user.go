package repository

import (
	"seatmap-backend/entities"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (repo *UserRepo)Create(userInput *entities.User) (*entities.User, error) {
	err := repo.DB.Create(userInput).Error
	if err != nil {
		return nil, err
	}
	return userInput ,nil
}

func (repo *UserRepo)Get(username string) (u *entities.User, err error) {
	return getUser(username, repo)
}

func getUser(username string, repo *UserRepo) (u *entities.User, err error) {
	user := new(entities.User)
	err = repo.DB.Select("role","full_name","username", "password", "salt").Where(map[string]interface{}{"username": username}).Find(&user).Error
	if err!= nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepo)List() (user *[]entities.User,err error) {
	users := make([]entities.User, 10)
	err = repo.DB.Select("role","full_name","username").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}