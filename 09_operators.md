# Операторы (Operators)

Операторы сравнения и логические операторы могут быть применены только к базовым типам Го, которые их поддерживают, либо к простым типам, в основе которых находятся базовые типы. Композитные типы не поддерживают операторы, за исключением операторов сравнения (обсудим ниже).

Операторы поределяются с помощью type sets, которые очень похожи на обычные интерфейсы по принципам формирования и пересечения методов.

```go
type Integer interface{ int }

type MyString string

// AnyString matches any type whose underlying type is string.
// This includes, among others, the type string itself, and
// the type MyString.
type AnyString interface {
	~string
}

// ApproximateMyString is INVALID.
type ApproximateMyString interface {
	~MyString // INVALID: underlying type of MyString is not MyString
}

// ApproximateParameter is INVALID.
type ApproximateParameter[T any] interface {
	~T // INVALID: T is a type parameter
}

type PredeclaredSignedInteger interface {
	int | int8 | int16 | int32 | int64
}

// SignedInteger is a constraint that matches any signed integer type.
type SignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

```

Новый пакет constraints:

```go
package constraints

// Ordered is a type constraint that matches any ordered type.
// An ordered type is one that supports the <, <=, >, and >= operators.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// Smallest returns the smallest element in a slice.
// It panics if the slice is empty.
func Smallest[T constraints.Ordered](s []T) T {
	r := s[0] // panics if slice is empty
	for _, v := range s[1:] {
		if v < r {
			r = v
		}
	}
	return r
}
```

Операторы сравнения == и != могут использоваться для struct, array и interface типов.
Для этого предопледелено ограничение comparable:

```go
// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

// ComparableHasher is a type constraint that matches all
// comparable types with a Hash method.
type ComparableHasher interface {
	comparable
	Hash() uintptr
}

// ImpossibleConstraint is a type constraint that no type can satisfy,
// because slice types are not comparable.
type ImpossibleConstraint interface {
	comparable
	[]int
}
```