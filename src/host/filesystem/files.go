package filesystem

import (
	"os"
)

func FileExists(filePath string) bool {
	f, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	if f.IsDir() {
		return false
	}

	return true // Kein Fehler oder ein anderer Fehler, der nicht bedeutet, dass die Datei nicht existiert
}
