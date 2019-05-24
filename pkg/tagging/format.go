package tagging

import (
	"regexp"
	"strings"
)

type FormatType int
const (
	SEMVER FormatType = iota
	RC
	NUM
	DATEH
	DATED
)

type Format struct {
	Value string
	Types []FormatType
}

func (f *Format) Parse() {
	f.extractTypes()
}

func (f *Format) extractTypes() {
	if strings.Contains(f.Value, "{SEMVER}") {
		f.Types = append(f.Types, SEMVER)
	}

	if strings.Contains(f.Value, "{RC}") {
		f.Types = append(f.Types, RC)
	}

	if strings.Contains(f.Value, "#") {
		f.Types = append(f.Types, NUM)
	}

	if matched, _ := regexp.Match(`\d{4}-\d{2}-\d{2}`, []byte(f.Value)); matched {
		f.Types = append(f.Types, DATEH)
	}

	if matched, _ := regexp.Match(`\d{4}\.\d{2}\.\d{2}`, []byte(f.Value)); matched {
		f.Types = append(f.Types, DATED)
	}
}

