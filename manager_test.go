package pectx_test

import (
	"context"
	"reflect"
	"sort"
	"testing"

	"github.com/pragmaticengineering/pectx"
)

const contextKey = "test123123123"

func TestManager_SetDuplicateKeys(t *testing.T) {
	t.Parallel()
	m := pectx.NewManager(contextKey)
	ctx := context.Background()
	f := map[string]string{"k": "v"}
	ctx = m.Set(ctx, f)
	f2 := map[string]string{"k": "v2"}
	ctx = m.Set(ctx, f2)
	store, _ := m.Get(ctx)

	if store.Get("k") != "v2" {
		t.Errorf("expected v2, got %s", store.Get("k"))
	}
}

func TestManager_Set(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc     string
		expected []string
		input    map[string]string
	}{
		{
			desc:     "Test that the context is set correctly",
			expected: []string{"k", "v"},
			input:    map[string]string{"k": "v"},
		},
		{
			desc:     "Test if multiple fields are set correctly",
			expected: []string{"k", "v", "k2", "v2"},
			input:    map[string]string{"k": "v", "k2": "v2"},
		},
		{
			desc:     "Test if the context is set correctly with no fields",
			expected: []string{},
			input:    map[string]string{},
		},
	}
	for _, tC := range testCases {
		testCase := tC
		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()
			m := pectx.NewManager(contextKey)
			ctx := context.Background()
			ctx = m.Set(ctx, testCase.input)
			keysAndValues := m.GetKeysAndValues(ctx)
			// sort the slices to compare them, as the order of the keys and values is not guaranteed.
			sort.Strings(keysAndValues)     // Sort the keysAndValues slice
			sort.Strings(testCase.expected) // Sort the expected slice
			if !reflect.DeepEqual(testCase.expected, keysAndValues) {
				t.Errorf("expected %v, got %v", testCase.expected, keysAndValues)
			}
		})
	}
}

func TestGetKeysAndValuesEmpty(t *testing.T) {
	t.Parallel()
	m := pectx.NewManager(contextKey)
	ctx := context.Background()
	keysAndValues := m.GetKeysAndValues(ctx)
	if len(keysAndValues) != 0 {
		t.Errorf("expected 0, got %d", len(keysAndValues))
	}
}

func TestGetKeysAndValues(t *testing.T) {

	tc := []struct {
		desc     string
		expected []string
		input    map[string]string
	}{
		{
			desc:     "Test if the context is set correctly",
			expected: []string{"k", "v"},
			input:    map[string]string{"k": "v"},
		},
		{
			desc:     "Test if multiple fields are set correctly",
			expected: []string{"k", "v", "k2", "v2"},
			input:    map[string]string{"k": "v", "k2": "v2"},
		},
		{
			desc:     "Test if the context is set correctly with no fields",
			expected: []string{},
			input:    map[string]string{},
		},
	}

	for _, tC := range tc {
		testCase := tC
		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()
			m := pectx.NewManager(contextKey)
			ctx := context.Background()
			ctx = m.Set(ctx, testCase.input)
			keysAndValues := m.GetKeysAndValues(ctx)
			// sort the slices to compare them, as the order of the keys and values is not guaranteed.
			sort.Strings(keysAndValues)     // Sort the keysAndValues slice
			sort.Strings(testCase.expected) // Sort the expected slice
			if !reflect.DeepEqual(testCase.expected, keysAndValues) {
				t.Errorf("expected %v, got %v", testCase.expected, keysAndValues)
			}
		})
	}

}

type mockStore struct {
	k string
	v string
}

func (m *mockStore) Get(k string) string {
	return m.v
}

func (m *mockStore) Set(k, v string) {
	m.k = k
	m.v = v
}

func (m *mockStore) ListKeys() []string {
	return []string{"mockKey", "mockValue"}
}

var _ pectx.KVStore = &mockStore{}

func TestWithStore(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc    string
		options []pectx.ManagerOption
		expectedKV []string
	}{
		{
			desc:    "Test that the store is set correctly",
			options: []pectx.ManagerOption{pectx.WithStore(&mockStore{})},
			expectedKV: []string{"mockKey", "mockValue"},
		},
	}

	for _, tC := range testCases {
		testCase := tC
		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()
			m := pectx.NewManager(contextKey, testCase.options...)
			if m == nil {
				t.Errorf("expected a manager, got nil")
			}

			ctx := m.Set(context.Background(), map[string]string{"k": "v"})
			ctx = m.Set(ctx, map[string]string{"k": "v2"})
			kv := m.GetKeysAndValues(ctx)
			sort.Strings(kv)
			sort.Strings(testCase.expectedKV)
			if !reflect.DeepEqual(testCase.expectedKV, kv){
				t.Errorf("expected %v, got %v", []string{"mockKey", "mockValue"}, kv)
			}
		})
	}
}
