package main

import (
	"github.com/pkg/errors"
	"log"
)

func main() {
	if err := prenda(17); err != nil {
		log.Fatal(err)
	}
}

func prenda(idade int) error {
	if idade < 18 {
		return errors.New("prenda: menores de 18 não podem ser presos")
	}
	return nil
}
