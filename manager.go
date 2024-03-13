package pectx

import (
	"context"
)

type ctxKey string

type Manager struct {
	// The context key used to store the data in the context.
	// This is to avoid collisions with other packages.
	ctxKey ctxKey
	// The store used to store the data.
	store KVStore
}

// NewManager creates a new manager with the given context key.
func NewManager(uniqueKey string, store KVStore) *Manager {
	return &Manager{
		ctxKey: ctxKey(uniqueKey),
		store:  store,
	}
}

// Get retrieves the values from the context.
func (m *Manager) Get(ctx context.Context) (KVStore, bool) {
	v := ctx.Value(m.ctxKey)

	if v == nil {
		return nil, false
	}

	return v.(KVStore), true
}

// Set stores the values in the context.
// If the key already exists, the value is overwritten.
func (m *Manager) Set(ctx context.Context, keyValues map[string]string) context.Context {
	kvStores, exist := m.Get(ctx) // ensure the context is initialized
	if !exist {
		kvStores = m.store
	}

	for key, value := range keyValues {
		kvStores.Set(key, value)
	}

	return context.WithValue(ctx, m.ctxKey, kvStores)
}

// GetKeysAndValues retrieves the keys and values from the context.
// if the context is empty, it returns an empty slice.
func (m *Manager) GetKeysAndValues(ctx context.Context) []string {

	fields, ok := m.Get(ctx)
	if !ok {
		return []string{}
	}

	keys := fields.ListKeys()
	keysAndValues := make([]string, 0, len(keys)*2)

	for _, key := range keys {
		value := fields.Get(key)
		keysAndValues = append(keysAndValues, key, value)
	}

	return keysAndValues
}
