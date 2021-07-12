package main

import (
	"fmt"
)

type Struct struct {
	s string
}

func (t *Struct) pointerMethod() {
	fmt.Printf("pointerMethod(): type %T, value %v\n", t, t)
}

func (t Struct) valueMethod() {
	fmt.Printf("valueMethod():   type  %T, value  %v\n", t, t)
}

type pointerMethodInterface interface {
	pointerMethod()
}

type valueMethodInterface interface {
	valueMethod()
}

func main() {

	pointerStruct := &Struct{s: "pointerStruct"}
	valueStruct := Struct{s: "valueStruct"}

	// 1
	pointerStruct.pointerMethod()
	pointerStruct.valueMethod()
	valueStruct.pointerMethod()
	valueStruct.valueMethod()

	fmt.Println()

	// 2
	var emptyInterfaceInstancePointerStruct interface{} = pointerStruct
	pp := emptyInterfaceInstancePointerStruct.(pointerMethodInterface)
	pp.pointerMethod()
	pv := emptyInterfaceInstancePointerStruct.(valueMethodInterface)
	pv.valueMethod()

	var emptyInterfaceInstanceValueStruct interface{} = valueStruct
	vp, ok := emptyInterfaceInstanceValueStruct.(pointerMethodInterface)
	if ok {
		vp.pointerMethod()
	} else {
		fmt.Printf("%v\n", vp)
	}
	vv := emptyInterfaceInstanceValueStruct.(valueMethodInterface)
	vv.valueMethod()

	fmt.Println()

	// 3
	valueStruct.s = "NewValue"
	valueStruct.valueMethod()
	vv.valueMethod()

}
