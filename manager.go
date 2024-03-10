package pectx

import (
	"context"
)

type ctxKey string

type Manager struct {
	// The context key used to store the data in the context.
	// This is to avoid collisions with other packages.
	ctxKey ctxKey
}

// NewManager creates a new manager with the given context key.
func NewManager(uniqueKey string) *Manager {
	return &Manager{ctxKey: ctxKey(uniqueKey)}
}

// Get retrieves the values from the context.
func (m *Manager) Get(ctx context.Context) ([]KVStore, bool) {
	v := ctx.Value(m.ctxKey)

	if v == nil {
		return nil, false
	}

	return v.([]KVStore), true
}

// Set stores the values in the context.
// If the key already exists, the value is overwritten.
func (m *Manager) Set(ctx context.Context, fields ...KVStore) context.Context {
	fieldMap := map[string]string{}
	kvStores, exist := m.Get(ctx) // ensure the context is initialized
	if !exist {
		kvStores = []KVStore{}
	}

	// store the fields in a map to avoid duplicates
	allFields := append(kvStores, fields...)
	for _, kv := range allFields {
		key, value := kv.Get()
		fieldMap[key] = value
	}

	// convert the map back to a slice of KVStore to store in the context
	uniqueFields := []KVStore{}
	for key, value := range fieldMap {
		uniqueFields = append(uniqueFields, NewField(key, value))
	}

	return context.WithValue(ctx, m.ctxKey, uniqueFields)
}

// GetKeysAndValues retrieves the keys and values from the context.
// if the context is empty, it returns an empty slice.
func (m *Manager) GetKeysAndValues(ctx context.Context) []string {
	keysAndValues := []string{}
	fields, ok := m.Get(ctx)
	if !ok {
		return keysAndValues
	}
	for _, field := range fields {
		key, value := field.Get()
		keysAndValues = append(keysAndValues, key, value)
	}

	return keysAndValues
}
