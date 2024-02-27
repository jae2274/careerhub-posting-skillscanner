package regexp_utils

import "strings"

var needReplaceSpecifics = []struct {
	original    string
	replacement string
}{
	{"+", "\\+"},
	{"(", "\\("},
	{")", "\\)"},
}

func preprocess(str string) string {
	processed := str
	for _, pair := range needReplaceSpecifics {
		processed = strings.ReplaceAll(processed, pair.original, pair.replacement)
	}
	processed = strings.ReplaceAll(processed, " ", "\\s*")
	return strings.ToLower(processed)
}

// func hasKoreanChar(str string) bool {
// 	return regexp.MustCompile(".*[ㄱ-ㅎ|ㅏ-ㅣ|가-힣]+.*").MatchString(str)
// }

// func initializeExactlyRegex(str string) string {
// 	return "^" + preprocess(str) + "$"
// }

func InitializeOnlyWordRegex(str string) string {
	preprocessed := preprocess(str)
	return "(?i)(^|[^a-zA-Z])" + preprocessed + "($|[^a-zA-Z\\+#])"
}

func InitializePrefixRegex(str string) string {
	preprocessed := preprocess(str)
	return "(?i)^" + preprocessed + ".*"
}
