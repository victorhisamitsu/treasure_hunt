package api

import (
	"encoding/json"
	"net/http"
)

func respondError(mensagem string, resposta map[string]any, w http.ResponseWriter) {
	resposta["Sucess"] = false
	resposta["Message"] = mensagem
	respostaByte, _ := json.Marshal(resposta)
	w.Header().Add("Content-Type", "application/json")
	w.Write(respostaByte)
}

func respondSucess(resposta map[string]any, w http.ResponseWriter) {
	respostaByte, err := json.Marshal(resposta)
	if err != nil {
		respondError(err.Error(), resposta, w)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(respostaByte)
}
