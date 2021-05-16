package main

import (
	log "fmt"
	"strings"
)

// type slice struct {
// 	ptr *[CAP]interface{}
// 	len uint
// 	cap uint
// }

func main() {
	alimentos := make([]interface{}, 0, 20)

	printTitle("SLICE COM CAPACIDADE INICIAL DEFINIDA")

	appendAndPrint(&alimentos, "arroz")
	appendAndPrint(&alimentos, "feijao")
	appendAndPrint(&alimentos, "macarrao")
	appendAndPrint(&alimentos, "cafe")
	appendAndPrint(&alimentos, "acucar")

	printTitle("SLICE SEM CAPACIDADE INICIAL")

	carros := []interface{}{}

	appendAndPrint(&carros, "corsa")
	appendAndPrint(&carros, "etios")
	appendAndPrint(&carros, "compass")
	appendAndPrint(&carros, "kicks")
	appendAndPrint(&carros, "gol")
}

func appendAndPrint(slice *[]interface{}, appending ...interface{}) {
	*slice = append(*slice, appending...)

	printSlice(slice)
}

func printTitle(title string) {
	width := 80
	line := strings.Repeat("#", width)

	marging := 4

	fillLength := int((width - marging - len(title)) / 2)
	fill := ""

	if fillLength > 0 {
		fill = strings.Repeat("#", fillLength)
	}

	marginString := strings.Repeat(" ", marging)

	headtitle := fill + marginString + title + marginString + fill

	log.Println(line)
	log.Println(headtitle)
	log.Println(line)
}

func printSlice(slice *[]interface{}) {
	log.Println(*slice)

	for index := range *slice {
		valor := &(*slice)[index]

		log.Printf("%d: %v, %p\n", index, *valor, valor)
	}

	log.Printf("len: %d, cap: %d\n", len(*slice), cap(*slice))
	log.Println("------------------------")
}
