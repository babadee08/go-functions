package semantics

import "fmt"

type SemanticVersion struct {
	major, minor, patch int
}

func NewSemanticVersion(major, minor, patch int) SemanticVersion {
	return SemanticVersion{
		major: major,
		minor: minor,
		patch: patch,
	}
}

func (sv SemanticVersion) String() string {
	return fmt.Sprintf("%v.%v.%v", sv.major, sv.minor, sv.patch)
}

func (sv *SemanticVersion) IncrementMajor() {
	sv.major++
}

func (sv *SemanticVersion) IncrementMinor() {
	sv.minor++
}

func (sv *SemanticVersion) IncrementPatch() {
	sv.patch++
}
