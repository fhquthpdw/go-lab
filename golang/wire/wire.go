//go:build wireinject

package main

import "github.com/google/wire"

var providerSet wire.ProviderSet = wire.NewSet(NewMessage, NewGreeter)

func InitializeEvent(phrase string, code int) (Event, error) {
	wire.Build(NewEvent, NewGreeter, wire.Struct(new(Message), "Content"))
	return Event{}, nil
}
