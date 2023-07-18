package api

import (
	"errors"
	"net/http"
)

type caminhoDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

func (h *ApiHandler) createCaminhoHandler(w http.ResponseWriter, r *http.Request) {
	// Ler o body
	bodyRequest := &caminhoDto{}
	resposta := map[string]any{"Sucess": true}
	err := readBody(r.Body, bodyRequest)
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}
	//Executar minha service
	err = h.caminhoService.InsertCaminho(bodyRequest.Name, bodyRequest.Description, bodyRequest.Active)
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}

	//Responder
	resposta["Message"] = "Caminho cadastrado"
	respondSucess(resposta, w)
}

func (h *ApiHandler) readCaminhoHandler(w http.ResponseWriter, r *http.Request) {
	// Ler o body
	bodyRequest := &createUserDto{}
	resposta := map[string]any{"Sucess": true}
	err := readBody(r.Body, bodyRequest)
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}
	//Executar minha service
	caminho, err := h.caminhoService.SearchCaminho(bodyRequest.Name)
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}

	//Responder
	resposta["Message"] = caminho
	respondSucess(resposta, w)
}

func (h *ApiHandler) deleteCaminhoHandler(w http.ResponseWriter, r *http.Request) {
	// Ler o body
	bodyRequest := &createUserDto{}
	resposta := map[string]any{"Sucess": true}
	err := readBody(r.Body, bodyRequest)
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}
	//Executar minha service
	delete, err := h.caminhoService.DeleteCaminho(bodyRequest.Name)
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}

	if !delete {
		err = errors.New("erro ao deletar caminho")
		respondError(err.Error(), resposta, w)
		return
	}

	//Responder
	resposta["Caminho"] = bodyRequest.Name
	resposta["Message"] = "Caminho deletado"
	respondSucess(resposta, w)
}

func (h *ApiHandler) changeCaminhoHandler(w http.ResponseWriter, r *http.Request) {
	// Ler o body
	bodyRequest := &caminhoDto{}
	resposta := map[string]any{"Sucess": true}
	err := readBody(r.Body, bodyRequest)
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}
	//Executar minha service
	caminho, err := h.caminhoService.ChangeCaminho(bodyRequest.Name, bodyRequest.Description, bodyRequest.Active)
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}

	//Responder
	resposta["Caminho"] = caminho
	resposta["Message"] = "Caminho alterado"
	respondSucess(resposta, w)
}

func (h *ApiHandler) searchAllCaminhosHandler(w http.ResponseWriter, r *http.Request) {
	// Ler o body
	resposta := map[string]any{"Sucess": true}

	//Executar minha service
	caminhos, err := h.caminhoService.SearchAllCaminhos()
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}

	//Responder
	resposta["Message"] = caminhos
	respondSucess(resposta, w)
}
