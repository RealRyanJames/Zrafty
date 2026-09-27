package main

import (
	commands "commands_app/Commands"
	extra "commands_app/Extrra"
	functtions "commands_app/Functtions"
	routes "commands_app/Routes"
	"errors"
	"fmt"
	"math/rand"
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
	} else {
		return false
	}

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

	return emptyStringValue.isEmptyMeshCleared == ""
}

type UILabelStyle string

type lineStringMesh struct {
	isLoadedOnUIChanged bool
	linesUI             UILabelStyle
}

type PositionsChanged struct {
	posX float64
	posY float64
	posZ float64
}

func GetPositions() float64 {

	positions := PositionsChanged{
		posX: 40.0,
		posY: 2.0,
		posZ: 0,
	}

	return float64(positions.posX*positions.posY/2) / 2 * 2

}

type DateLoggerMesh struct {
	isHourlyDay bool
	isMorning   bool
	isAfternoon bool
	isEvening   bool
}

type TimeOfDay struct {
	morning   bool
	afternoon bool
	evening   bool
}

type DateLoggingSystem struct {
	date any
}

func (date DateLoggingSystem) GetDate() any {
	return date.date
}

func (dateLoggerMesh DateLoggerMesh) GetDateMesh() string {
	timeOfDay := TimeOfDay{
		morning:   time.Now().Hour() < 11,
		afternoon: time.Now().Hour() > 11 && time.Now().Hour() < 16,
		evening:   time.Now().Hour() > 16,
	}

	isMorning := timeOfDay.morning
	isAfternoon := timeOfDay.afternoon
	isEvening := timeOfDay.evening

	if isMorning {
		return strings.ToUpper("Good Morning User")
	}

	if isAfternoon {

		return strings.ToUpper("Good Afternoon User")
	}

	if isEvening {
		return strings.ToUpper("Good Evening User")
	}

	return ""
}

func (styledLabelUI lineStringMesh) GetUIChanged() string {
	if styledLabelUI.isLoadedOnUIChanged {

		for i := 0; i < int(GetPositions()); i++ {
			fmt.Print("-")
		}
	}

	return string("-\n")
}

type NowTimeMesh = string

type ErrorLines struct {
	lineOnError []NowTimeMesh
}

type PickedOperationMesh struct {
	PickedDrawOPSize []string
	isKnownSize      bool
}

func (opPicked PickedOperationMesh) GetPickedOperator() string {

	if opPicked.isKnownSize {

		str_op := []string{"+", "-", "*"}
		r_op := len(str_op)

		selected_op_len := rand.Intn(r_op)

		return str_op[selected_op_len]
	}

	return string("")
}

func main() {

	var s string

	baseMeshes := BaseMeshes{
		isCatching: true,
	}

	MeshUIChanged := lineStringMesh{
		isLoadedOnUIChanged: true,
	}

	MeshUIChanged.GetUIChanged()

	fmt.Println("")

	for baseMeshes.GetBaseMesh() {

		dateNow := DateLoggerMesh{
			isMorning:   time.Now().Hour() < 11,
			isAfternoon: time.Now().Hour() > 11 && time.Now().Hour() < 17,
			isEvening:   time.Now().Hour() > 18,
		}

		r := extra.CustomType{
			MeshVersion: []string{"Current Version: ", "v1.", "0.", "0\n"},
			MeshLogger:  []string{},
		}

		for _, val := range r.MeshVersion {
			fmt.Print(val)
		}

		fmt.Println(strings.ToUpper(time.Now().Month().String()), "/", time.Now().Day(), "/", time.Now().Year())

		fmt.Println(dateNow.GetDateMesh())

		mesh := PrintMesh{
			messageUpper: strings.ToUpper("Enter Chosen Route: "),
		}

		fmt.Println(mesh.GetMesh())
		fmt.Scanln(&s)

		if strings.Contains(s, "/Math") {

			num1 := routes.GetIniialRoute()

			OP_PICKED := PickedOperationMesh{
				isKnownSize: true,
			}
			op := OP_PICKED.GetPickedOperator()

			answer := 0
			fmt.Printf("What is: %d %s %d\n", num1.NUM1, op, num1.NUM2)
			fmt.Scanln(&answer)

			if op == "+" && answer == functtions.Add(int(num1.NUM1), int(num1.NUM2)) {

				fmt.Println(strings.ToUpper("Correct Answer Good Job!"))
			}

			if op == "+" && answer != functtions.Add(int(num1.NUM1), int(num1.NUM2)) {

				fmt.Println(strings.ToUpper("Incorrect Answer Please Try Again Next Time?"))
			}

			if op == "-" && answer == functtions.Sub(int(num1.NUM1), int(num1.NUM2)) {

				fmt.Println(strings.ToUpper("Correct Answer Good Job!"))
			}

			if op == "-" && answer != functtions.Sub(int(num1.NUM1), int(num1.NUM2)) {

				fmt.Println(strings.ToUpper("Incorrect Answer Please Try Again Next Time?"))
			}

			if op == "*" && answer == functtions.Multiply(int(num1.NUM1), int(num1.NUM2)) {

				fmt.Println(strings.ToUpper("Correct Answer Good Job!"))
			}

			if op == "*" && answer != functtions.Multiply(int(num1.NUM1), int(num1.NUM2)) {

				fmt.Println(strings.ToUpper("Incorrect Answer Please Try Again Next Time?"))
			}

		} else {

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

			if emptyTextMesh.CheckEmptyValue() {

				break
			} else {
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

				}
			}

		}

		commands.Commands()
		baseMeshes.isCatching = false
		MeshUIChanged.GetUIChanged()
		fmt.Scanln()
	}

}
