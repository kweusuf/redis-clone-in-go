package boot

import (
	"bufio"
	"fmt"
	"log/slog"
	"net"
	"strings"

	"github.com/kweusuf/redis-clone-in-go/pkg/model"
	"github.com/kweusuf/redis-clone-in-go/pkg/service"
)

func Init() {

	listener, err := net.Listen("tcp", ":5001")
	if err != nil {
		panic(err)
	}

	defer listener.Close()

	// Initialize the data store
	s := model.Store{Data: make(map[string]string)}

	// Initialize the service
	dbService := service.MakeDBService(s)

	for {
		conn, err := listener.Accept()
		if err != nil {
			slog.Error("error in receiving message", "err", err)
		}
		go handleConnection(conn, dbService)
	}
}

func handleConnection(conn net.Conn, dbService service.DBService) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		input := scanner.Text()
		parts := strings.Split(input, " ")

		if len(parts) < 2 {
			fmt.Fprintln(conn, "ERROR: Unknown command")
			continue
		}

		command := parts[0]
		args := parts[1:]

		response := dbService.HandleCommand(command, args)

		fmt.Fprintln(conn, response)
	}
}
