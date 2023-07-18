package service

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/Hitsa/CacaTesouro/internal/repository"
)

type UserService struct {
	Repository *repository.UsersRepository
}

func NewUserService(r *repository.UsersRepository) *UserService {
	user := UserService{r}
	return &user
}

func (s UserService) CreateUser(name string, userName string, password string, email string, role string) error {
	err := validateName(name)
	if err != nil {
		return err
	}
	valid := validateEmail(email)
	if !valid {
		return errors.New("email invalido")
	}

	dateNow := time.Now()
	err = s.Repository.InsertDbUsers(name, userName, password, email, role, dateNow)
	if err != nil {
		return err
	}
	return nil
}

func validateName(name string) error {
	reg := regexp.MustCompile(`[[:alpha:]]+`)
	nameExtract := reg.FindAllString(name, -1)
	if nameExtract == nil {
		return fmt.Errorf("nome inválido")
	}
	return nil
}

func validateEmail(email string) bool {
	reg := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	emailExtract := reg.MatchString(email)
	return emailExtract
}

func (s UserService) SearchUser(userName string) (*repository.Users, error) {

	UserArray, err := s.Repository.SearchDbUser(userName)
	if err != nil {
		return nil, err
	}
	return UserArray, nil
}

func (s UserService) DeleteUser(userName string) (bool, error) {

	valid, err := s.Repository.DeleteUserDB(userName)
	if err != nil {
		return false, nil
	}
	return valid, nil
}

func (s UserService) ChangeUserData(name string, userName string, password string, email string, role string) (*repository.Users, error) {

	// Consulta Db para ver se já existe usuário

	userArray, err := s.Repository.ChangeUserDataDB(name, userName, password, email, role)
	if err != nil {
		return nil, err
	}
	// Return
	return userArray, nil
}
