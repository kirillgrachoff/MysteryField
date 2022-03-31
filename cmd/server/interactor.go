package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"server/interconnect"
	"server/util"
)

type ObserverImpl struct {
	interconnect.UnimplementedObserverServer

	word string
	fuel uint32
}

func NewLocalImpl(fuel uint32, word string) *ObserverImpl {
	return &ObserverImpl{
		fuel: fuel,
		word: word,
	}
}

func (impl *ObserverImpl) ProcessGame(server interconnect.Observer_ProcessGameServer) error {
	game := impl.NewGameInstance(server)
	game.server = server
	return game.PlayImpl()
}

func (impl *ObserverImpl) Play() {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	interconnect.RegisterObserverServer(s, impl)

	log.Printf("Starting Game server at port %v\n", port)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (impl *ObserverImpl) NewGameInstance(server interconnect.Observer_ProcessGameServer) *GameInstance {
	return &GameInstance{
		fuel:   impl.fuel,
		word:   util.NewSecretString(impl.word),
		server: server,
	}
}
