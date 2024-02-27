package main

import (
	"context"
	"os"

	"github.com/jae2274/Careerhub-SkillScanner/careerhub/scanner/app"
	"github.com/jae2274/Careerhub-SkillScanner/careerhub/scanner/logger"
	"github.com/jae2274/Careerhub-SkillScanner/careerhub/scanner/scanner_grpc"
	"github.com/jae2274/Careerhub-SkillScanner/careerhub/scanner/vars"
	"github.com/jae2274/goutils/llog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	mainCtx := context.Background()
	envVars, err := vars.Variables()
	checkErr(mainCtx, err)

	initLogger(mainCtx, envVars.PostLogUrl)

	conn, err := grpc.Dial(envVars.GrpcEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	checkErr(mainCtx, err)

	grpcClient := scanner_grpc.NewScannerGrpcClient(conn)

	llog.Info(mainCtx, "Start scanning for new skills")
	err = app.StartScanForNewSkills(grpcClient, app.Skill) //신규 스킬을 대상으로 기존 채용공고를 스캔
	checkErr(mainCtx, err)

	llog.Info(mainCtx, "Start scanning for new job postings")
	err = app.StartScanForNewSkills(grpcClient, app.JobPosting) //신규 채용공고를 대상으로 기존 스킬을 스캔
	checkErr(mainCtx, err)

}

const (
	appName = "skill_scanner"
	service = "careerhub"

	ctxKeyTraceID = "trace_id"
)

func initLogger(ctx context.Context, postUrl string) error {
	llog.SetMetadata("service", service)
	llog.SetMetadata("app", appName)
	llog.SetDefaultContextData(ctxKeyTraceID)

	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	llog.SetMetadata("hostname", hostname)

	appLogger, err := logger.NewAppLogger(ctx, postUrl)
	if err != nil {
		return err
	}

	llog.SetDefaultLLoger(appLogger)

	return nil
}

func checkErr(ctx context.Context, err error) {
	if err != nil {
		llog.LogErr(ctx, err)
		os.Exit(1)
	}
}
