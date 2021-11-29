package service

import (
	
	fmt "fmt"
	"strconv"
)


func Mlike(name string, jml_like string) string {
	var helloOutput string
	var like, err = strconv.Atoi(jml_like)
	if name == "1" {
		like = like + 1
	} else if name == "0" {
		like = like - 1
	}
	if err == nil {
		fmt.Println(like) // 124
	}
	helloOutput = fmt.Sprintf("%d", like)
	return helloOutput
} 
