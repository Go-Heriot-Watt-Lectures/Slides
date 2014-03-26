import (
	"net"
	"log"
)

func main() {
	handle, err := net.Listen("tcp", "localhost:3048")
	if err != nil {
		panic("cannot listen")
	}
	for {
		connection, err := handle.Accept()
		if err != nil {
			log.Println("Error handling a connection")
		}
		go handleConnection(connection)
	}
}
