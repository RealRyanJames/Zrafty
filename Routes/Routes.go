package routes

import "math/rand"

type CustomRouteMesh string
type NumMesh int

type RAND_NUMBER struct {
	NUM1 NumMesh
	NUM2 NumMesh
}

func GetIniialRoute() RAND_NUMBER {

	numMesh := RAND_NUMBER{
		NUM1: NumMesh(rand.Intn(6)),
		NUM2: NumMesh(rand.Intn(9)),
	}

	return numMesh
}
