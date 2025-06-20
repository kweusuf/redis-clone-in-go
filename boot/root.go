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

// Init starts the TCP server, initializes the data store and service, and listens for incoming connections.
// It accepts connections on port 5001 and handles each connection in a separate goroutine.
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

// handleConnection processes commands from a single client connection.
// It reads input from the connection, parses commands and arguments, and writes responses back to the client.
//
// Parameters:
//   - conn: The network connection to the client.
//   - dbService: The database service used to handle commands.
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
