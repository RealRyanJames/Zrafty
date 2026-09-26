package commands

import (
	"fmt"
	"strings"
)

type MeshsCommands struct {
	is_running bool
}

func (mesh MeshsCommands) GetMesh() {
	if mesh.is_running {
		fmt.Println(strings.ToUpper("Route Visited"))
	}

}

func Commands() {
	command := MeshsCommands{
		is_running: true,
	}

	command.GetMesh()
	fmt.Scanln()
}
