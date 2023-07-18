package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Hitsa/CacaTesouro/internal/service"
	"github.com/go-chi/chi/v5"
)

type ApiHandler struct {
	caminhoService *service.CaminhoService
	userService    *service.UserService
	pistaService   *service.PistaService
}

func API(caminhoService *service.CaminhoService, userService *service.UserService, pistaService *service.PistaService) {
	r := chi.NewRouter()

	handler := &ApiHandler{
		caminhoService,
		userService,
		pistaService,
	}

	r.Post("/", bodyApiHandler)
	r.Post("/createUser", handler.createUserHandler)
	r.Post("/readUser", handler.readUserHandler)
	r.Post("/deleteUser", handler.deleteUserHandler)
	r.Post("/changeUser", handler.changeUserDataHandlers)
	r.Post("/createCaminho", handler.createCaminhoHandler)
	r.Post("/readCaminho", handler.readCaminhoHandler)
	r.Post("/deleteCaminho", handler.deleteCaminhoHandler)
	r.Post("/changeCaminho", handler.changeCaminhoHandler)
	r.Post("/searchAllCaminhos", handler.searchAllCaminhosHandler)
	r.Post("/createPista", handler.createPistaHandler)
	r.Post("/readPista", handler.readPistaHandler)
	r.Post("/deletePista", handler.deletePistaHandler)
	r.Post("/changePista", handler.changePistaHandler)
	r.Post("/searchAllPistas", handler.searchAllPistasHandler)
	r.Post("/searchPistasByCaminho", handler.searchPistasByCaminhoHandler)
	http.ListenAndServe(":4000", r)
}

func bodyApiHandler(w http.ResponseWriter, r *http.Request) {
	resposta := map[string]any{}
	body, _ := io.ReadAll(r.Body)
	var minhaVariavel map[string]string
	json.Unmarshal(body, &minhaVariavel)
	variavelJson, _ := json.Marshal(resposta)

	w.Write(variavelJson)
}

func readBody[T any](b io.ReadCloser, bodyRequest *T) error {
	body, err := io.ReadAll(b)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, bodyRequest)
	if err != nil {
		return err
	}
	return nil
}
