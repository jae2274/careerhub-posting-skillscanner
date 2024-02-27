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
	PostLogUrl   string
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

	postLogUrl, err := getFromEnv("POST_LOG_URL")
	if err != nil {
		return nil, err
	}

	return &Vars{
		GrpcEndpoint: grpcEndpoint,
		PostLogUrl:   postLogUrl,
	}, nil
}

func getFromEnv(envVar string) (string, error) {
	ev := os.Getenv(envVar)

	if ev == "" {
		return "", fmt.Errorf("%s is not existed", envVar)
	}

	return ev, nil
}
