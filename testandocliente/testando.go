package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Nextc3/notificacao-covid-blockchain/entidade"
	"log"

	"github.com/Nextc3/fabric-samples/testandocliente/cliente"
)

func main() {
	fmt.Println("Iniciando teste do cliente")

	var conex cliente.Conexao
	var contra cliente.Contrato
	aux, gw := conex.IniciarConexao()
	defer gw.Close()

	if aux == nil && gw == nil {
		log.Fatalf("Falha em começar uma conexão. No método principal")
	}
	contra.SetContrato(aux)
	/*
		fmt.Println("Todas as notificações :")

		notis, err := contra.ObterTodasNotificacoes(false)
		if err != nil {
			log.Fatalf("Erro em obter todas as notificações")
		}
		for _, resultado := range notis {
			notiponti := resultado.Ativo

			imprimirNotificacao(*notiponti)

		}
	*/

	noti, err := contra.ConsultarNotificacao(false, 1)
	if err != nil {
		log.Fatalf("Erro em obter resposta do contrato")
	}
	fmt.Println("Notificação completa")
	imprimirNotificacao(noti)

	fmt.Println("Existe notificação 2?")
	resposta := contra.ExisteNotificacao(false, 2)
	fmt.Println(resposta)

	fmt.Println("Todas as notificações :")

	notis, err := contra.ObterTodasNotificacoes(false)
	if err != nil {
		log.Fatalf("Erro em obter todas as notificações")
	}
	for _, resultado := range notis {
		notiponti := resultado.Ativo

		imprimirNotificacao(*notiponti)

	}
	fmt.Println("Teste de criar notificação com Bolsonaro")
	contra.CriarNotificacao(criarNotificacao())
	fmt.Println("Existe notificação 3?")
	resposta = contra.ExisteNotificacao(false, 3)
	fmt.Println(resposta)
	fmt.Println("Consultando notificação de Bolsonaro")
	noti, err = contra.ConsultarNotificacao(false, 3)
	if err != nil {
		log.Fatalf("Erro em obter resposta do contrato")
	}
	fmt.Println("Notificação completa de Bolsonaro")
	imprimirNotificacao(noti)
	fmt.Println("Imprimindo todas as notificações com a de Bolsonaro")
	notis, err = contra.ObterTodasNotificacoes(false)
	if err != nil {
		log.Fatalf("Erro em obter todas as notificações")
	}
	for _, resultado := range notis {
		notiponti := resultado.Ativo

		imprimirNotificacao(*notiponti)

	}

}
func formatarBonitoJSON(data []byte) string {
	var bonito bytes.Buffer
	if err := json.Indent(&bonito, data, " ", ""); err != nil {
		panic(fmt.Errorf("Falhou em fomartar JSON: %w", err))
	}
	return bonito.String()
}
func imprimirNotificacao(n entidade.Notificacao) {
	notiEmBytes, _ := json.Marshal(n)
	fmt.Println(formatarBonitoJSON(notiEmBytes))

}
func criarNotificacao() entidade.Notificacao {
	return entidade.Notificacao{
		Id: 3,
		CidadaoNotificador: entidade.Notificador{
			Id:             2,
			Email:          "reydeplastico@gmail.com",
			Cpf:            "741.852.963-01",
			DataNascimento: "01/10/1961",
			Nome:           "Dr. Rey",
			NomeDaMae:      "Roberta Miguel Rey",
			Estado:         "São Paulo",
			Municipio:      "São Paulo",
			Telefone:       "11 1234-4567",
			Ocupacao:       "Médico",
		},
		TemCPF:                    true,
		EhProfissionalDeSaude:     false,
		EhProfissionalDeSeguranca: false,
		Cpf:                       "951.753.654-17",
		Ocupacao:                  "Presidente",
		Nome:                      "Jair Messias Bolsonaro",
		DataNascimento:            "21/03/1955",
		Sexo:                      false,
		Raca:                      1,
		PovoTradicional:           false,
		Cep:                       "70150-900",
		Logradouro:                "Zona Cívico-Administrativa",
		NumeroEndereco:            "s/n",
		Complemento:               "Palácio do Planalto",
		Bairro:                    "PRAÇA DOS TRÊS PODERES",
		Estado:                    "Distrito Federal",
		Municipio:                 "Brasília",
		Telefone:                  "(61) 3411-1065",
		Email:                     "jairbolsonaro@gmail.com",
		Estrategia:                1,
		LocalizacaoTeste:          2,
		DataNotificacao:           "06/07/2021",
		Condicoes: map[string]bool{
			"bolsa de colostomia": true,
		},
		Sintomas: map[string]bool{
			"febre":            true,
			"oxigenação baixa": true,
		},
		Vacinas: map[string]bool{
			"nenhuma": true,
		},
		Teste: []entidade.TesteCovid{
			{
				Id:            3,
				TipoDeTeste:   "rt-pcr",
				EstadoDoTeste: 3,
				DataDaColeta:  "06/07/2021",
				Resultado:     1,
				Lote:          "1234",
				Fabricante:    "Fiocruz",
			},
		},
	}

}
