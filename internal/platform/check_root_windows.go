//go:build windows

package platform

import (
	"errors"
	"fmt"

	"golang.org/x/sys/windows"
)

func CheckRoot(_ string) error {
	var token windows.Token
	err := windows.OpenProcessToken(windows.CurrentProcess(), windows.TOKEN_QUERY, &token)
	if err != nil {
		return fmt.Errorf("failed to check admin privileges: %w", err)
	}
	defer token.Close()

	adminSid, err := windows.CreateWellKnownSid(windows.WinBuiltinAdministratorsSid)
	if err != nil {
		return fmt.Errorf("failed to create admin SID: %w", err)
	}

	isAdmin, err := token.IsMember(adminSid)
	if err != nil {
		return fmt.Errorf("failed to check admin membership: %w", err)
	}

	if isAdmin {
		return errors.New(`Execution of PostgreSQL by a user with administrative permissions is not
 permitted.
 The server must be started under an unprivileged user ID to prevent
 possible system security compromises.  See the documentation for
 more information on how to properly start the server`)
	}

	return nil
}
