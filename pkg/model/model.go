package model

// Store represents a simple in-memory key-value store.
type Store struct {
	// Data holds the key-value pairs for the store.
	Data map[string]string
	List map[string][]string
}
