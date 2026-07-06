package environment

/**
Variables that are set during build
*/

// IsDocker has to be true if compiled for the Docker image (auto-generated value)
var IsDocker = "false"

// BuildTime is the time of the build (auto-generated value)
var BuildTime = "Dev Build"

// Builder is the name of builder (auto-generated value)
var Builder = "Manual Build"

// Commit is the git commit the binary was built from (auto-generated value)
var Commit = ""

// SourceRepo is the base URL of the corresponding (AGPL) source repository for this build
var SourceRepo = "https://github.com/manoangazi/gokapi"

// SourceCodeURL returns a link to the source corresponding to this exact build, as
// required by AGPL-3.0 section 13. It pins to the built-from commit when known and
// otherwise falls back to the hardening branch (e.g. a manual build).
func SourceCodeURL() string {
	if Commit != "" {
		return SourceRepo + "/tree/" + Commit
	}
	return SourceRepo + "/tree/hardening"
}

// IsDockerInstance returns true if the binary was compiled with the official docker makefile, which
// sets IsDocker to true
func IsDockerInstance() bool {
	return IsDocker != "false"
}
