package main

import (
	commands "commands_app/Commands"
	"errors"
	"fmt"
	"os"
	"strings"
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

func GetInput(str string) string {

	fmt.Scan(&str)
	return str
}

type PrintMesh struct {
	messageUpper string
}

type ErrorMesh struct {
	err error
}

func (errHandledMesh ErrorMesh) GetErrorMesh() {
	if errHandledMesh.err == errors.New("Exit") {

		os.Exit(0)

	} else {
		os.Exit(1)
	}
}

func (println PrintMesh) GetMesh() string {

	if len(println.messageUpper) < 4 {

		err := ErrorMesh{
			err: errors.New("Exit"),
		}

		err.GetErrorMesh()
	}

	return println.messageUpper
}

func main() {

	var s string
	s = GetInput(s)

	mesh := PrintMesh{
		messageUpper: strings.ToUpper("Enter Chosen Route: "),
	}

	fmt.Println(mesh.GetMesh())

	sleep := SLEEP{
		Time: 3.0,
	}

	if strings.Contains(string(s), "/") {

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

	}

	commands.Commands()
	fmt.Scanln()
	mesh.GetMesh()
}
