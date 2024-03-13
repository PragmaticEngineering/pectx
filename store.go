package pectx

// This line of code doesn't do anything at runtime, but if Store does not implement KVStore,
// the code will not compile.
// This is a common idiom in Go for checking interface implementation at compile time.
var _ KVStore = (*Store)(nil)

// Store is an interface that defines the methods to store and retrieve data from a Store.
type KVStore interface {
	KVGetter
	KVSetter
}

// KVGetter is an interface that defines the method to retrieve data from a Store.
type KVGetter interface {
	Get(key string) (value string)
	ListKeys() []string
}

// KVSetter is an interface that defines the method to store data in a Store.
type KVSetter interface {
	Set(key string, value string)
}

// Store represents a key-value pair. It is used to store data in the context.
type Store map[string]string

// NewStore creates a new Store with the given key and value.
func NewStore(keyValue map[string]string) *Store {
	store := Store(keyValue)
	return &store
}

// listKeys retrieves the keys from the Store and returns them as a slice of strings.
func (f *Store) ListKeys() []string {
	keys := make([]string, 0, len(*f))
	for k := range *f {
		keys = append(keys, k)
	}
	return keys
}

// Get retrieves the value from the Store and returns it along with a boolean indicating if the value exists.
// If the value doesn't exist, the value is an empty string and the boolean is false.
func (f *Store) Get(key string) (value string) {
	return (*f)[key]
}

// Set stores the value in the Store.
// If the key already exists, the value is overwritten.
func (f *Store) Set(key string, value string) {
	(*f)[key] = value
}
