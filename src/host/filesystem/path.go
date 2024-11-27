package filesystem

import "path/filepath"

// Prüfen, ob ein Pfad existiert und ob er korrekt ist
func IsPathValid(path string) (bool, error) {
	_, err := filepath.Abs(path)
	if err != nil {
		return false, err // Pfad ist ungültig
	}
	return true, nil
}
