package pectx_test

import (
	"context"
	"fmt"
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

func helperFields(amount int) []pectx.KVStore {
	var fields []pectx.KVStore
	for i := 0; i < amount; i++ {
		fields = append(fields, pectx.NewField(fmt.Sprintf("k%d", i), fmt.Sprintf("v%d", i)))
	}
	return fields
}

func helperFieldsDuplicates(amount int) []pectx.KVStore {
	var fields []pectx.KVStore
	for i := 0; i < amount; i++ {
		fields = append(fields, pectx.NewField(fmt.Sprintf("k%d", i), fmt.Sprintf("v%d", i)))
	}
	return fields
}

func BenchmarkManager_Set(b *testing.B) {
	var contextKey = "test123123123"

	testCases := []struct {
		name   string
		fields []pectx.KVStore
		mgr    *pectx.Manager
		ctx    context.Context
	}{
		{
			name:   "manager with 1 field",
			fields: helperFields(1),
			mgr:    pectx.NewManager(contextKey),
			ctx:    context.Background(),
		},
		{
			name:   "manager with 10 fields",
			fields: helperFields(10),
			mgr:    pectx.NewManager(contextKey),
			ctx:    context.Background(),
		},
		{
			name:   "manager with 100 fields",
			fields: helperFields(100),
			mgr:    pectx.NewManager(contextKey),
			ctx:    context.Background(),
		},
		{
			name:   "manager with 10 fields - All Duplicate",
			fields: helperFieldsDuplicates(10),
			mgr:    pectx.NewManager(contextKey),
			ctx:    context.Background(),
		},
		{
			name:   "manager with 100 fields - All Duplicate",
			fields: helperFieldsDuplicates(100),
			mgr:    pectx.NewManager(contextKey),
			ctx:    context.Background(),
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tc.mgr.Set(tc.ctx, tc.fields...)
			}
		})
	}
}

func helperFieldsGet(amount int) (context.Context, *pectx.Manager) {
	ctx := context.Background()
	m := pectx.NewManager(contextKey)

	m.Set(ctx, helperFields(amount)...)

	return ctx, m
}

func BenchmarkManager_Get(b *testing.B) {
	testCases := []struct {
		name   string
		amount int
	}{
		{
			name:   "context with 1 field",
			amount: 1,
		},
		{
			name:   "context with 10 fields",
			amount: 10,
		},
		{
			name:   "context with 100 fields",
			amount: 100,
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			ctx, m := helperFieldsGet(tc.amount)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				m.Get(ctx)
			}
		})
	}
}

func BenchmarkManager_GetKeysAndValues(b *testing.B) {
	testCases := []struct {
		name   string
		amount int
	}{
		{
			name:   "context with 1 field",
			amount: 1,
		},
		{
			name:   "context with 10 fields",
			amount: 10,
		},
		{
			name:   "context with 100 fields",
			amount: 100,
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			ctx, m := helperFieldsGet(tc.amount)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				m.GetKeysAndValues(ctx)
			}
		})
	}
}
