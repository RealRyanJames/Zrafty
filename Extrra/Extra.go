package extra

type MeshLogging = []string

type CustomType struct {
	MeshVersion MeshLogging
	MeshLogger  MeshLogging
}

func GetMeshCompatable(MeshInit CustomType) *CustomType {
	return &CustomType{
		MeshVersion: []string{"1.0.0"},
		MeshLogger:  []string{},
	}

}
