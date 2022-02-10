# Ограничения типа (constraints)

Чтобы использовать поведение типа с помощью его методов, нужно явно указать ограничения типа в виде типа-интерфейса:

```go
func Print[T fmt.Stringer](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

```

https://go2goplay.golang.org/p/z_2qAnPOec3 - string does not satisfy fmt.Stringer (missing method String)

