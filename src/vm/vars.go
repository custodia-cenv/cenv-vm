package vm

import (
	"sync"

	"github.com/custodia-cenv/bngsocket-go"
)

var (
	clientIpcVmSokcet *bngsocket.BngConn
	vmMutex           = new(sync.Mutex)
)
