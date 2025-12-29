package core

type Atributos struct {
	FOR int8 `json:"for"`
	DES int8 `json:"des"`
	CON int8 `json:"con"`
	INT int8 `json:"int"`
	SAB int8 `json:"sab"`
	CAR int8 `json:"car"`
}

type ModATR struct {
	ModFOR int8 `json:"modfor"`
	ModDES int8 `json:"moddes"`
	ModCON int8 `json:"modcon"`
	ModINT int8 `json:"modint"`
	ModSAB int8 `json:"modsab"`
	ModCAR int8 `json:"modcar"`
}

type Pericia struct {
	Nome               string `json:"nome"`
	AtributoChave      string `json:"atributo"`
	SomenteTreinada    bool   `json:"somente-treinada"`
	PenalidadeArmadura bool   `json:"penalidade-armadura"`
}

type PericiaTreinada struct {
	Nome  string `json:"nome"`
	Bonus int8   `json:"bonus"`
}
