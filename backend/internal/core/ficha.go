// Package core é responsável por definir os modelos
// de ficha e outros campos relevantes
package core

import (
	"github.com/google/uuid"
)

type Ficha struct {
	ID        string    `json:"id"`
	Nome      string    `json:"nome"`
	Raca      Raca      `json:"raca"`
	Classe    Classe    `json:"classe"`
	Nivel     uint8     `json:"nivel"`
	Origem    Origem    `json:"origem"`
	Divindade Divindade `json:"divindade"`
	Atributos Atributos `json:"atributos"`
	Mod       ModATR    `json:"mod"`
	PVTotal   int16     `json:"pvtotal"`
	PVAtual   int16     `json:"pvatual"`
	PMTotal   int16     `json:"pmtotal"`
	PMAtual   int16     `json:"pmatual"`
	Defesa    int8      `json:"defesa"`
}

func FichaLimpa(nome string, raca Raca, classe Classe, origem Origem, divindade Divindade) *Ficha {
	return &Ficha{
		uuid.NewString(),
		nome,
		raca,
		classe,
		1,
		origem,
		divindade,
		Atributos{10, 10, 10, 10, 10, 10},
		ModATR{0, 0, 0, 0, 0, 0},
		0,
		0,
		0,
		0,
		10,
	}
}
