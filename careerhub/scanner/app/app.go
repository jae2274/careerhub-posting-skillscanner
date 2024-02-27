package app

import (
	"context"
	"io"
	"regexp"

	"github.com/jae2274/Careerhub-SkillScanner/careerhub/scanner/regexp_utils"
	"github.com/jae2274/Careerhub-SkillScanner/careerhub/scanner/scanner_grpc"
	"github.com/jae2274/goutils/enum"
	"github.com/jae2274/goutils/terr"
)

type App struct {
	grpcClient scanner_grpc.ScannerGrpcClient
}

func NewApp(grpcClient scanner_grpc.ScannerGrpcClient) *App {
	return &App{grpcClient: grpcClient}
}

type ScanTargetValues struct{}

type ScanTarget = enum.Enum[ScanTargetValues]

const (
	Skill      = ScanTarget("skill")
	JobPosting = ScanTarget("jobPosting")
)

func (ScanTargetValues) Values() []string {
	return []string{
		string(Skill),
		string(JobPosting),
	}
}

func StartScanForNewSkills(grpcClient scanner_grpc.ScannerGrpcClient, scanTarget ScanTarget) error {

	isTargetSkill := scanTarget == Skill
	isTargetJobPosting := scanTarget == JobPosting

	mainCtx := context.Background()
	skills, err := grpcClient.GetSkills(mainCtx, &scanner_grpc.ScanComplete{IsScanComplete: !isTargetSkill}) //스캔 목적이 스킬이 아닌 경우 이미 완료된 스킬 목록을 가져옴
	if err != nil {
		return err
	}

	jobPostingStream, err := grpcClient.GetJobPostings(mainCtx, &scanner_grpc.ScanComplete{IsScanComplete: !isTargetJobPosting}) //스캔 목적이 채용공고가 아닌 경우 이미 완료된 채용공고 목록을 가져옴
	if err != nil {
		return err
	}

	sendRequestStream, err := grpcClient.SetRequiredSkills(mainCtx)
	if err != nil {
		return err
	}

	for {
		jobPosting, err := jobPostingStream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		alreadyExistedSkills := make(map[string]bool)
		for _, requiredSkill := range jobPosting.RequiredSkill {
			alreadyExistedSkills[requiredSkill] = true
		}

		additionalSkills := make([]string, 0)
		for _, skillName := range skills.SkillNames {
			if _, ok := alreadyExistedSkills[skillName]; !ok { //존재하지 않는다면 스캔
				if CheckSkillRequirement(jobPosting, skillName) {
					additionalSkills = append(additionalSkills, skillName)
				}
			}
		}

		if len(additionalSkills) > 0 {
			err = sendRequestStream.Send(&scanner_grpc.SetRequiredSkillsRequest{
				Site:          jobPosting.Site,
				PostingId:     jobPosting.PostingId,
				RequiredSkill: additionalSkills,
			})
			if err != nil {
				return err
			}
		}
	}

	response, err := sendRequestStream.CloseAndRecv()
	if err != nil {
		return err
	}
	if !response.Success {
		return terr.New("Failed to set required skills")
	}

	err = jobPostingStream.CloseSend()
	if err != nil {
		return err
	}

	if isTargetSkill { //스킬 스캔이 목적인 경우 별도로 스킬의 스캔 완료를 알림
		response, err = grpcClient.SetScanComplete(mainCtx, skills)
		if err != nil {
			return err
		}
		if !response.Success {
			return terr.New("Failed to set required skills")
		}
	}

	return nil
}

func CheckSkillRequirement(jobPosting *scanner_grpc.JobPostingInfo, skillName string) bool {
	regexString := regexp_utils.InitializeOnlyWordRegex(skillName)
	if regexp.MustCompile(regexString).MatchString(jobPosting.Title) {
		return true
	}
	if regexp.MustCompile(regexString).MatchString(jobPosting.Qualifications) {
		return true
	}
	if regexp.MustCompile(regexString).MatchString(jobPosting.Preferred) {
		return true
	}

	return false
}
