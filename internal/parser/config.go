package parser

import "modernc.org/cc/v3"

type Config struct {
	IncludePaths []string
	Defines      []string

	CppCommand string

	SkipHostConfig bool
	OverrideConfig *cc.Config
}
