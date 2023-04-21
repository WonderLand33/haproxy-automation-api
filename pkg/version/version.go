package version

import (
	"encoding/json"
	"fmt"
	"runtime"
)

var (
	gitBranch = "unknown"
	gitCommit = "unknown"
	gitTag    = "unknown"
	buildUser = "unknown"
	buildDate = "unknown"
)

type Info struct {
	GitBranch string `json:"gitBranch"`
	GitCommit string `json:"gitCommit"`
	GitTag    string `json:"gitTag"`
	BuildUser string `json:"buildUser"`
	BuildDate string `json:"buildDate"`
	GoVersion string `json:"goVersion"`
	Compiler  string `json:"compiler"`
	Platform  string `json:"platform"`
}

func (info Info) String() string {
	jsonString, _ := json.Marshal(info)
	return string(jsonString)
}

// Get returns the overall codebase version. It's for
// detecting what code a binary was built from.
func Get() Info {
	// These variables typically come from -ldflags settings and
	// in their absence fallback to the default settings
	return Info{
		GitBranch: gitBranch,
		GitCommit: gitCommit,
		GitTag:    gitTag,
		BuildUser: buildUser,
		BuildDate: buildDate,
		GoVersion: runtime.Version(),
		Compiler:  runtime.Compiler,
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
