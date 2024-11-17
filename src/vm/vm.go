package vm

import (
	"github.com/custodia-cenv/cenvx-core/src/core"
)

func InitVmProcessInstance() error {
	initCoreVmIpcClientSession([]*core.ACL{})
	return nil
}
