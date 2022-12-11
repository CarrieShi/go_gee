package main

type Factory func() Server

var factory Factory

func RegisterFactory(f Factory) {
	factory = f
}

func NewServer() Server {
	return factory()
}
