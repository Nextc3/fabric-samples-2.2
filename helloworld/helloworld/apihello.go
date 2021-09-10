package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Oi struct {
	Id        string `json:"id,omitempty"`
	Saudacao  string `json:"saudacao,omitempty"`
	Despedida string `json:"despedida,omitempty"`
	Oidenovo  string `json:"oidenovo,omitempty"`
	Pessoa    string `json:"pessoa,omitempty"`
}

var ois []Oi

// GetOis mostra todos os contatos da variável ois
func GetOis(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ois)
}

// GetOi mostra apenas um contato
func GetOi(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range ois {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Oi{})
}

// CreatePerson cria um novo contato
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var oi Oi
	_ = json.NewDecoder(r.Body).Decode(&oi)
	oi.Id = params["id"]
	ois = append(ois, oi)
	json.NewEncoder(w).Encode(ois)
}

// DeletePerson deleta um contato
func DeleteOi(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range ois {
		if item.Id == params["id"] {
			ois = append(ois[:index], ois[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(ois)
	}
}

// função principal para executar a api
func main() {

	router := mux.NewRouter()
	ois = append(ois, Oi{Id: "OI1", Saudacao: "e ai fdp", Despedida: "xau fdp", Oidenovo: "que viagem", Pessoa: "El Diablo"})
	ois = append(ois, Oi{Id: "OI2", Saudacao: "arrombado vc", Despedida: "sai sacana", Oidenovo: "de novo", Pessoa: "El Putaria"})

	router.HandleFunc("/contato", GetOis).Methods("GET")
	router.HandleFunc("/contato/{id}", GetOi).Methods("GET")
	router.HandleFunc("/contato/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/contato/{id}", DeleteOi).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
