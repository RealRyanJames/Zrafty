package math_route

import "fmt"

type MeshType struct {
	isMathProd bool
}

func CustomTypeMesh[T any](r T) T {

	fmt.Println("Version is:", r)
	return r
}
