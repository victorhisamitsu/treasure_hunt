package api

import (
	"errors"
	"fmt"
	"net/http"
)

type createUserDto struct {
	Name     string `json:"name"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func (h *ApiHandler) createUserHandler(w http.ResponseWriter, r *http.Request) {
	// Ler o body
	bodyRequest := &createUserDto{}
	resposta := map[string]any{"Sucess": true}
	err := readBody(r.Body, bodyRequest)
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}
	//Executar minha service
	err = h.userService.CreateUser(bodyRequest.Name, bodyRequest.UserName, bodyRequest.Password, bodyRequest.Email, bodyRequest.Role)
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}

	//Responder
	resposta["Message"] = "usuário cadastrado"
	respondSucess(resposta, w)
}

func (h *ApiHandler) readUserHandler(w http.ResponseWriter, r *http.Request) {
	// Ler o body
	bodyRequest := &createUserDto{}
	resposta := map[string]any{"Sucess": true}
	err := readBody(r.Body, bodyRequest)
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}
	//Executar minha service
	user, err := h.userService.SearchUser(bodyRequest.UserName)
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}
	fmt.Println(*user)
	//Responder
	resposta["Message"] = *user
	respondSucess(resposta, w)
}

func (h *ApiHandler) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Ler o body
	bodyRequest := &createUserDto{}
	resposta := map[string]any{"Sucess": true}
	err := readBody(r.Body, bodyRequest)
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}
	//Executar minha service
	delete, err := h.userService.DeleteUser(bodyRequest.UserName)
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}
	if !delete {
		err = errors.New("usuário não encontrado")
		respondError(err.Error(), resposta, w)
		return
	}

	//Responder
	resposta["UserName"] = bodyRequest.UserName
	resposta["Message"] = "usuário deletado"
	respondSucess(resposta, w)
}

func (h *ApiHandler) changeUserDataHandlers(w http.ResponseWriter, r *http.Request) {
	// Ler o body
	bodyRequest := &createUserDto{}
	resposta := map[string]any{"Sucess": true}
	err := readBody(r.Body, bodyRequest)
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}
	//Executar minha service
	user, err := h.userService.ChangeUserData(bodyRequest.Name, bodyRequest.UserName, bodyRequest.Password, bodyRequest.Email, bodyRequest.Role)
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}
	if user == nil {
		respondError(err.Error(), resposta, w)
		return
	}
	fmt.Println(*user)
	//Responder
	resposta["Message"] = *user
	respondSucess(resposta, w)
}
