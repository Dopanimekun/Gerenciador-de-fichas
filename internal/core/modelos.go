package core

type Raca struct {
	ID             string          `json:"id"`
	Nome           string          `json:"nome"`
	BonusAtributos map[string]int8 `json:"bonus-atributos"`
	Habilidades    []Poder         `json:"poder"`
}

type Classe struct {
	ID       string   `json:"id"`
	Nome     string   `json:"nome"`
	PVnvl    int8     `json:"pv-por-nvl"`
	PMnvl    int8     `json:"pm-por-nvl"`
	Pericias []string `json:"pericias"`
}

type Origem struct {
	ID       string   `json:"id"`
	Nome     string   `json:"nome"`
	Pericias []string `json:"pericias"`
	Poderes  []Poder  `json:"poderes"`
}

type Divindade struct {
	ID      string   `json:"id"`
	Nome    string   `json:"nome"`
	Energia string   `json:"energia"`
	Devotos []string `json:"devotos"`
	Poderes []Poder  `json:"poderes"`
}

type Poder struct {
	ID         string `json:"id"`
	Nome       string `json:"nome"`
	Descricao  string `json:"descricao"`
	Requisitos string `json:"requisitos"`
}

type Equipamento struct {
	ID          string `json:"id"`
	Nome        string `json:"Nome"`
	Descricao   string `json:"Descricao"`
	Requisitos  string `json:"Requisitos"`
	BonusTeste  int8   `json:"Bonus"`
	BonusDano   int8   `json:"bonusdano"`
	BonusDef    int8   `json:"bonusdef"`
	BonusGerais []int8 `json:"bonusgerais"`
	Penalidade  int8   `json:"penalidade"`
}
