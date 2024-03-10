# Welcome to PragmaticEngineering/pectx
pectx is a context helper for golang. It is a simple and easy to use package that helps you to manage your context in a more pragmatic way.

- store and retrieve values from context in a type-safe way
    - avoids conflicting keys in context so your values are safe
- easily retrieve and set values in context

# How to use

## Set a value in context
```go
// create a new context manager
// The contextkey is a string that is used to store and retrieve values from context
// it should be unique to your application
contextKey := "my-unique-context-key"
m := pectx.NewManager(contextKey)

ctx := context.Background()

// setting data utilizes the KVStore interface
// This package provides a default implementation of KVStore called Field
f := pectx.NewField("my-key", "my-value")

// set the value in context
// The set method returns a new context with the value set
// The Set function can accept multiple values to set in context.
ctx := m.Set(ctx, f)
```

## Duplicate keys
Duplicate keys will be overwritten by the last value set in context.

```go
// create a new context manager
// The contextkey is a string that is used to store and retrieve values from context
// it should be unique to your application
contextKey := "my-unique-context-key"
m := pectx.NewManager(contextKey)

ctx := context.Background()

// setting data utilizes the KVStore interface
// This package provides a default implementation of KVStore called Field
f := pectx.NewField("my-key", "my-value")
f2 := pectx.NewField("my-key", "my-value2")
// set the value in context
// The set method returns a new context with the value set
// The Set function can accept multiple values to set in context.
ctx := m.Set(ctx, f, f2)

// The value of my-key will be "my-value2"
// The last value set will be the value of the key
// The order returned by GetKeysAndValues is not guaranteed as it uses a map to ensure unique keys
keysAndValues := m.GetKeysAndValues(ctx)

```


# FAQ
## How do I choose my keys?
Keys are fairly flexible, and can hold more or less any string value. For best compatibility with implementations and consistency with existing code in other projects, there are a few conventions you should consider.

## How should I make my keys 
- Make your keys human-readable.
- Constant keys are generally a good idea.
- Be consistent across your codebase.
- Keys should naturally match parts of the message string.
- Use lower case for simple keys and lowerCamelCase for more complex ones. Kubernetes is one example of a project that has adopted that convention.
- While key names are mostly unrestricted (and spaces are acceptable), it's generally a good idea to stick to printable ascii characters, or at least match the general character set of your log lines.

## Why should keys be constant values?
The point of structured logging is to make later log processing easier. Your keys are, effectively, the schema of each log message. If you use different keys across instances of the same log line, you will make your structured logs much harder to use. Sprintf() is for values, not for keys!