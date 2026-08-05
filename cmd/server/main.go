package main

import (
	"encoding/json"
	"fmt"
	"os"

	// IMPORTANTE: Substitua NOME_DO_SEU_MODULO pelo nome que está no seu go.mod
	// Exemplo: se o go.mod diz "module projetorpg", coloque "projetorpg/internal/core"
	"github.com/Dopanimekun/Gerenciador-de-fichas/internal/core"
)

func main() {
	// 1. Instanciamos tipos vazios para passar para a função
	// (Assumindo que Raca, Classe, Origem e Divindade estão definidos em internal/core/modelos.go)
	raca := core.Raca{}
	classe := core.Classe{}
	origem := core.Origem{}
	divindade := core.Divindade{}

	// 2. Chamamos a função de criação de ficha
	ficha := core.FichaLimpa("Teste", raca, classe, origem, divindade)

	// 3. Convertendo a struct para formato JSON com recuos (pretty print)
	jsonData, err := json.MarshalIndent(ficha, "", "  ")
	if err != nil {
		fmt.Printf("Erro ao gerar JSON: %v\n", err)
		return
	}

	// 4. Exibe o JSON no terminal para visualização rápida
	fmt.Println("--- JSON GERADO ---")
	fmt.Println(string(jsonData))
	fmt.Println("-------------------")

	// 5. Salva o JSON em um arquivo na raiz de onde o comando for executado
	err = os.WriteFile("ficha_teste.json", jsonData, 0o644)
	if err != nil {
		fmt.Printf("Erro ao salvar arquivo JSON: %v\n", err)
		return
	}

	fmt.Println("Arquivo 'ficha_teste.json' salvo com sucesso!")
}
