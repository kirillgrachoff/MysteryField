package main

import (
	"log"
	"strconv"

	"github.com/kirillgrachoff/MysteryField/pkg/util"
)

const (
	defaultPort       = "5000"
	defaultSecretWord = "banana"
	defaultFuelBudget = "10"
)

var (
	port       string
	fuelInit   uint32
	secretWord string
)

func init() {
	port = util.GetOrDefault("PORT", defaultPort)

	secretWord = util.GetOrDefault("SECRET_WORD", defaultSecretWord)

	fuel := util.GetOrDefault("FUEL_BUDGET", defaultFuelBudget)

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
