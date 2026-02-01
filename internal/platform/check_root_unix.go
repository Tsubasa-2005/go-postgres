//go:build !windows

package platform

import (
	"errors"
	"fmt"
	"os"
)

func CheckRoot(progname string) error {
	if os.Geteuid() == 0 {
		return errors.New(`"root" execution of the PostgreSQL server is not permitted.
The server must be started under an unprivileged user ID to prevent
possible system security compromise.  See the documentation for
more information on how to properly start the server.`)
	}

	if os.Getuid() != os.Geteuid() {
		return fmt.Errorf("%s: real and effective user IDs must match", progname)
	}

	return nil
}
