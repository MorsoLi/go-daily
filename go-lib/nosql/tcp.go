package main

import {
	"fmt"
}

type Config struct {

}
type Handler interface {
	Handle(ctx context.Context, conn net.Conn)
	Close() error
}
type Config struct {
	Address    string        `yaml:"address"`
	MaxConnect uint32        `yaml:"max-connect"`
	Timeout    time.Duration `yaml:"timeout"`
}
func ListenAndServerWithSignal(cfg *Config, handler Handler){
	fmt.Println("hhhhhhh")
}