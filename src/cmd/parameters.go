package main

import (
	"flag"
	"fmt"
	"strings"
)

// VFSPair repräsentiert ein Paar aus Typ und Datei
type VFSPair struct {
	Type string
	File string
}

// VFSFlags ist ein Slice von VFSPair und implementiert die flag.Value Schnittstelle
type VFSFlags []VFSPair

// String gibt eine kommagetrennte Liste der VFS-Paare zurück
func (v *VFSFlags) String() string {
	var pairs []string
	for _, pair := range *v {
		pairs = append(pairs, pair.Type+" "+pair.File)
	}
	return strings.Join(pairs, ", ")
}

// Set fügt ein neues VFS-Paar hinzu
func (v *VFSFlags) Set(value string) error {
	// Teile den Eingabewert in Typ und Datei
	parts := strings.SplitN(value, " ", 2)
	if len(parts) != 2 {
		return fmt.Errorf("ungültiges VFS-Paar: %s. Erwartet Format: 'type file'", value)
	}

	// Erstelle ein neues Paar und füge es dem Slice hinzu
	pair := VFSPair{
		Type: parts[0],
		File: parts[1],
	}
	*v = append(*v, pair)
	return nil
}

var (
	// Definiere die Flags
	image    = flag.String("img", "", "Pfad zur Bilddatei (erforderlich)")
	vfsFlags VFSFlags
)

func init() {
	// Registriere das benutzerdefinierte -vfs Flag
	flag.Var(&vfsFlags, "vfs", "VFS-Paare im Format 'type file'. Kann mehrmals verwendet werden.")
}
