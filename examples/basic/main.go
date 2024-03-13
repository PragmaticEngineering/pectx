package main

import (
	"context"
	"fmt"

	"github.com/pragmaticengineering/pectx"
)

func main() {
	// create a new context manager
	// The contextkey is a string that is used to store and retrieve values from context
	// it should be unique to your application
	contextKey := "my-unique-context-key"
	m := pectx.NewManager(contextKey, &pectx.Store{})

	ctx := context.Background()

	// setting data utilizes the KVStore interface
	// This package provides a default implementation of KVStore called Store
	f := map[string]string{"my-key": "my-value"}

	// set the value in context
	// The set method returns a new context with the value set
	// The Set function can accept multiple values to set in context.
	ctx = m.Set(ctx, f)

	kvs := m.GetKeysAndValues(ctx)
	fmt.Println(kvs) // [my-key my-value]
}
