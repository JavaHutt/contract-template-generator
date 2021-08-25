package main

import (
	"log"

	"github.com/nguyenthenguyen/docx"
)

func main() {
	r, err := docx.ReadDocxFile("./template.docx")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	docx1 := r.Editable()
	// Replace like https://golang.org/pkg/strings/#Replace
	if err := docx1.Replace("общая сумма контракта", "8 000 000 (восемь миллионов) Долларов США", -1); err != nil {
		log.Fatal(err)
	}
	docx1.WriteToFile("./result.docx")
}
