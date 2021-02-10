package main

import (
	"testing"
)

type value struct {
	number float64
	primo  bool
}

func TestCheckIfIsPrimo(t *testing.T) {

	v := []value{
		value{
			number: 0,
			primo:  false,
		},
		value{
			number: 1,
			primo:  false,
		},
		value{
			number: 2,
			primo:  true,
		},
		value{
			number: 3,
			primo:  true,
		},
		value{
			number: 4,
			primo:  false,
		},
		value{
			number: 5,
			primo:  true,
		},
		value{
			number: 6,
			primo:  false,
		},
		value{
			number: 7,
			primo:  true,
		},
		value{
			number: 8,
			primo:  false,
		},
		value{
			number: 9,
			primo:  false,
		},
		value{
			number: 10,
			primo:  false,
		},
	}

	for i := range v {

		check := CheckIfIsPrimo(v[i].number)

		if check != v[i].primo {
			t.Error("Se Esperaba un True")
		}

	}

}
