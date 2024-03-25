package app

import "github.com/jae2274/goutils/enum"

type SkillFromValues struct{}
type SkillFrom = enum.Enum[SkillFromValues]

const (
	Origin             = SkillFrom("ORIGIN")
	FromTitle          = SkillFrom("FROM_TITLE")
	FromMainTask       = SkillFrom("FROM_MAIN_TASK")
	FromQualifications = SkillFrom("FROM_QUALIFICATIONS")
	FromPreferred      = SkillFrom("FROM_PREFERRED")
)

func (SkillFromValues) Values() []string {
	return []string{string(Origin), string(FromTitle), string(FromMainTask), string(FromQualifications), string(FromPreferred)}
}
