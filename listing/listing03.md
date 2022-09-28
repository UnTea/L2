Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых
интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil

	return err
}

func main() {
	err := Foo()

	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Вывод:   
nil  
false

`error` представляет собой интерфейс. `Foo` возвращает указатель на структуру `PathError`, которая реализует интерфейс `error`. Получается, что интерфейс не пустой, в нем лежит структура, которая уже содержит `nil`. Поэтому выводится значение `nil`, но интерфейс не является `nil`.

[Внутреннее устройство интерфейса](https://github.com/golang/go/blob/b32689f6c3156da19d469f35cc68fc155d401ef9/src/runtime/runtime2.go#L202-L205):

```go
type iface struct {
    tab  *itab
    data unsafe.Pointer
}
```

Где [tab](https://github.com/golang/go/blob/b32689f6c3156da19d469f35cc68fc155d401ef9/src/runtime/runtime2.go#L925-L931) это указатель на [itab](https://github.com/golang/go/blob/b32689f6c3156da19d469f35cc68fc155d401ef9/src/runtime/type.go#L350-L354) и [unsafe.Pointer](https://github.com/golang/go/blob/b32689f6c3156da19d469f35cc68fc155d401ef9/src/unsafe/unsafe.go#L184):

```go
type itab struct {
    inter *interfacetype
    _type *_type
    hash  uint32 // copy of _type.hash. Used for type switches.
    _     [4]byte
    fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
}
```

```go
type interfacetype struct {
	typ     _type
	pkgpath name
	mhdr    []imethod
}
```

`interfacetype` — структура, которая хранит в себе: [typ](https://github.com/golang/go/blob/b32689f6c3156da19d469f35cc68fc155d401ef9/src/runtime/type.go#L35-L52) тип интерфеса, [pkgpath](https://github.com/golang/go/blob/b32689f6c3156da19d469f35cc68fc155d401ef9/src/runtime/type.go#L430) имя модуля в котором обявлен интерфейс, [mhdr](https://github.com/golang/go/blob/b32689f6c3156da19d469f35cc68fc155d401ef9/src/reflect/type.go#L396) массив объявленных функций внутри интерфейса.:

```go
type Pointer *ArbitraryType
```

`unsafe.Pointer` - его существование обусловлено тем, что над ним можно применять 4 операции, недоступные над другими типами данных:
- Значение указателя любого типа может быть преобразовано в `unsafe.Pointer`.
- `unsafe.Pointer` может быть преобразован в значение указателя любого типа.
- `uintptr` может быть преобразован в `unsafe.Pointer`.
- `unsafe.Pointer` может быть преобразован в `uintptr`.

В данном случае `unsafe.Pointer` указывает на фактическую переменную с конкретным (статическим) типом, пустой интерфейс хранит только `unsafe.Pointer`, так как у него нет методов.
