package service

import (
	"github.com/Hitsa/CacaTesouro/internal/repository"
)

type PistaService struct {
	Repository *repository.PistaRepository
}

func NewPistaService(r *repository.PistaRepository) *PistaService {
	pista := PistaService{r}
	return &pista
}

func (s PistaService) InsertPistas(nomePista string, texto string, nomeCaminho string, active bool, dica string, ordemPista int) error {

	err := s.Repository.InsertDbPistas(nomePista, texto, nomeCaminho, active, dica, ordemPista)
	if err != nil {
		return err
	}
	return nil

}

func (s PistaService) SearchPistas(id string) (*repository.Pista, error) {

	// Buscar no Db usuário
	pistas, err := s.Repository.SearchDbPistas(id)
	if err != nil {
		return nil, err
	}
	return pistas, nil
}

func (s PistaService) DeletePistas(id string) bool {

	//Executar query
	valid := s.Repository.DeleteDbPistas(id)
	if !valid {
		return valid
	}
	return true
}

func (s PistaService) SearchPistasByCaminho(nome string) ([]repository.Pista, error) {

	// Buscar no Db usuário
	pistas, err := s.Repository.SearchCaminhoByPistaDB(nome)
	if err != nil {
		return nil, err
	}
	return pistas, nil
}

func (s PistaService) SearchAllPistas() ([]repository.Pista, error) {

	// Buscar no Db usuário
	pistas, err := s.Repository.SearchAllPistasDB()
	if err != nil {
		return nil, err
	}
	return pistas, nil
}

func (s PistaService) ChangePista(id string, nome string, texto string, ativo bool, dica string, ordemPista int) (*repository.Pista, error) {

	pista, err := s.Repository.ChangePistaDb(id, nome, texto, ativo, dica, ordemPista)
	if err != nil {
		return nil, err
	}
	return pista, nil
}
