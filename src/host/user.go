package host

import (
	"os/exec"
	"os/user"
	"strings"

	cenvvm "github.com/custodia-cenv/cenv-vm/src"
)

// UserHasSystemPrivileges gib an ob die VM als Root/Administrator oder mit normalen Benutzerberechtigungen ausgeführt wird
func UserHasSystemPrivileges() bool {
	currentUser, err := user.Current()
	if err != nil {
		return false
	}
	return currentUser.Uid == "0"
}

// IsUserInGroup gib an ob der Aktuelle Benutzer Mitglied einer Spiziellen Gruppe ist
func IsUserInGroup(groupName cenvvm.HostUserGroupName) bool {
	// Hole den aktuellen Benutzer
	currentUser, err := user.Current()
	if err != nil {
		return false
	}

	// Hole die Gruppen des Benutzers mithilfe des Betriebssystems
	cmd := exec.Command("id", "-Gn", currentUser.Username)
	output, err := cmd.Output()
	if err != nil {
		return false
	}

	// Parse die Gruppen
	groups := strings.Fields(string(output))
	for _, group := range groups {
		if group == string(groupName) {
			return true
		}
	}

	return false
}

// GetUsernameNameAndPrimaryGroupName gibt den Benutzernamen des Aktuellen Benutezrs sowie seine Primäre Gruppe zurück
func GetUsernameNameAndPrimaryGroupName() (string, string) {
	// Aktuellen Benutzer abrufen
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	// Primäre Gruppe des Benutzers abrufen
	groupID := currentUser.Gid
	group, err := user.LookupGroupId(groupID)
	if err != nil {
		panic(err)
	}
	return currentUser.Username, group.Name
}
