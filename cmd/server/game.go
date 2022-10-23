package main

import (
	"fmt"
	"strings"

	"github.com/kirillgrachoff/MysteryField/pkg/interconnect"
	"github.com/kirillgrachoff/MysteryField/pkg/util"
)

type GameInstance struct {
	server interconnect.Observer_ProcessGameServer
	word   *util.SecretString
	fuel   uint32
}

func (game *GameInstance) PlayImpl() error {
	for ; game.fuel > 0 && !game.Won(); game.fuel -= 1 {
		err := game.SendGameState()
		if err != nil {
			return err
		}

		word, err := game.Receive()
		if err != nil {
			return err
		}

		err = game.GuessCharacter(word)
		if err != nil {
			return err
		}
	}
	var err error
	err = nil
	if game.fuel == 0 {
		err = game.SendDefeat()
	} else {
		err = game.SendWin()
	}
	return err
}

func (game *GameInstance) SendGameState() error {
	err := game.server.Send(&interconnect.WordStatus{
		GameState:  game.word.Get(),
		GameStatus: interconnect.WordStatus_RUNNING,
		Fuel:       game.fuel,
	})
	return err
}

func (game *GameInstance) SendWin() error {
	return game.server.Send(&interconnect.WordStatus{
		GameState:  game.word.Get(),
		GameStatus: interconnect.WordStatus_WIN,
		Fuel:       game.fuel,
	})
}

func (game *GameInstance) SendDefeat() error {
	return game.server.Send(&interconnect.WordStatus{
		GameState:  game.word.Get(),
		GameStatus: interconnect.WordStatus_DEFEAT,
		Fuel:       game.fuel,
	})
}

func (game *GameInstance) Receive() (*interconnect.WordGuess, error) {
	return game.server.Recv()
}

func (game *GameInstance) GuessCharacter(word *interconnect.WordGuess) error {
	if len(word.Character) > 1 {
		return fmt.Errorf("input: %v with len %v\n", word.Character, len(word.Character))
	}
	game.word.Guess([]rune(word.Character)[0])
	return nil
}

func (game *GameInstance) Won() bool {
	return !strings.Contains(game.word.Get(), "*")
}
