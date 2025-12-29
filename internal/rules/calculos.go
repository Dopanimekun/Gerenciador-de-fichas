package rules

import (
	"github.com/Dopanimekun/Gerenciador-de-fichas/internal/core"
)

// calculo utilizando reflect
//func Modificadores(a core.Atributos) map[string]int {
//	v := reflect.ValueOf(a)
//	t := reflect.TypeOf(a)
//
//	mods := make(map[string]int)
//
//	for i := 0; i < v.NumField(); i++ {
//		tag := t.Field(i).Tag.Get("json")
//		valor := int(v.Field(i).Int())
//		mods[tag] = (valor - 10) / 2
//	}
//
//	return mods
//}

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

func CalcularModificador(valorAtributo int8) int8 {
	return int8(valorAtributo-10) / 2
}
