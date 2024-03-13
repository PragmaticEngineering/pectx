package main

import (
	"context"
	"fmt"
	pectx "github.com/pragmaticengineering/pectx"
)

func main() {
	// create a new context manager
	// The contextkey is a string that is used to store and retrieve values from context
	// it should be unique to your application
	contextKey := "my-unique-context-key"
	m := pectx.NewManager(contextKey)

	ctx := context.Background()

	// setting data utilizes the KVStore interface
	// This package provides a default implementation of KVStore called Field
	f := pectx.NewStore(map[string]string{"my-key": "my-value", "my-key2": "my-value"})
	// set the value in context
	// The set method returns a new context with the value set
	// The Set function can accept multiple values to set in context.
	ctx = m.Set(ctx, f)
	f2 := pectx.NewStore(map[string]string{"my-key": "my-value2", "my-key2": "my-value2"})
	ctx = m.Set(ctx, f2)
	// The value of my-key will be "my-value2"
	// The last value set will be the value of the key
	// The order returned by GetKeysAndValues is not guaranteed as it uses a map to ensure unique keys
	keysAndValues := m.GetKeysAndValues(ctx)
	fmt.Println(keysAndValues) // [my-key my-value2 my-key2 my-value2]
}
