package vars

import (
	"fmt"
	"os"
)

type DBUser struct {
	Username string
	Password string
}

type Vars struct {
	GrpcEndpoint string
}

type ErrNotExistedVar struct {
	VarName string
}

func NotExistedVar(varName string) *ErrNotExistedVar {
	return &ErrNotExistedVar{VarName: varName}
}

func (e *ErrNotExistedVar) Error() string {
	return fmt.Sprintf("%s is not existed", e.VarName)
}

func Variables() (*Vars, error) {
	grpcEndpoint, err := getFromEnv("GRPC_ENDPOINT")
	if err != nil {
		return nil, err
	}

	return &Vars{
		GrpcEndpoint: grpcEndpoint,
	}, nil
}

func getFromEnv(envVar string) (string, error) {
	ev := os.Getenv(envVar)

	if ev == "" {
		return "", fmt.Errorf("%s is not existed", envVar)
	}

	return ev, nil
}
