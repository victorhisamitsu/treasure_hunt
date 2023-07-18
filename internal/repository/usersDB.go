package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Users struct {
	bun.BaseModel `bun:"users"`
	Id            string    `json:"id" bun:"id,pk,type:uuid"`
	Name          string    `json:"name" bun:"name"`
	UserName      string    `json:"userName" bun:"username"`
	Password      string    `json:"password" bun:"password"`
	Email         string    `json:"email" bun:"email"`
	CreatedDate   time.Time `json:"createdDate" bun:"data"`
	Role          string    `json:"role" bun:"role"`
}

type UsersRepository struct {
	DB *bun.DB
}

func NewRepositoryUser(d *bun.DB) *UsersRepository {
	return &UsersRepository{
		DB: d,
	}
}

func (r *UsersRepository) InsertDbUsers(name string, userName string, password string, email string, role string, createdDate time.Time) error {

	// Consulta Db para ver se já existe usuário
	userArray := []Users{}
	count, err := r.DB.NewSelect().Model(&userArray).Where("username=?", userName).ScanAndCount(context.Background())
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("userName já ultilizado")
	}
	count, err = r.DB.NewSelect().Model(&userArray).Where("email=?", email).ScanAndCount(context.Background())
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("email já ultilizado")
	}

	// Caso não exista adicionar ao banco
	user := Users{
		Id:          uuid.NewString(),
		Name:        name,
		UserName:    userName,
		Password:    password,
		Email:       email,
		CreatedDate: createdDate,
		Role:        role,
	}

	res, err := r.DB.NewInsert().Model(&user).Exec(context.Background())
	if err != nil {
		return err
	}

	// Return
	fmt.Println(res)
	fmt.Println("Dados inseridos com sucesso.")
	return nil
}

func (r *UsersRepository) SearchDbUser(userName string) (*Users, error) {

	// Buscar no Db usuário

	UserArray := Users{}
	_, err := r.DB.NewSelect().Model(&UserArray).Where("username=?", userName).Exec(context.Background(), &UserArray)
	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}
	return &UserArray, nil
}

func (r *UsersRepository) DeleteUserDB(userName string) (bool, error) {

	//Ligação no db
	userArray := []Users{}

	//Executar query
	res, err := r.DB.NewDelete().Model(&userArray).Where("username=?", userName).Exec(context.Background())
	if err != nil {
		return false, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, errors.New("nenhum usuário encontrado")
	}
	return true, nil
}

func (r *UsersRepository) ChangeUserDataDB(name string, userName string, password string, email string, role string) (*Users, error) {

	// Configuração da conexão com o banco de dados

	// Consulta Db para ver se já existe usuário
	userArray := []Users{}
	count, err := r.DB.NewSelect().Model(&userArray).Where("username=?", userName).ScanAndCount(context.Background())
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, errors.New("usuário não encontrado")
	}

	// Caso não exista adicionar ao banco
	User := Users{
		Name:     name,
		UserName: userName,
		Password: password,
		Email:    email,
		Role:     role,
	}
	res, err := r.DB.NewUpdate().Model(&User).OmitZero().Where("userName=?", userName).Returning("*").Exec(context.Background())
	if err != nil {
		return nil, err
	}

	// Return
	fmt.Println(res)
	return &User, nil
}
