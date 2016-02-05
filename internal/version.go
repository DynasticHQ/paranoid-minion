package version

import (
	"fmt"
	"runtime"
)

// http://semver.org/
const Binary = "0.0.1-alpha"

func String(app string) string {
	return fmt.Sprintf("%s v%s (built w/%s)", app, Binary, runtime.Version())
}
