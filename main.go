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

type BaseMeshes struct {
	isCatching bool
}

func (baseMeshes BaseMeshes) GetBaseMesh() bool {
	if baseMeshes.isCatching {
		return baseMeshes.isCatching
	}

	return true
}

type MeshIncludesText struct {
	isLoaded   bool
	TextMeshes string
}

func (meshesLengthLoaded MeshIncludesText) GetLoadedMeshes() string {
	m := PrintMesh{}

	if meshesLengthLoaded.isLoaded && strings.Contains(meshesLengthLoaded.TextMeshes, "Break") {
		meshesLengthLoaded.TextMeshes = strings.ToUpper("Loaded Text")
		return m.GetMesh()
	}

	return string(meshesLengthLoaded.TextMeshes)
}

type EmptyTextMesh struct {
	isEmptyMeshCleared string
}

func (emptyStringValue EmptyTextMesh) CheckEmptyValue() bool {

	if emptyStringValue.isEmptyMeshCleared == "" {

		fmt.Println(strings.ToUpper("Input is: Empty Value"))
	}

	fmt.Println(strings.ToUpper("Input is Not Empty Value"))

	return emptyStringValue.isEmptyMeshCleared != ""
}

func main() {

	var s string

	baseMeshes := BaseMeshes{
		isCatching: true,
	}

	for baseMeshes.GetBaseMesh() {

		mesh := PrintMesh{
			messageUpper: strings.ToUpper("Enter Chosen Route: "),
		}

		fmt.Println(mesh.GetMesh())

		s = GetInput(s)

		isMeshTextContains := MeshIncludesText{
			isLoaded:   bool(baseMeshes.isCatching),
			TextMeshes: s,
		}

		sleep := SLEEP{
			Time: 3.0,
		}

		emptyTextMesh := EmptyTextMesh{
			isEmptyMeshCleared: s,
		}

		for emptyTextMesh.CheckEmptyValue() {

			fmt.Println(mesh.GetMesh())
			isMeshTextContains.GetLoadedMeshes()

			if strings.Contains(string(s), "/") {

				commandsSender := MeshLength{
					SIZE: map[string]string{
						"Home":    "/",
						"Math":    "/math",
						"Counter": "/counter",
					},
				}

				for i := range 1 {
					for key, val := range commandsSender.SIZE {
						i += 1

						var file INTERMesh

						file = Setup{
							argsPassed: key,
							commands:   val,
						}

						sleep.Sleep()

						fmt.Println(i, key, ":", val)

						file.getPrintln()
					}

				}

			} else {

				emptyTextMesh.CheckEmptyValue()
			}

			commands.Commands()
			mesh.GetMesh()
			fmt.Scanln()
		}
	}
}
