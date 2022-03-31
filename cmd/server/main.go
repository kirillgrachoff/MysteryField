package main

import (
	"log"
	"os"
	"strconv"
)

const (
	defaultPort       = "5000"
	defaultSecretWord = "banana"
	defaultFuelInit   = "10"
)

var (
	port       string
	fuelInit   uint32
	secretWord string
)

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	secretWord = os.Getenv("SECRET_WORD")
	if secretWord == "" {
		secretWord = defaultSecretWord
	}

	fuel := os.Getenv("FUEL_INIT")
	if fuel == "" {
		fuel = defaultFuelInit
	}
	value, err := strconv.Atoi(fuel)
	if err != nil {
		log.Fatalln(err)
	}
	fuelInit = uint32(value + len(secretWord))
}

func main() {
	srv := NewLocalImpl(fuelInit, secretWord)
	srv.Play()
}
