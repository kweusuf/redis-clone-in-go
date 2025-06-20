package util

import (
	"fmt"
	"strconv"

	"github.com/kweusuf/redis-clone-in-go/pkg/model"
)

// Set stores the given value for the specified key in the store.
// args[0]: key, args[1]: value
// Returns the value that was set, or an error message if arguments are invalid.
func Set(store *model.Store, args []string) string {
	if len(args) < 2 {
		return "ERROR: SET command requires a key and a value"
	}
	key := args[0]
	value := args[1]
	store.Data[key] = value
	return value
}

// Get retrieves the value associated with the specified key from the store.
// args[0]: key
// Returns the value as a string, or an error message if arguments are invalid.
func Get(store *model.Store, args []string) string {
	if len(args) < 1 {
		return "ERROR: GET command requires a key"
	}
	key := args[0]
	return store.Data[key]
}

// Del removes the specified key and its value from the store.
// args[0]: key
// Does nothing if arguments are invalid.
func Del(store *model.Store, args []string) {
	if len(args) < 1 {
		return
	}
	key := args[0]
	delete(store.Data, key)
}

// Incr increments the integer value stored at the specified key by 1.
// args[0]: key
// Returns the new value as a string, or an error message if the value is not an integer or key does not exist.
func Incr(store *model.Store, args []string) string {
	if len(args) < 1 {
		return "ERROR: INCR command requires a key"
	}
	valStr := Get(store, args)
	if valStr == "" {
		return "ERROR: key does not exist"
	}
	value, err := strconv.Atoi(valStr)
	if err != nil {
		return "ERROR: value is not an integer"
	}
	newValue := value + 1
	newArgs := []string{args[0], strconv.Itoa(newValue)}
	Set(store, newArgs)
	return fmt.Sprintf("%v", newValue)
}

// Decr decrements the integer value stored at the specified key by 1.
// args[0]: key
// Returns the new value as a string, or an error message if the value is not an integer or key does not exist.
func Decr(store *model.Store, args []string) string {
	if len(args) < 1 {
		return "ERROR: DECR command requires a key"
	}
	valStr := Get(store, args)
	if valStr == "" {
		return "ERROR: key does not exist"
	}
	value, err := strconv.Atoi(valStr)
	if err != nil {
		return "ERROR: value is not an integer"
	}
	newValue := value - 1
	newArgs := []string{args[0], strconv.Itoa(newValue)}
	Set(store, newArgs)
	return fmt.Sprintf("%v", newValue)
}

// IncrBy increments the integer value stored at the specified key by the given increment.
// args[0]: key, args[1]: increment value
// Returns the new value as a string, or an error message if the value is not an integer or key does not exist.
func IncrBy(store *model.Store, args []string) string {
	if len(args) < 2 {
		return "ERROR: INCRBY command requires a key and an increment value"
	}
	increment, err := strconv.Atoi(args[1])
	if err != nil {
		return "ERROR: increment must be an integer"
	}
	valStr := Get(store, args)
	if valStr == "" {
		return "ERROR: key does not exist"
	}
	value, err := strconv.Atoi(valStr)
	if err != nil {
		return "ERROR: value is not an integer"
	}
	newValue := value + increment
	newArgs := []string{args[0], strconv.Itoa(newValue)}
	Set(store, newArgs)
	return fmt.Sprintf("%v", newValue)
}

// DecrBy decrements the integer value stored at the specified key by the given decrement.
// args[0]: key, args[1]: decrement value
// Returns the new value as a string, or an error message if the value is not an integer or key does not exist.
func DecrBy(store *model.Store, args []string) string {
	if len(args) < 2 {
		return "ERROR: DECRBY command requires a key and a decrement value"
	}
	decrement, err := strconv.Atoi(args[1])
	if err != nil {
		return "ERROR: decrement must be an integer"
	}
	valStr := Get(store, args)
	if valStr == "" {
		return "ERROR: key does not exist"
	}
	value, err := strconv.Atoi(valStr)
	if err != nil {
		return "ERROR: value is not an integer"
	}
	newValue := value - decrement
	newArgs := []string{args[0], strconv.Itoa(newValue)}
	Set(store, newArgs)
	return fmt.Sprintf("%v", newValue)
}
