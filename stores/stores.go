package stores

import (
	"github.com/pragmaticengineering/pectx"
	ms "github.com/pragmaticengineering/pectx/stores/map"
)

// compile-time check that Map implements KVStore
var _ pectx.KVStore = (*ms.Map)(nil)
