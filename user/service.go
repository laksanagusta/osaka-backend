package user

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	UserDetails(id int) (User, error)
	UserFindAll() ([]User, error)
	UpdateUser(inputID GetUserDetailInput, inputData UpdateUserInput) (User, error)
	DeleteUser(inputID GetUserDetailInput) (string, error)
	CreateUserBulk(workerIndex int, counter int, jobs []interface{}) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Username = strings.ToLower(input.Username)
	user.Name = input.Name
	user.Email = strings.ToLower(input.Email)
	user.Occupation = input.Occupation
	email := input.Email

	//Check if email already taken
	isEmailExist, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if isEmailExist.ID != 0 {
		return user, errors.New("Register failed, Email already exist")
	}

	isUsernameExist, err := s.repository.FindByUsername(input.Username)

	if err != nil {
		return user, err
	}

	if isUsernameExist.ID != 0 {
		return user, errors.New("Register failed, username already been taken")
	}

	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(PasswordHash)

	user.Role = "USER"
	if input.Role != "" {
		user.Role = strings.ToUpper(input.Role)
	}

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) UserDetails(id int) (User, error) {
	user_id := id
	user, err := s.repository.FindById(user_id)
	if user.ID == 0 {
		return user, errors.New("User not found")
	}
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *service) UserFindAll() ([]User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *service) UpdateUser(inputID GetUserDetailInput, inputData UpdateUserInput) (User, error) {
	user, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return user, err
	}

	user.Username = strings.ToLower(inputData.Username)
	user.Name = inputData.Name
	user.Email = strings.ToLower(inputData.Email)
	user.Occupation = inputData.Occupation
	user.Role = "USER"

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil

}

func (s *service) DeleteUser(inputID GetUserDetailInput) (string, error) {
	updatedUser, err := s.repository.Delete(inputID.ID)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) UpdateUserRole(inputID GetUserDetailInput, inputData UpdateUserInput) (User, error) {
	user, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return user, err
	}

	user.Name = inputData.Name
	user.Email = inputData.Email
	user.Occupation = inputData.Occupation
	user.Role = "ADMIN"

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil

}

func (s *service) CreateUserBulk(workerIndex int, counter int, jobs []interface{}) (User, error) {
	user := User{}
	user.Name = fmt.Sprintf("%v", jobs[0])

	newUsers, err := s.repository.Save(user)
	if err != nil {
		return newUsers, err
	}

	return newUsers, nil
}
