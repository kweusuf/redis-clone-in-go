package service

import (
	"fmt"
	"strconv"
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
	case "INCR":
		if len(args) < 1 {
			return "ERROR: INCR command requires a key"
		}
		if _, err := strconv.Atoi(d.get(args[0])); err != nil {
			return "ERROR: value is not an integer"
		}
		return fmt.Sprintf("%v", d.incr(args[0]))
	case "DECR":
		if len(args) < 1 {
			return "ERROR: DECR command requires a key"
		}
		if _, err := strconv.Atoi(d.get(args[0])); err != nil {
			return "ERROR: value is not an integer"
		}
		return fmt.Sprintf("%v", d.decr(args[0]))
	case "INCRBY":
		if len(args) < 2 {
			return "ERROR: INCRBY command requires a key and an increment value"
		}
		increment, err := strconv.Atoi(args[1])
		if err != nil {
			return "ERROR: increment must be an integer"
		}
		result, err := d.incrBy(args[0], increment)
		if err != nil {
			return err.Error()
		}
		return fmt.Sprintf("%v", *result)
	case "DECRBY":
		if len(args) < 2 {
			return "ERROR: DECRBY command requires a key and a decrement value"
		}
		decrement, err := strconv.Atoi(args[1])
		if err != nil {
			return "ERROR: decrement must be an integer"
		}
		result, err := d.decrBy(args[0], decrement)
		if err != nil {
			return err.Error()
		}
		return fmt.Sprintf("%v", *result)
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

func (d *dbService) incr(key string) int {
	value, _ := strconv.Atoi(d.get(key))
	d.set(key, strconv.Itoa(value+1))
	return value + 1
}

func (d *dbService) decr(key string) int {
	value, _ := strconv.Atoi(d.get(key))
	d.set(key, strconv.Itoa(value-1))
	return value - 1
}

func (d *dbService) incrBy(key string, increment int) (*int, error) {
	valStr := d.get(key)
	value, err := strconv.Atoi(valStr)
	if err != nil {
		return nil, fmt.Errorf("ERROR: value is not an integer")
	}
	d.set(key, strconv.Itoa(value+increment))
	result := value + increment
	return &result, nil
}

func (d *dbService) decrBy(key string, decrement int) (*int, error) {
	valStr := d.get(key)
	value, err := strconv.Atoi(valStr)
	if err != nil {
		return nil, fmt.Errorf("ERROR: value is not an integer")
	}
	d.set(key, strconv.Itoa(value-decrement))
	result := value - decrement
	return &result, nil
}
