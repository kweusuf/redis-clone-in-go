package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kweusuf/redis-clone-in-go/pkg/model"
	"github.com/kweusuf/redis-clone-in-go/pkg/util"
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
		return util.Set(&d.Store, args[0], strings.Join(args[1:], " "))
	case "GET":
		return util.Get(&d.Store, args[0])
	case "DEL":
		util.Del(&d.Store, args[0])
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

// incr increments the integer value stored at the specified key by 1.
// Returns the new value as a string, or an error message if the value is not an integer.
func (d *dbService) incr(args []string) string {
	if len(args) < 1 {
		return "ERROR: INCR command requires a key"
	}
	valStr := util.Get(&d.Store, args[0])
	if valStr == "" {
		return "ERROR: key does not exist"
	}
	value, err := strconv.Atoi(valStr)
	if err != nil {
		return "ERROR: value is not an integer"
	}
	newValue := value + 1
	util.Set(&d.Store, args[0], strconv.Itoa(newValue))
	return fmt.Sprintf("%v", newValue)
}

// decr decrements the integer value stored at the specified key by 1.
// Returns the new value as a string, or an error message if the value is not an integer.
func (d *dbService) decr(args []string) string {
	if len(args) < 1 {
		return "ERROR: DECR command requires a key"
	}
	valStr := util.Get(&d.Store, args[0])
	if valStr == "" {
		return "ERROR: key does not exist"
	}
	value, err := strconv.Atoi(valStr)
	if err != nil {
		return "ERROR: value is not an integer"
	}
	newValue := value - 1
	util.Set(&d.Store, args[0], strconv.Itoa(newValue))
	return fmt.Sprintf("%v", newValue)
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
	valStr := util.Get(&d.Store, args[0])
	if valStr == "" {
		return "ERROR: key does not exist"
	}
	value, err := strconv.Atoi(valStr)
	if err != nil {
		return "ERROR: value is not an integer"
	}
	newValue := value + increment
	util.Set(&d.Store, args[0], strconv.Itoa(newValue))
	return fmt.Sprintf("%v", newValue)
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
	valStr := util.Get(&d.Store, args[0])
	if valStr == "" {
		return "ERROR: key does not exist"
	}
	value, err := strconv.Atoi(valStr)
	if err != nil {
		return "ERROR: value is not an integer"
	}
	newValue := value - decrement
	util.Set(&d.Store, args[0], strconv.Itoa(newValue))
	return fmt.Sprintf("%v", newValue)
}
