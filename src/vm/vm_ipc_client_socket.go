package vm

import (
	"fmt"
	"net"
	"os"
	"path/filepath"

	"github.com/CustodiaJS/bngsocket"
	cenvvm "github.com/custodia-cenv/cenv-vm/src"
	"github.com/custodia-cenv/cenv-vm/src/host"
	"github.com/custodia-cenv/cenv-vm/src/host/filesystem"
	cenvxcore "github.com/custodia-cenv/cenvx-core/src"
	"github.com/custodia-cenv/cenvx-core/src/core"
)

// initCoreVmIpcClientSession Initalisiert den VM-IPC Core Socket
func initCoreVmIpcClientSession(userGroups []*core.ACL) error {
	// Es wird geprüft ob der Benutzer Systemrechte hat
	if host.UserHasSystemPrivileges() {
		// Es wird versucht für den Root User eine IPC Verbindung zu registrieren
		rouvmErr := initRootUserVmIpcClientSession(string(cenvxcore.CoreVmIpcRootSocketPath))
		if rouvmErr == nil {
			return nil
		}
	}

	// Es versucht anhand der Angegeben Benutzer Gruppen eine Verbindung mit dem Core herzustellen
	if len(userGroups) > 0 {
		for _, item := range userGroups {
			usergr := initSpeficUserGroupVmIpcClientSession(*item.Groupname)
			if usergr == nil {
				return nil
			}
		}
	}

	// Es wird versucht eine Current User Sitzung mit dem Core Service aufzubauen
	currUserSessionErr := initCurrentUserVmIpcClientSession()
	if currUserSessionErr == nil {
		return nil
	}

	// Es konnte keine Core Service Instanz gefunden werden
	return cenvvm.ErrNoCoreServiceRunning
}

// initRootUserVmIpcClientSession erstellt eine RootUser basierte IPC Verbindung mit dem Core her
func initRootUserVmIpcClientSession(basePath string) error {
	// Check if basePath is empty
	if basePath == "" {
		return fmt.Errorf("%w: the provided basePath is empty", cenvvm.ErrEmptyBasePath)
	}

	// Der Altuelle Path wird erstellt
	socketPath := filepath.Join(basePath, fmt.Sprintf("%s_root.sock", cenvvm.CORE_SOCKET_PREFIX))

	// Es wird geprüft ob die Datei vorhanden ist
	if !filesystem.FileExists(socketPath) {
		return fmt.Errorf("%w: socket file not found at %s", cenvvm.ErrSocketNotFound, socketPath)
	}

	// Additional check: Permissions
	info, err := os.Stat(socketPath)
	if err != nil {
		if os.IsPermission(err) {
			return fmt.Errorf("%w: unable to access socket at %s", cenvvm.ErrPermission, socketPath)
		}
		return fmt.Errorf("%w: unexpected error accessing socket at %s", cenvvm.ErrInvalidSocket, socketPath)
	}

	// Check if the path is a file, not a directory
	if info.IsDir() {
		return fmt.Errorf("%w: socket path is a directory, not a file: %s", cenvvm.ErrInvalidSocket, socketPath)
	}

	// Verbindung zum Unix-Socket herstellen
	socketConn, socketOpenningError := net.Dial("unix", socketPath)
	if socketOpenningError != nil {
		return socketOpenningError
	}

	// Die Verbindung wird geupgradet und der Prozess wird Initalisiert
	if upgradeInitError := upgradeSocketToBngSocketAndInitNewProcess(socketConn); upgradeInitError != nil {
		return upgradeInitError
	}

	// Der Vorgang ist ohne Fehler durchgeführt wurden
	return nil
}

// initSpeficUserGroupVmIpcClientSession wird verwendet um eine IPC Verbindung auf basis einer Gruppenmitgliedschaft herzustellen
func initSpeficUserGroupVmIpcClientSession(groupName string) error {
	// Der Altuelle Path wird erstellt
	socketPath := string(cenvxcore.GetCoreSpeficSocketUserGroupPath(groupName))

	// Es wird geprüft ob die Datei vorhanden ist
	if !filesystem.FileExists(socketPath) {
		return fmt.Errorf("%w: socket file not found at %s", cenvvm.ErrSocketNotFound, socketPath)
	}

	// Additional check: Permissions
	info, err := os.Stat(socketPath)
	if err != nil {
		if os.IsPermission(err) {
			return fmt.Errorf("%w: unable to access socket at %s", cenvvm.ErrPermission, socketPath)
		}
		return fmt.Errorf("%w: unexpected error accessing socket at %s", cenvvm.ErrInvalidSocket, socketPath)
	}

	// Check if the path is a file, not a directory
	if info.IsDir() {
		return fmt.Errorf("%w: socket path is a directory, not a file: %s", cenvvm.ErrInvalidSocket, socketPath)
	}

	// Verbindung zum Unix-Socket herstellen
	socketConn, socketOpenningError := net.Dial("unix", socketPath)
	if socketOpenningError != nil {
		return socketOpenningError
	}

	// Die Verbindung wird geupgradet und der Prozess wird Initalisiert
	if upgradeInitError := upgradeSocketToBngSocketAndInitNewProcess(socketConn); upgradeInitError != nil {
		return upgradeInitError
	}

	// Der Vorgang ist ohne Fehler durchgeführt wurden
	return nil
}

// initCurrentUserVmIpcClientSession wird verwendet um eine IPC Verbindung auf basis des Aktuellen Benutzers herzustellen
func initCurrentUserVmIpcClientSession() error {
	// Der Aktuelle Benutezrname sowie Primäre Gruppenname wird abgerufen
	username, primarygroupname := host.GetUsernameNameAndPrimaryGroupName()

	// Der Path für den Aktuellen Benutzer wird erstellt
	socketPath := string(cenvxcore.GetCoreSpeficSocketUserAndGroupPath(username, primarygroupname))

	// Es wird geprüft ob die Datei vorhanden ist
	if !filesystem.FileExists(socketPath) {
		return fmt.Errorf("%w: socket file not found at %s", cenvvm.ErrSocketNotFound, socketPath)
	}

	// Additional check: Permissions
	info, err := os.Stat(socketPath)
	if err != nil {
		if os.IsPermission(err) {
			return fmt.Errorf("%w: unable to access socket at %s", cenvvm.ErrPermission, socketPath)
		}
		return fmt.Errorf("%w: unexpected error accessing socket at %s", cenvvm.ErrInvalidSocket, socketPath)
	}

	// Check if the path is a file, not a directory
	if info.IsDir() {
		return fmt.Errorf("%w: socket path is a directory, not a file: %s", cenvvm.ErrInvalidSocket, socketPath)
	}

	// Verbindung zum Unix-Socket herstellen
	socketConn, socketOpenningError := net.Dial("unix", socketPath)
	if socketOpenningError != nil {
		return socketOpenningError
	}

	// Die Verbindung wird geupgradet und der Prozess wird Initalisiert
	if upgradeInitError := upgradeSocketToBngSocketAndInitNewProcess(socketConn); upgradeInitError != nil {
		return upgradeInitError
	}

	// Der Vorgang ist ohne Fehler durchgeführt wurden
	return nil
}

// upgradeSocketToBngSocket wird verwenet um eine Socket Verbindung zu upgraden
func upgradeSocketToBngSocketAndInitNewProcess(socket net.Conn) error {
	// Die Verbindung wird geupgradet
	var upgradeError error
	clientIpcVmSokcet, upgradeError = bngsocket.UpgradeSocketToBngConn(socket)
	if upgradeError != nil {
		return upgradeError
	}

	return nil
}
