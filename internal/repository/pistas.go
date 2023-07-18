package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Pista struct {
	bun.BaseModel `bun:"pista"`
	Id            string    `json:"id" bun:"id,pk,type:uuid"`
	Nome          string    `json:"nome" bun:"nome"`
	Texto         string    `json:"texto" bun:"texto"`
	IdCaminho     string    `json:"idCaminho" bun:"idcaminho"`
	NomeCaminho   string    `json:"nomeCaminho" bun:"caminho"`
	CreatedDate   time.Time `json:"date" bun:"datadica"`
	NextDatePista time.Time `json:"nextDatePista" bun:"dataproximadica"`
	Active        bool      `json:"active" bun:"ativo"`
	User          string    `json:"user" bun:"usuario"`
	Dica          string    `json:"dica" bun:"dica"`
	OrdemPista    int       `json:"ordemPista" bun:"ordempista"`
}

type PistaRepository struct {
	DB *bun.DB
}

func NewRepositoryPista(d *bun.DB) *PistaRepository {
	return &PistaRepository{
		DB: d,
	}
}

func (r *PistaRepository) InsertDbPistas(nomePista string, texto string, nomeCaminho string, active bool, dica string, ordemPista int) error {

	// Verifica se existe o caminho
	caminhos := Caminho{}
	_, err := r.DB.NewSelect().Model(&caminhos).Where("nome=?", nomeCaminho).Exec(context.Background(), &caminhos)
	if err != nil {
		return errors.New("caminho não encontrado")
	}

	// Verifica se já existe a pista
	pista := []Pista{}
	count, err := r.DB.NewSelect().Model(&pista).Where("caminho=? AND ordempista=?", nomeCaminho, ordemPista).ScanAndCount(context.Background())
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("já existe uma pista nesta ordem")
	}

	// Caso não exista adicionar ao banco
	dateNow := time.Now()
	afterOneHour := dateNow.Add(time.Minute * 60)

	novoCaminho := Pista{
		Id:            uuid.NewString(),
		Nome:          nomePista,
		Texto:         texto,
		IdCaminho:     caminhos.Id,
		NomeCaminho:   nomeCaminho,
		CreatedDate:   dateNow,
		NextDatePista: afterOneHour,
		Active:        true,
		Dica:          dica,
		OrdemPista:    ordemPista,
	}

	res, err := r.DB.NewInsert().Model(&novoCaminho).Exec(context.Background())
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Return
	fmt.Println(res)
	fmt.Println("Pista criada com sucesso.")
	return nil

}

func (r *PistaRepository) SearchDbPistas(id string) (*Pista, error) {

	// Buscar no Db usuário

	pistas := Pista{}
	res, err := r.DB.NewSelect().Model(&pistas).Where("id=?", id).Exec(context.Background(), &pistas)
	if err != nil {
		return nil, errors.New("pista não encontrada")
	}
	resultado, _ := res.RowsAffected()

	if resultado == 0 {
		return nil, errors.New("pista não encontrada")
	}
	return &pistas, nil
}

func (r *PistaRepository) DeleteDbPistas(nome string) bool {

	userArray := []Pista{}

	//Executar query
	err := r.DB.NewDelete().Model(&userArray).Where("id=?", nome).Scan(context.Background())
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (r *PistaRepository) SearchCaminhoByPistaDB(nomeCaminho string) ([]Pista, error) {

	listaPistas := make([]Pista, 0)
	err := r.DB.NewRaw("SELECT * FROM public.pista WHERE caminho = ?", nomeCaminho).Scan(context.Background(), &listaPistas)
	if err != nil {
		return nil, err
	}
	resultados := len(listaPistas)
	fmt.Println(resultados)
	if resultados < 1 {
		return nil, errors.New("nenhum pista encontrada")
	}
	return listaPistas, nil
}

func (r *PistaRepository) SearchAllPistasDB() ([]Pista, error) {

	listaPistas := make([]Pista, 0)
	_, err := r.DB.NewRaw("SELECT * FROM public.pista").Exec(context.Background(), &listaPistas)
	if err != nil {
		return nil, err
	}
	resultados := len(listaPistas)
	fmt.Println(resultados)
	if resultados < 1 {
		return nil, errors.New("nenhum pista encontrada")
	}
	return listaPistas, nil
}

func (r *PistaRepository) ChangePistaDb(id string, nome string, texto string, ativo bool, dica string, ordemPista int) (*Pista, error) {

	// Consulta Db para ver se existe caminho
	pistas := []Pista{}
	count, err := r.DB.NewSelect().Model(&pistas).Where("id=?", id).ScanAndCount(context.Background())
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, errors.New("pista não encontrada")
	}

	// Caso não exista adicionar ao banco
	pista := Pista{
		Nome:       nome,
		Texto:      texto,
		Active:     ativo,
		Dica:       dica,
		OrdemPista: ordemPista,
	}
	res, err := r.DB.NewUpdate().Model(&pista).OmitZero().Where("id=?", id).Returning("*").Exec(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Return
	fmt.Println(res)
	fmt.Println("Dados alterados com sucesso.")
	return &pista, nil
}
