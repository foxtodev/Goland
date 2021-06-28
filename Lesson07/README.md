##### С какими интерфейсами мы уже сталкивались в предыдущих уроках? Обратите внимание на уроки, в которых мы читали из стандартного ввода и писали в стандартный вывод.

```go
func fmt.Println(a ...interface{}) (n int, err error)
func fmt.Scanf(format string, a ...interface{}) (n int, err error)
```
##### Посмотрите примеры кода в своём портфолио. Везде ли ошибки обрабатываются грамотно? Хотите ли вы переписать какие-либо функции?

Можно было бы сделать обработку ошибок ввода, например

```go
n, err = fmt.Scanf("%f\n", &value)
if err != nil || n != 1 {
    // invalid input
    fmt.Println(n, err)
}
```


##### Дополнительное практическое задание

```go
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

```

```
pixel@pixel-machine:~/go/Goland$ go run Lesson07/main.go
pointerMethod(): type *main.Struct, value &{pointerStruct}
valueMethod():   type  main.Struct, value  {pointerStruct}
pointerMethod(): type *main.Struct, value &{valueStruct}
valueMethod():   type  main.Struct, value  {valueStruct}

pointerMethod(): type *main.Struct, value &{pointerStruct}
valueMethod():   type  main.Struct, value  {pointerStruct}
<nil>
valueMethod():   type  main.Struct, value  {valueStruct}

valueMethod():   type  main.Struct, value  {NewValue}
valueMethod():   type  main.Struct, value  {valueStruct}
```

<br />