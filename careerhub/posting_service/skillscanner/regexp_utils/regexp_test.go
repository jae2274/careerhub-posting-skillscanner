package regexp_utils_test

import (
	"regexp"
	"testing"

	"github.com/jae2274/careerhub-posting-skillscanner/careerhub/posting_service/skillscanner/regexp_utils"
	"github.com/stretchr/testify/require"
)

func TestInitializeOnlyWordRegex(t *testing.T) {

	assertContainsMatchIn(t, regexp_utils.InitializeOnlyWordRegex("kotlin"), "javascript, typescript 개발 경험", false)
	assertContainsMatchIn(t, regexp_utils.InitializeOnlyWordRegex("java"), "javascript, typescript 개발 경험", false)
	assertContainsMatchIn(t, regexp_utils.InitializeOnlyWordRegex("springboot"), "spring boot api 개발 경험", false)

	assertContainsMatchIn(t, regexp_utils.InitializeOnlyWordRegex("javascript"), "node.js,JAVAscript,typescript 개발 경험", true)
	assertContainsMatchIn(t, regexp_utils.InitializeOnlyWordRegex("spring boot"), "springboot api 개발 경험", true)
	assertContainsMatchIn(t, regexp_utils.InitializeOnlyWordRegex("typescript"), "node.js,javascript,TYPESCRIPT 개발 경험", true)
	assertContainsMatchIn(t, regexp_utils.InitializeOnlyWordRegex("java"), "node.js,javascript,TYPESCRIPT 개발 경험", false)
	assertContainsMatchIn(t, regexp_utils.InitializeOnlyWordRegex("script"), "node.js,javascripts,TYPESCRIPT 개발 경험", false)

	// C++의 + 같은 특수문자는 유의하여 \\+로 대체하도록 한다.
	assertContainsMatchIn(t, regexp_utils.InitializeOnlyWordRegex("C++"), "C++개발 경험", true)
	assertContainsMatchIn(t, regexp_utils.InitializeOnlyWordRegex("C++"), "C/C++개발 경험", true)

	assertContainsMatchIn(t, regexp_utils.InitializeOnlyWordRegex("C"), "C++ 개발 경험", false)
	assertContainsMatchIn(t, regexp_utils.InitializeOnlyWordRegex("C"), "C# 개발 경험", false)

	assertContainsMatchIn(t, regexp_utils.InitializeOnlyWordRegex("C#"), "C#개발 경험", true)
}

func TestInitializePrefixRegex(t *testing.T) {
	assertContainsMatchIn(t, regexp_utils.InitializePrefixRegex("java"), "javascript", true)
	assertContainsMatchIn(t, regexp_utils.InitializePrefixRegex("java"), "java", true)
	assertContainsMatchIn(t, regexp_utils.InitializePrefixRegex("java"), "JAVASCRIPT", true)
	assertContainsMatchIn(t, regexp_utils.InitializePrefixRegex("java"), "JAVA", true)
	assertContainsMatchIn(t, regexp_utils.InitializePrefixRegex("JAVA"), "javascript", true)
	assertContainsMatchIn(t, regexp_utils.InitializePrefixRegex("JAVA"), "java", true)

	assertContainsMatchIn(t, regexp_utils.InitializePrefixRegex("spring boot"), "spring boot", true)
	assertContainsMatchIn(t, regexp_utils.InitializePrefixRegex("spring boot"), "springboot", true)
	assertContainsMatchIn(t, regexp_utils.InitializePrefixRegex("springboot"), "spring boot", false)
}

func assertContainsMatchIn(t *testing.T, regexpString string, text string, expected bool) {
	isContainsMatchIn := regexp.MustCompile(regexpString).MatchString(text)

	require.Equal(t, expected, isContainsMatchIn)
}
