package service

import (
	"strings"

	"github.com/kweusuf/redis-clone-in-go/pkg/model"
)

type dbService struct {
	Store model.Store
}

type DBService interface {
	HandleCommand(command string, args []string) string
}

func MakeDBService(store model.Store) DBService {
	return &dbService{
		Store: store,
	}
}

// HandleCommand implements DBService.
func (d *dbService) HandleCommand(command string, args []string) string {
	switch command {
	case "SET":
		// Using Join to save the rest of the received data
		return d.set(args[0], strings.Join(args[1:], " "))
	case "GET":
		return d.get(args[0])
	case "DEL":
		d.del(args[0])
		return "DELETED"
	default:
		return "ERROR: Unknown command"
	}
}

func (d *dbService) set(key string, value string) string {
	d.Store.Data[key] = value

	return value
}

func (d *dbService) get(key string) string {
	return d.Store.Data[key]
}

func (d *dbService) del(key string) {
	delete(d.Store.Data, key)
}
