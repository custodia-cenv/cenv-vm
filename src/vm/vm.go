package vm

import (
	"github.com/custodia-cenv/cenvx-core/src/core"
)

// Initalisiert den VM Core Prozess
func InitVmProcessInstance() error {
	// Die Verbindung zum Core wird Initalisiert
	err := initCoreVmIpcClientSession([]*core.ACL{})
	if err != nil {
		return err
	}

	// Es ist kein Fehler beim Initalisieren der Verbindung aufgetreten
	return nil
}

// Wird verwendet nachdem das VM Image geladen wurde um die Manifest Datei zu Initalisieren
func InitVmImgManifestWithCore() error {
	return nil
}

// Registriert die VM in anderen Services
func InitAntoherVmSystemServices() error {
	return nil
}

// Registriert das Working dir
func SetupWorkingDir() error {
	return nil
}
