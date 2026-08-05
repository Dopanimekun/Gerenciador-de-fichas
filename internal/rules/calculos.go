package rules

import (
	"github.com/Dopanimekun/Gerenciador-de-fichas/internal/core"
)

func CalcularModificador(valorAtributo int8) int8 {
	modificador := (valorAtributo - 10) / 2
	return modificador
}

func PVinicial(modCON *core.ModATR, pvClasse *core.Classe) int8 {
	pvInicial := pvClasse.PVnvl + modCON.ModCON
	return pvInicial
}

func CalcularDefesa(modDES *core.ModATR, bonusDef *core.Equipamento) int8 {
	Defesa := 10 + modDES.ModDES + bonusDef.BonusDef
	return Defesa
}

func AtributosToMap(a core.Atributos) map[string]int8 {
	return map[string]int8{
		"for": int8(a.FOR),
		"des": int8(a.DES),
		"con": int8(a.CON),
		"int": int8(a.INT),
		"sab": int8(a.SAB),
		"car": int8(a.CAR),
	}
}

func MapToAtributos(m map[string]int8) core.Atributos {
	return core.Atributos{
		FOR: int8(m["for"]),
		DES: int8(m["des"]),
		CON: int8(m["con"]),
		INT: int8(m["int"]),
		SAB: int8(m["sab"]),
		CAR: int8(m["car"]),
	}
}
