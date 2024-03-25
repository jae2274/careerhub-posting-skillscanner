package app_test

import (
	"testing"

	"github.com/jae2274/careerhub-posting-skillscanner/careerhub/posting_service/skillscanner/app"
	"github.com/jae2274/careerhub-posting-skillscanner/careerhub/posting_service/skillscanner/scanner_grpc"
	"github.com/stretchr/testify/require"
)

func TestCheckSkillRequirement(t *testing.T) {
	jobPosting := &scanner_grpc.JobPostingInfo{
		Site:           "jumpit",
		PostingId:      "123",
		Title:          "Java 개발자",
		MainTask:       "Java를 이용한 서버 개발, aws 환경에서 서버 구축",
		Qualifications: "Java 3년 이상, Spring Boot 2년 이상",
		Preferred:      "Kotlin 1년 이상",
		RequiredSkill:  []string{},
	}

	require.Equal(t, *app.CheckSkillRequirement(jobPosting, "kotlin"), app.FromPreferred)
	require.Equal(t, *app.CheckSkillRequirement(jobPosting, "java"), app.FromTitle)
	require.Equal(t, *app.CheckSkillRequirement(jobPosting, "spring boot"), app.FromQualifications)
	require.Equal(t, *app.CheckSkillRequirement(jobPosting, "aws"), app.FromMainTask)
	require.Nil(t, app.CheckSkillRequirement(jobPosting, "javascript"))
	require.Nil(t, app.CheckSkillRequirement(jobPosting, "typescript"))
	require.Nil(t, app.CheckSkillRequirement(jobPosting, "springboot"))
}
