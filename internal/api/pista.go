package api

import (
	"net/http"
)

type pistaDto struct {
	Nome        string `json:"nome"`
	Texto       string `json:"texto"`
	NomeCaminho string `json:"caminho"`
	Active      bool   `json:"ativo"`
	Dica        string `json:"dica"`
	OrdemPista  int    `json:"ordemPista"`
	Id          string `json:"id"`
}

func (h *ApiHandler) createPistaHandler(w http.ResponseWriter, r *http.Request) {
	// Ler o body
	bodyRequest := &pistaDto{}
	resposta := map[string]any{"Sucess": true}
	err := readBody(r.Body, bodyRequest)
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}
	//Executar minha service
	err = h.pistaService.InsertPistas(bodyRequest.Nome, bodyRequest.Texto, bodyRequest.NomeCaminho, bodyRequest.Active, bodyRequest.Dica, bodyRequest.OrdemPista)
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}

	//Responder
	resposta["Message"] = "Pista cadastrada"
	respondSucess(resposta, w)
}

func (h *ApiHandler) readPistaHandler(w http.ResponseWriter, r *http.Request) {
	// Ler o body
	bodyRequest := &pistaDto{}
	resposta := map[string]any{"Sucess": true}
	err := readBody(r.Body, bodyRequest)
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}
	//Executar minha service
	user, err := h.pistaService.SearchPistas(bodyRequest.Id)
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}
	//Responder
	resposta["Message"] = user
	respondSucess(resposta, w)
}

func (h *ApiHandler) deletePistaHandler(w http.ResponseWriter, r *http.Request) {
	// Ler o body
	bodyRequest := &pistaDto{}
	resposta := map[string]any{"Sucess": true}
	err := readBody(r.Body, bodyRequest)
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}
	//Executar minha service
	delete := h.pistaService.DeletePistas(bodyRequest.Id)
	if !delete {
		respondError(err.Error(), resposta, w)
		return
	}

	//Responder
	resposta["Message"] = "Pista deletada"
	respondSucess(resposta, w)
}

func (h *ApiHandler) searchPistasByCaminhoHandler(w http.ResponseWriter, r *http.Request) {
	bodyRequest := &pistaDto{}
	resposta := map[string]any{"Sucess": true}
	err := readBody(r.Body, bodyRequest)
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}
	//Executar minha service
	pistas, err := h.pistaService.SearchPistasByCaminho(bodyRequest.NomeCaminho)
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}

	//Responder
	resposta["Message"] = pistas
	respondSucess(resposta, w)
}

func (h *ApiHandler) searchAllPistasHandler(w http.ResponseWriter, r *http.Request) {

	resposta := map[string]any{"Sucess": true}

	//Executar minha service
	pistas, err := h.pistaService.SearchAllPistas()
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}

	//Responder
	resposta["Message"] = pistas
	respondSucess(resposta, w)
}

func (h *ApiHandler) changePistaHandler(w http.ResponseWriter, r *http.Request) {
	// Ler o body
	bodyRequest := &pistaDto{}
	resposta := map[string]any{"Sucess": true}
	err := readBody(r.Body, bodyRequest)
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}
	//Executar minha service
	caminho, err := h.pistaService.ChangePista(bodyRequest.Id, bodyRequest.Nome, bodyRequest.Texto, bodyRequest.Active, bodyRequest.Dica, bodyRequest.OrdemPista)
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}

	//Responder
	resposta["Caminho"] = caminho
	resposta["Message"] = "Caminho alterado"
	respondSucess(resposta, w)
}
