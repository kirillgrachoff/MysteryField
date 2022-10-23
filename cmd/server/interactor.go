package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	"github.com/kirillgrachoff/MysteryField/pkg/interconnect"
	"github.com/kirillgrachoff/MysteryField/pkg/util"
)

type ObserverImpl struct {
	interconnect.UnimplementedObserverServer

	word string
	fuel uint32
}

func NewObserverImpl(fuel uint32, word string) *ObserverImpl {
	return &ObserverImpl{
		fuel: fuel,
		word: word,
	}
}

func (impl *ObserverImpl) ProcessGame(server interconnect.Observer_ProcessGameServer) error {
	log.Println("Client Accepted")
	log.Printf("Word: %v, fuel: %v", impl.word, impl.fuel)
	game := impl.NewGameInstance(server)
	game.server = server
	err := game.PlayImpl()
	log.Println("End Client game")
	return err
}

func (impl *ObserverImpl) ServeGrpc() {
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
