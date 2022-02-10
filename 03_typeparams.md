# Type parameters
Тип-параметр - абстрактный тип, который при компиляции заменяется на тип-аргумент (type arguments).
Определяется рядом с именем функции в квадратных скобках, и затем может использоваться где угодно дальше в этой функции - в параметрах, в возврате и в теле функции.

https://go2goplay.golang.org/p/MFQPvbbyvEk

```go
func Print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func UsePrint(){
    Print[int]([]int{1, 2, 3})
    Print([]string{"Hello", "Generics"})
}
```

Смотрим дизассемблер:
https://godbolt.org/z/TGbsMefed - генерятся функции, но не инлайнятся! Пока нет никаких оптимизаций в этой части.

