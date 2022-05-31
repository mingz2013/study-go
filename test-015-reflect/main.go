package main

import (
	"log"
	"reflect"
)

type Calc struct {
	a int
	b int
}

func (cale *Calc) Add(a int, b int) int {
	return a + b
}

func main() {
	calc := new(Calc)
	ty := reflect.TypeOf(calc)
	log.Println(ty)
	v := reflect.ValueOf(calc)
	log.Println(v)
	dire := reflect.Indirect(v)
	log.Println(dire)
	dire.Type().Name()

}
