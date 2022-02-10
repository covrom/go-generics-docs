# Выведение типов (Type inference)

Выведение типов делается автоматически, но для более явного читаемого кода можно указывать типы:

```go
func Map[F, T any](s []F, f func(F) T) []T { ... }

var s []int
f := func(i int) int64 { return int64(i) }
var r []int64

// Specify both type arguments explicitly.
r = Map[int, int64](s, f)
// Specify just the first type argument, for F,
// and let T be inferred.
r = Map[int](s, f)
// Don't specify any type arguments, and let both be inferred.
r = Map(s, f)
```

То же самое можно проделывать и с вызовом функций:

```go
Print[int]([]int{1, 2, 3})

s1 := []int{1, 2, 3}
Print(s1)

// Map calls the function f on every element of the slice s,
// returning a new slice of the results.
func Map[F, T any](s []F, f func(F) T) []T {
	r := make([]T, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

strs := Map([]int{1, 2, 3}, strconv.Itoa)

```

# Унификация типов

Типы-параметры соответствуют типам-аргументам различной сложности за счет общей унификации:

```go
[]map[int]bool

T1 (T1 matches []map[int]bool)
[]T1 (T1 matches map[int]bool)
[]map[T1]T2 (T1 matches int, T2 matches bool)

cannot be unified with:
    int
    struct{}
    []struct{}
    []map[T1]string
```
