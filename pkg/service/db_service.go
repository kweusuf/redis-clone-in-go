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
// HandleCommand processes a given command with its arguments and performs the corresponding operation on the database store.
// Supported commands include:
//   - "SET": Sets the value for a given key.
//   - "GET": Retrieves the value for a given key.
//   - "DEL": Deletes the specified keys.
//   - "INCR": Increments the integer value of a key by one.
//   - "DECR": Decrements the integer value of a key by one.
//   - "INCRBY": Increments the integer value of a key by a specified amount.
//   - "DECRBY": Decrements the integer value of a key by a specified amount.
//
// Returns the result of the operation as a string, or an error message for unknown commands.
//
// Parameters:
//   - command: The command to execute (e.g., "SET", "GET").
//   - args: A slice of arguments for the command.
//
// Returns:
//   - A string representing the result of the command execution.
func (d *dbService) HandleCommand(command string, args []string) string {
	switch command {
	case "SET":
		// return util.Set(store, args[0], strings.Join(args[1:], " "))
		return util.Set(&d.Store, args)
	case "GET":
		return util.Get(&d.Store, args)
	case "DEL":
		util.Del(&d.Store, args)
		return "DELETED"
	case "INCR":
		return util.Incr(&d.Store, args)
	case "DECR":
		return util.Decr(&d.Store, args)
	case "INCRBY":
		return util.IncrBy(&d.Store, args)
	case "DECRBY":
		return util.DecrBy(&d.Store, args)
	default:
		return "ERROR: Unknown command"
	}
}
