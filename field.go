package pectx

// Store is an interface that defines the methods to store and retrieve data from a field.
type KVStore interface {
	KVGetter
	KVSetter
}

// KVGetter is an interface that defines the method to retrieve data from a field.
type KVGetter interface {
	Get() (key string, value string)
}

// KVSetter is an interface that defines the method to store data in a field.
type KVSetter interface {
	Set(key string, value string)
}

// Field represents a key-value pair. It is used to store data in the context.
type Field struct {
	key   string
	value string
}

// NewField creates a new field with the given key and value.
func NewField(key string, value string) *Field {
	return &Field{
		key:   key,
		value: value,
	}
}

// Get retrieves the value from the field and returns it along with a boolean indicating if the value exists.
// If the value doesn't exist, the value is an empty string and the boolean is false.
func (f *Field) Get() (key string, value string) {
	return f.key, f.value
}

// Set stores the value in the field.
// If the key already exists, the value is overwritten.
func (f *Field) Set(key string, value string) {
	f.key = key
	f.value = value
}
