# Обобщенные типы (Generic types)

```go
type Vector[T any] []T
```

Внутри определения типа, тип-параметр может использоваться везде и теми же способами, что и любой другой тип.

```go
func (v *Vector[T]) Push(x T) { *v = append(*v, x) }

var v Vector[int]

type List[T any] struct {
	next *List[T] // this reference to List[T] is OK
	val  T
}

type P[T1, T2 any] struct {
	F *P[T1, T2] // [T2, T1] is invalid!
}

```

Ограничения так же применимы, как и для типов-параметров функций:

```go
type StringableVector[T Stringer] []T

func (s StringableVector[T]) String() string {
	var sb strings.Builder
	for i, v := range s {
		if i > 0 {
			sb.WriteString(", ")
		}
		// It's OK to call v.String here because v is of type T
		// and T's constraint is Stringer.
		sb.WriteString(v.String())
	}
	return sb.String()
}
```

https://go2goplay.golang.org/p/xLpLFI22K_E
https://godbolt.org/z/1qqfhYMEv

