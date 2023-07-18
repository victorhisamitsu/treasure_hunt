package service

import (
	"fmt"

	"github.com/Hitsa/CacaTesouro/internal/repository"
)

type CaminhoService struct {
	Repository *repository.CaminhoRepository
}

func NewCaminhoService(r *repository.CaminhoRepository) *CaminhoService {
	caminho := CaminhoService{r}
	return &caminho
}

func (s CaminhoService) InsertCaminho(name string, descricao string, active bool) error {

	err := s.Repository.InsertDbCaminhoDb(name, descricao, active)
	if err != nil {
		return err
	}
	// Return
	fmt.Println("Caminho criado com sucesso.")
	return nil

}

func (s CaminhoService) SearchCaminho(nome string) (*repository.Caminho, error) {

	// Buscar no Db usuário

	caminhos, err := s.Repository.SearchDbCaminho(nome)
	if err != nil {
		return nil, err
	}

	return caminhos, nil
}

func (s CaminhoService) DeleteCaminho(nome string) (bool, error) {

	valid, err := s.Repository.DeleteDbCaminho(nome)
	if !valid {
		return valid, err
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s CaminhoService) ChangeCaminho(name string, description string, active bool) (*repository.Caminho, error) {

	caminho, err := s.Repository.ChangeDbCaminho(name, description, active)
	if err != nil {
		return nil, err
	}
	return caminho, nil
}

func (s CaminhoService) SearchAllCaminhos() ([]repository.Caminho, error) {

	listaCaminhos, err := s.Repository.SearchAllCaminhosDb()
	if err != nil {
		return nil, err
	}
	return listaCaminhos, nil
}
