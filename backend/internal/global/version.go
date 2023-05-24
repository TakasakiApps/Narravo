package global

import "fmt"

const (
	VersionMajor = 0
	VersionSub   = 1
	VersionPatch = 0
	VersionExtra = "-rc1"
)

func GetVersion() string {
	return fmt.Sprintf("%v.%v.%v.%s", VersionMajor, VersionSub, VersionPatch, VersionExtra)
}
