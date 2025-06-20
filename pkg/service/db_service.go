package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kweusuf/redis-clone-in-go/pkg/model"
)

// dbService provides methods to interact with the key-value store.
type dbService struct {
	// Store is the underlying key-value data store.
	Store model.Store
}

// DBService defines the interface for handling database commands.
type DBService interface {
	// HandleCommand processes a command with arguments and returns the result as a string.
	HandleCommand(command string, args []string) string
}

// MakeDBService creates a new DBService with the provided store.
func MakeDBService(store model.Store) DBService {
	return &dbService{
		Store: store,
	}
}

// HandleCommand processes the given command and arguments, and returns the result as a string.
// Supported commands: SET, GET, DEL, INCR, DECR, INCRBY, DECRBY.
func (d *dbService) HandleCommand(command string, args []string) string {
	switch command {
	case "SET":
		return d.set(args[0], strings.Join(args[1:], " "))
	case "GET":
		return d.get(args[0])
	case "DEL":
		d.del(args[0])
		return "DELETED"
	case "INCR":
		return d.incr(args)
	case "DECR":
		return d.decr(args)
	case "INCRBY":
		return d.incrBy(args)
	case "DECRBY":
		return d.decrBy(args)
	default:
		return "ERROR: Unknown command"
	}
}

// set stores the given value for the specified key in the store.
// Returns the value that was set.
func (d *dbService) set(key string, value string) string {
	d.Store.Data[key] = value

	return value
}

// get retrieves the value associated with the specified key from the store.
// Returns the value as a string.
func (d *dbService) get(key string) string {
	return d.Store.Data[key]
}

// del removes the specified key and its value from the store.
func (d *dbService) del(key string) {
	delete(d.Store.Data, key)
}

// incr increments the integer value stored at the specified key by 1.
// Returns the new value as a string, or an error message if the value is not an integer.
func (d *dbService) incr(args []string) string {
	if len(args) < 1 {
		return "ERROR: INCR command requires a key"
	}
	if _, err := strconv.Atoi(d.get(args[0])); err != nil {
		return "ERROR: value is not an integer"
	}
	return fmt.Sprintf("%v", d.increment(args[0]))
}

// decr decrements the integer value stored at the specified key by 1.
// Returns the new value as a string, or an error message if the value is not an integer.
func (d *dbService) decr(args []string) string {
	if len(args) < 1 {
		return "ERROR: DECR command requires a key"
	}
	if _, err := strconv.Atoi(d.get(args[0])); err != nil {
		return "ERROR: value is not an integer"
	}
	return fmt.Sprintf("%v", d.decrement(args[0]))
}

// incrBy increments the integer value stored at the specified key by the given increment.
// Returns the new value as a string, or an error message if the value is not an integer.
func (d *dbService) incrBy(args []string) string {
	if len(args) < 2 {
		return "ERROR: INCRBY command requires a key and an increment value"
	}
	increment, err := strconv.Atoi(args[1])
	if err != nil {
		return "ERROR: increment must be an integer"
	}
	result, err := d.incrementBy(args[0], increment)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("%v", *result)
}

// decrBy decrements the integer value stored at the specified key by the given decrement.
// Returns the new value as a string, or an error message if the value is not an integer.
func (d *dbService) decrBy(args []string) string {
	if len(args) < 2 {
		return "ERROR: DECRBY command requires a key and a decrement value"
	}
	decrement, err := strconv.Atoi(args[1])
	if err != nil {
		return "ERROR: decrement must be an integer"
	}
	result, err := d.decrementBy(args[0], decrement)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("%v", *result)
}

// increment increases the integer value at the given key by 1.
// Returns the new value as an int.
func (d *dbService) increment(key string) int {
	value, _ := strconv.Atoi(d.get(key))
	d.set(key, strconv.Itoa(value+1))
	return value + 1
}

// decrement decreases the integer value at the given key by 1.
// Returns the new value as an int.
func (d *dbService) decrement(key string) int {
	value, _ := strconv.Atoi(d.get(key))
	d.set(key, strconv.Itoa(value-1))
	return value - 1
}

// incrementBy increases the integer value at the given key by the specified increment.
// Returns the new value as a pointer to int, or an error if the value is not an integer.
func (d *dbService) incrementBy(key string, increment int) (*int, error) {
	valStr := d.get(key)
	value, err := strconv.Atoi(valStr)
	if err != nil {
		return nil, fmt.Errorf("ERROR: value is not an integer")
	}
	d.set(key, strconv.Itoa(value+increment))
	result := value + increment
	return &result, nil
}

// decrementBy decreases the integer value at the given key by the specified decrement.
// Returns the new value as a pointer to int, or an error if the value is not an integer.
func (d *dbService) decrementBy(key string, decrement int) (*int, error) {
	valStr := d.get(key)
	value, err := strconv.Atoi(valStr)
	if err != nil {
		return nil, fmt.Errorf("ERROR: value is not an integer")
	}
	d.set(key, strconv.Itoa(value-decrement))
	result := value - decrement
	return &result, nil
}
