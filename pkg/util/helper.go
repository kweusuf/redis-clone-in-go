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

// LPush prepends one or multiple values to a list stored at the specified key.
// args[0]: key, args[1:]: values to prepend
// Returns the length of the list after the operation, or an error message if arguments are invalid.
func LPush(store *model.Store, args []string) string {
	if len(args) < 2 {
		return "ERROR: LPUSH command requires a key and at least one value"
	}
	key := args[0]
	values := args[1:]

	if store.List[key] == nil {
		store.List[key] = make([]string, 0)
	}

	// Prepend values in reverse order to maintain left-push order
	for i := len(values) - 1; i >= 0; i-- {
		store.List[key] = append([]string{values[i]}, store.List[key]...)
	}

	return fmt.Sprintf("%d", len(store.List[key]))
}

// RPush appends one or multiple values to a list stored at the specified key.
// args[0]: key, args[1:]: values to append
// Returns the length of the list after the operation, or an error message if arguments are invalid.
func RPush(store *model.Store, args []string) string {
	if len(args) < 2 {
		return "ERROR: RPUSH command requires a key and at least one value"
	}
	key := args[0]
	values := args[1:]

	if store.List[key] == nil {
		store.List[key] = make([]string, 0)
	}

	store.List[key] = append(store.List[key], values...)

	return fmt.Sprintf("%d", len(store.List[key]))
}

// LPop removes and returns the first element of the list stored at the specified key.
// args[0]: key
// Returns the value of the first element, or an error message if the list is empty or does not exist.
func LPop(store *model.Store, args []string) string {
	if len(args) < 1 {
		return "ERROR: LPOP command requires a key"
	}
	key := args[0]

	if len(store.List[key]) == 0 {
		return "ERROR: list is empty or does not exist"
	}

	value := store.List[key][0]
	store.List[key] = store.List[key][1:]

	return value
}

// RPop removes and returns the last element of the list stored at the specified key.
// args[0]: key
// Returns the value of the last element, or an error message if the list is empty or does not exist.
func RPop(store *model.Store, args []string) string {
	if len(args) < 1 {
		return "ERROR: RPOP command requires a key"
	}
	key := args[0]

	if len(store.List[key]) == 0 {
		return "ERROR: list is empty or does not exist"
	}

	value := store.List[key][len(store.List[key])-1]
	store.List[key] = store.List[key][:len(store.List[key])-1]

	return value
}

// LLen returns the length of the list stored at the specified key.
// args[0]: key
// Returns the length as a string, or "0" if the list does not exist.
func LLen(store *model.Store, args []string) string {
	if len(args) < 1 {
		return "ERROR: LLEN command requires a key"
	}
	key := args[0]

	if store.List[key] == nil {
		return "0"
	}

	return fmt.Sprintf("%d", len(store.List[key]))
}

// LIndex returns the element at the specified index in the list stored at the given key.
// args[0]: key, args[1]: index
// Returns the value at the index, or an error message if the index is out of range or arguments are invalid.
func LIndex(store *model.Store, args []string) string {
	if len(args) < 2 {
		return "ERROR: LINDEX command requires a key and an index"
	}
	key := args[0]
	index, err := strconv.Atoi(args[1])
	if err != nil {
		return "ERROR: index must be an integer"
	}

	if store.List[key] == nil || index < 0 || index >= len(store.List[key]) {
		return "ERROR: index out of range"
	}

	return store.List[key][index]
}
