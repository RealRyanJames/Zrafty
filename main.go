package main

import (
	commands "commands_app/Commands"
	"fmt"
	"time"
)

type MeshLength struct {
	SIZE map[string]string
}

type INTERMesh interface {
	getPrintln() string
}

type Setup struct {
	commands   string
	argsPassed string
}

func (r Setup) getPrintln() string {

	return r.argsPassed + " " + r.commands
}

type SLEEP struct {
	Time int
}

func (sleep SLEEP) Sleep() {
	if sleep.Time > 1 {

		time.Sleep(time.Duration(sleep.Time) * time.Second)
	} else {

		time.Sleep(time.Duration(sleep.Time) * time.Second)
	}
}

func main() {

	sleep := SLEEP{
		Time: 3.0,
	}

	commandsSender := MeshLength{
		SIZE: map[string]string{
			"Home":    "/",
			"Math":    "/math",
			"Counter": "/counter",
		},
	}

	for key, val := range commandsSender.SIZE {

		var file INTERMesh

		file = Setup{
			argsPassed: key,
			commands:   val,
		}

		sleep.Sleep()
		fmt.Println(file.getPrintln())
	}

	commands.Commands()

}
