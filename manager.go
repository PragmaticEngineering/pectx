package pectx

import (
	"context"

	mapStore "github.com/pragmaticengineering/pectx/stores/map"
)

// Store is an interface that defines the methods to store and retrieve data from a Store.
type KVStore interface {
	KVGetter
	KVSetter
}

// KVGetter is an interface that defines the method to retrieve data from a Store.
type KVGetter interface {
	Get(key string) (value string)
	ListKeys() []string
}

// KVSetter is an interface that defines the method to store data in a Store.
type KVSetter interface {
	Set(key string, value string)
}

type ctxKey string

// ManagerOption is a function that sets some option on the manager.
type ManagerOption func(*Manager)

type Manager struct {
	// The context key used to store the data in the context.
	// This is to avoid collisions with other packages.
	ctxKey ctxKey
	// The store used to store the data.
	store KVStore
}

// WithStore changes the store used by the manager.
func WithStore(store KVStore) ManagerOption {
	return func(m *Manager) {
		m.store = store
	}
}

// NewManager creates a new manager with the given context key.
func NewManager(uniqueKey string, opts ...ManagerOption) *Manager {
	m := &Manager{
		ctxKey: ctxKey(uniqueKey),
		store:  &mapStore.Map{},
	}
	for _, opt := range opts {
		opt(m)
	}

	return m
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
