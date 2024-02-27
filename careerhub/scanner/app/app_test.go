package app_test

import (
	"testing"

	"github.com/jae2274/Careerhub-SkillScanner/careerhub/scanner/app"
	"github.com/jae2274/Careerhub-SkillScanner/careerhub/scanner/scanner_grpc"
	"github.com/stretchr/testify/require"
)

func TestCheckSkillRequirement(t *testing.T) {
	jobPosting := &scanner_grpc.JobPostingInfo{
		Site:           "jumpit",
		PostingId:      "123",
		Title:          "Java 개발자",
		Qualifications: "Java 3년 이상, Spring Boot 2년 이상",
		Preferred:      "Kotlin 1년 이상",
		RequiredSkill:  []string{},
	}

	require.True(t, app.CheckSkillRequirement(jobPosting, "kotlin"))
	require.True(t, app.CheckSkillRequirement(jobPosting, "java"))
	require.True(t, app.CheckSkillRequirement(jobPosting, "spring boot"))
	require.False(t, app.CheckSkillRequirement(jobPosting, "javascript"))
	require.False(t, app.CheckSkillRequirement(jobPosting, "typescript"))
	require.False(t, app.CheckSkillRequirement(jobPosting, "springboot"))
}
