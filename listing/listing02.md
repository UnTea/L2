Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	
	x = 1
	
	return
}


func anotherTest() int {
	var x int
	
	defer func() {
		x++
	}()
	
	x = 1
	
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:  
Вывод:  
2  
1

Оператор defer откладывает выполнение функции до тех пор, пока не вернется окружающая функция, либо в обычном режиме, 
либо через panic. Аргументы отложенного вызова вычисляются немедленно, даже если вызов функции не выполняется до тех пор
, пока не вернется окружающая функция. Если имеется несколько отложенных вызовов функций, они выполняются в порядке 
"последний пришел - первым вышел".

[В документации описывается следующее](https://go.dev/ref/spec#Defer_statements):
For instance, if the deferred function is a function literal and the surrounding function has named result parameters 
that are in scope within the literal, the deferred function may access and modify the result parameters before they are 
returned. If the deferred function has any return values, they are discarded when the function completes.

В нашем примере произойдёт следующее:

```go
func anotherTest() int {
	var x int // инициализация переменной
	
	defer func() {
		x++ // инкрементирование переменной x, возвращаемое значение функцией не изменяется
	}()
	
	x = 1 // присваивание переменной x значение равное 1
	
	return x // возвращаемое значение переменной будет равно 1
}

func test() (x int) { // x является возвращаемым параметром функции
	defer func() {
		x++ // инкрементирование переменной x, возвращаемое значение функции измениться
	}()
	
	x = 1 // присваивание переменной x значение равное 1
	
	return // возвращаемое значение переменной будет равно 2
}
```
