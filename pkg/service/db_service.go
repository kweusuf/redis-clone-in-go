package service

import (
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
	//
	// Parameters:
	//   - command: The command to execute (e.g., "SET", "GET", "LPUSH").
	//   - args: A slice of arguments for the command.
	//
	// Returns:
	//   - A string representing the result of the command execution.
	HandleCommand(command string, args []string) string
}

// MakeDBService creates a new DBService with the provided store.
//
// Parameters:
//   - store: The underlying key-value data store.
//
// Returns:
//   - A DBService instance for handling database commands.
func MakeDBService(store model.Store) DBService {
	return &dbService{
		Store: store,
	}
}

// HandleCommand processes the given command and arguments, and returns the result as a string.
//
// Supported commands include:
//   - "SET": Sets the value for a given key.
//   - "GET": Retrieves the value for a given key.
//   - "DEL": Deletes the specified keys.
//   - "INCR": Increments the integer value of a key by one.
//   - "DECR": Decrements the integer value of a key by one.
//   - "INCRBY": Increments the integer value of a key by a specified amount.
//   - "DECRBY": Decrements the integer value of a key by a specified amount.
//   - "LPUSH": Prepends one or multiple values to a list.
//   - "RPUSH": Appends one or multiple values to a list.
//   - "LPOP": Removes and returns the first element of a list.
//   - "RPOP": Removes and returns the last element of a list.
//   - "LLEN": Returns the length of a list.
//   - "LINDEX": Returns the element at a given index in a list.
//
// Returns the result of the operation as a string, or an error message for unknown commands.
//
// Parameters:
//   - command: The command to execute (e.g., "SET", "GET", "LPUSH").
//   - args: A slice of arguments for the command.
//
// Returns:
//   - A string representing the result of the command execution.
func (d *dbService) HandleCommand(command string, args []string) string {
	var result string
	switch command {
	case "SET":
		result = util.Set(&d.Store, args)
	case "GET":
		result = util.Get(&d.Store, args)
	case "DEL":
		util.Del(&d.Store, args)
		result = "DELETED"
	case "INCR":
		result = util.Incr(&d.Store, args)
	case "DECR":
		result = util.Decr(&d.Store, args)
	case "INCRBY":
		result = util.IncrBy(&d.Store, args)
	case "DECRBY":
		result = util.DecrBy(&d.Store, args)
	case "LPUSH":
		result = util.LPush(&d.Store, args)
	case "RPUSH":
		result = util.RPush(&d.Store, args)
	case "LPOP":
		result = util.LPop(&d.Store, args)
	case "RPOP":
		result = util.RPop(&d.Store, args)
	case "LLEN":
		result = util.LLen(&d.Store, args)
	case "LINDEX":
		result = util.LIndex(&d.Store, args)
	default:
		result = "ERROR: Unknown command"
	}
	return result
}
