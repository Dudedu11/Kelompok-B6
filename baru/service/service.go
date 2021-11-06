package service

import (
	fmt "fmt"
)

func MLike(name string) string {
	var helloOutput string
	var like int = 100
	if name == "1" {
		like = like + 1
	} else if name == "0" {
		like = like - 1
	}
	helloOutput = fmt.Sprintf("%d", like)
	return helloOutput
}
