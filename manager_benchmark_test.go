package pectx_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/pragmaticengineering/pectx"
)

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
	m := pectx.NewManager(contextKey)

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
	manager := pectx.NewManager(contextKey)

	// Define test cases
	testCases := []struct {
		name string
		data map[string]string
	}{
		{
			name: "Set 1 key-value pair",
			data: map[string]string{
				"key1": "value1",
			},
		},
		{
			name: "Set 10 key-value pairs",
			data: helperFields(10),
		},
		{
			name: "Set 100 key-value pairs",
			data: helperFields(100),
		},
	}

	for _, tc := range testCases {
		testcase := tc
		b.Run(testcase.name, func(b *testing.B) {
			ctx := manager.Set(context.Background(), testcase.data)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				manager.GetKeysAndValues(ctx)
			}
		})
	}
}
