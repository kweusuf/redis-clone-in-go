package util

import "github.com/kweusuf/redis-clone-in-go/pkg/model"

// Set stores the given value for the specified key in the store.
func Set(store *model.Store, key string, value string) string {
	store.Data[key] = value
	return value
}

// Get retrieves the value associated with the specified key from the store.
func Get(store *model.Store, key string) string {
	return store.Data[key]
}

// Del removes the specified key and its value from the store.
func Del(store *model.Store, key string) {
	delete(store.Data, key)
}
