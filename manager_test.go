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

func TestManager_SetDuplicateKeys(t *testing.T) {
	t.Parallel()
	m := pectx.NewManager(contextKey, &pectx.Store{})
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
			m := pectx.NewManager(contextKey, &pectx.Store{})
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
	m := pectx.NewManager(contextKey, &pectx.Store{})
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
			m := pectx.NewManager(contextKey, &pectx.Store{})
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

func helperFields(amount int) map[string]string {

	store := map[string]string{}
	for i := 0; i < amount; i++ {
		key := fmt.Sprintf("k%d", i)
		store[key] = fmt.Sprintf("v%d", i)

	}
	return store
}

func BenchmarkManager_Set(b *testing.B) {
	var contextKey = "test123123123"

	testCases := []struct {
		name   string
		fields map[string]string
		mgr    *pectx.Manager
		ctx    context.Context
	}{
		{
			name:   "manager with 1 field",
			fields: helperFields(1),
			mgr:    pectx.NewManager(contextKey, &pectx.Store{}),
			ctx:    context.Background(),
		},
		{
			name:   "manager with 10 fields",
			fields: helperFields(10),
			mgr:    pectx.NewManager(contextKey, &pectx.Store{}),
			ctx:    context.Background(),
		},
		{
			name:   "manager with 100 fields",
			fields: helperFields(100),
			mgr:    pectx.NewManager(contextKey, &pectx.Store{}),
			ctx:    context.Background(),
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tc.mgr.Set(tc.ctx, tc.fields)
			}
		})
	}
}

func helperFieldsGet(amount int) (context.Context, *pectx.Manager) {
	ctx := context.Background()
	m := pectx.NewManager(contextKey, &pectx.Store{})

	m.Set(ctx, helperFields(amount))

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

func BenchmarkKeysAndValues_forloop(b *testing.B) {
	manager := pectx.NewManager("uniqueKey", &pectx.Store{})

	// Define test cases
	testCases := []struct {
		name string
		data context.Context
	}{
		{
			name: "TestCase1",
			data: manager.Set(context.Background(), map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			}),
		},
		{
			name: "TestCase2",
			data: manager.Set(context.Background(), map[string]string{
				"key4": "value4",
				"key5": "value5",
			}),
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				manager.GetKeysAndValues(tc.data)
			}
		})
	}
}
