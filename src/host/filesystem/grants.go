package filesystem

import "os"

func HasUserAccess(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsPermission(err) {
			return false, nil // Keine Berechtigung
		}
		return false, err // Anderer Fehler
	}

	// Überprüfen, ob der Benutzer Lese-/Schreibzugriff hat
	file, err := os.Open(path)
	if err != nil {
		if os.IsPermission(err) {
			return false, nil // Keine Berechtigung
		}
		return false, err // Anderer Fehler
	}
	defer file.Close()
	return true, nil
}
