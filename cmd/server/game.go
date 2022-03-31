package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"server/interconnect"
	"server/util"
	"strings"
)

type LocalImpl struct {
	interconnect.UnimplementedObserverServer

	word string
	fuel uint32
}

func NewLocalImpl(fuel uint32, word string) *LocalImpl {
	return &LocalImpl{
		fuel: fuel,
		word: word,
	}
}

func (l *LocalImpl) ProcessGame(server interconnect.Observer_ProcessGameServer) error {
	game := l.NewGameInstance(server)
	game.server = server
	return game.PlayImpl()
}

func (l *LocalImpl) Play() {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	interconnect.RegisterObserverServer(s, l)

	log.Printf("Starting Game server at port %v\n", port)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (l *LocalImpl) NewGameInstance(server interconnect.Observer_ProcessGameServer) *GameInstance {
	return &GameInstance{
		fuel:   l.fuel,
		word:   util.NewSecretString(l.word),
		server: server,
	}
}

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
