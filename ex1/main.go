package main

import "fmt"

// type Repo struct{}

// func GetByID[T any](r Repo, id string) (T, error) {
// }

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
}
