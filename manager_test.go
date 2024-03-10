package pectx_test

import (
	"context"
	"reflect"
	"sort"
	"testing"

	"github.com/pragmaticengineering/pectx"
)

const contextKey = "test123123123"

func TestManager_Set(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc     string
		expected []string
		input    []pectx.KVStore
	}{
		{
			desc:     "Test that the context is set correctly",
			expected: []string{"k", "v"},
			input:    []pectx.KVStore{pectx.NewField("k", "v")},
		},
		{
			desc:     "Test if multiple fields are set correctly",
			expected: []string{"k", "v", "k2", "v2"},
			input:    []pectx.KVStore{pectx.NewField("k", "v"), pectx.NewField("k2", "v2")},
		},
		{
			desc:     "Test if the context is set correctly with no fields",
			expected: []string{},
			input:    []pectx.KVStore{},
		},
		{
			desc:     "Test if the values are overwritten if the key already exists",
			expected: []string{"k", "v2"},
			input:    []pectx.KVStore{pectx.NewField("k", "v"), pectx.NewField("k", "v2")},
		},
	}
	for _, tC := range testCases {
		testCase := tC
		t.Run(testCase.desc, func(t *testing.T) {
			t.Parallel()
			m := pectx.NewManager(contextKey)
			ctx := context.Background()
			ctx = m.Set(ctx, testCase.input...)
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

func TestGetKeysAndValues(t *testing.T) {
	t.Run("Test if the context is empty", func(t *testing.T) {
		t.Parallel()
		m := pectx.NewManager(contextKey)
		ctx := context.Background()
		keysAndValues := m.GetKeysAndValues(ctx)

		if len(keysAndValues) != 0 {
			t.Errorf("expected empty slice, got %v", keysAndValues)
		}

		if !reflect.DeepEqual([]string{}, keysAndValues) {
			t.Errorf("expected empty slice, got %v", keysAndValues)
		}
	})

}
