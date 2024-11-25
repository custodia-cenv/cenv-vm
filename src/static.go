package cenvvm

const (
	Version VmVersion = 100000000

	// Gibt den Status des Core Osbjektes an
	NEW      VmState = 1
	INITED   VmState = 2
	SERVING  VmState = 3
	SHUTDOWN VmState = 4
	CLOSED   VmState = 5
)
