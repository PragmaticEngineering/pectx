package mapStore

// Map represents a key-value pair. It is used to store data in the context.
type Map map[string]string

// NewMap creates a new Store with the given key and value.
func NewMap(keyValue map[string]string) *Map {
	store := Map(keyValue)

	return &store
}

// listKeys retrieves the keys from the Store and returns them as a slice of strings.
func (f *Map) ListKeys() []string {
	keys := make([]string, 0, len(*f))
	for k := range *f {
		keys = append(keys, k)
	}
	return keys
}

// Get retrieves the value from the Store and returns it along with a boolean indicating if the value exists.
// If the value doesn't exist, the value is an empty string and the boolean is false.
func (f *Map) Get(key string) (value string) {
	return (*f)[key]
}

// Set stores the value in the Store.
// If the key already exists, the value is overwritten.
func (f *Map) Set(key string, value string) {
	(*f)[key] = value
}
