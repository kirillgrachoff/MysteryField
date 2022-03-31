package main

import (
	"log"
	"server/util"
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
	port = util.GetOrDefault("PORT", defaultPort)

	secretWord = util.GetOrDefault("SECRET_WORD", defaultSecretWord)

	fuel := util.GetOrDefault("FUEL_INIT", defaultFuelInit)

	value, err := strconv.Atoi(fuel)
	if err != nil {
		log.Fatalln(err)
	}

	fuelInit = uint32(value + len(util.UniqueCharacters(secretWord)))
}

func main() {
	srv := NewObserverImpl(fuelInit, secretWord)
	srv.ServeGrpc()
}
