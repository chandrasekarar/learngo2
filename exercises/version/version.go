package version

import "runtime"

// Version gets current Go version
func Version() string {
	return runtime.Version()
}
