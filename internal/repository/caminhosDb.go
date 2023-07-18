package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Caminho struct {
	bun.BaseModel `bun:"caminhos"`
	Id            string `json:"id" bun:"id,pk,type:uuid"`
	Name          string `json:"name" bun:"nome"`
	Descrição     string `json:"descrição" bun:"descricao"`
	Active        bool   `json:"active" bun:"ativo"`
	Date          string `json:"date" bun:"data"`
	IdUser        string `json:"idUser" bun:"usuario"`
}

type CaminhoRepository struct {
	DB *bun.DB
}

func NewRepositoryCaminho(d *bun.DB) *CaminhoRepository {
	return &CaminhoRepository{
		DB: d,
	}
}

func (r *CaminhoRepository) InsertDbCaminhoDb(name string, descricao string, active bool) error {

	// Consulta Db para ver se já existe caminho
	caminhos := []Caminho{}
	count, err := r.DB.NewSelect().Model(&caminhos).Where("nome=?", name).ScanAndCount(context.Background())
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("já existe um caminho com esse nome")
	}

	// Caso não exista adicionar ao banco
	dateNow := time.Now()
	dataFormated := dateNow.Format("02/01/2006 15:04:05")

	novoCaminho := Caminho{
		Id:        uuid.NewString(),
		Name:      name,
		Descrição: descricao,
		Active:    active,
		Date:      dataFormated,
	}

	res, err := r.DB.NewInsert().Model(&novoCaminho).Exec(context.Background())
	if err != nil {
		fmt.Println(err)
		return err
	}
	if res == nil {
		return nil
	}

	// Return
	fmt.Println(res)
	return nil

}

func (r *CaminhoRepository) SearchDbCaminho(nome string) (*Caminho, error) {

	// Ligação com db

	// Buscar no Db usuário

	caminhos := Caminho{}
	_, err := r.DB.NewSelect().Model(&caminhos).Where("nome=?", nome).Exec(context.Background(), &caminhos)
	if err != nil {
		return nil, errors.New("caminho não encontrado")
	}

	return &caminhos, nil
}

func (r *CaminhoRepository) DeleteDbCaminho(nome string) (bool, error) {

	//Ligação no db
	userArray := []Caminho{}

	//Executar query
	res, err := r.DB.NewDelete().Model(&userArray).Where("nome=?", nome).Exec(context.Background())
	if err != nil {
		return false, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, errors.New("nenhum caminho encontrado")
	}
	return true, nil
}

func (r CaminhoRepository) ChangeDbCaminho(name string, description string, active bool) (*Caminho, error) {

	// Consulta Db para ver se já existe usuário
	listaCarrinho := []Caminho{}
	count, err := r.DB.NewSelect().Model(&listaCarrinho).Where("nome=?", name).ScanAndCount(context.Background())
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, errors.New("caminho não encontrado")
	}

	// Caso não exista adicionar ao banco
	caminho := Caminho{
		Name:      name,
		Descrição: description,
		Active:    active,
	}
	res, err := r.DB.NewUpdate().Model(&caminho).OmitZero().Where("nome=?", name).Returning("*").Exec(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(res)
	// Return
	fmt.Println("Dados alterados com sucesso.")
	return &caminho, nil
}

func (r *CaminhoRepository) SearchAllCaminhosDb() ([]Caminho, error) {
	listaCaminhos := make([]Caminho, 0)
	_, err := r.DB.NewRaw("SELECT * FROM public.caminhos").Exec(context.Background(), &listaCaminhos)
	if err != nil {
		return nil, errors.New("nenhum caminho cadastrado")
	}

	return listaCaminhos, nil
}

