package network
import (
    "fmt"
    "net"
    "bufio"
	"strings"
	"log"
)
type Client struct{
	bufferSize int
	ip string
	port string
}
func CreateClient(bufferSize int, ip string, port string)*Client{

	return & Client{
		bufferSize:bufferSize,
		ip : ip,
		port: port,
	}
}
func (client *Client) SendAndReceive(request Request)Response{
	buffer :=  make([]byte, client.bufferSize)
	conn, err := net.Dial("udp", client.ip +":"+client.port)
	if err != nil {
   		 fmt.Printf("Some error %v", err)
    return nil
	}
	var response Response
	log.Println("sent:",request.String())
	fmt.Fprintf(conn, request.String())
	_, err = bufio.NewReader(conn).Read(buffer)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
	} else {
		data := strings.TrimSpace(string(buffer))
		log.Println("recievied:",data)

		action := strings.Split(data,";")[0]
		log.Println("action:",action)
		switch action{
		case "GetGame":
			response = &GetGameResponse{}
			break;

		case "InGame":
			response = &InGameResponse{}
			break;

		case "CreateGame":
			response = &CreateGameResponse{}
			break;

		case "JoinGame":
			response = &JoinGameResponse{}
			break;
		}
		response.FromString(data)
		log.Println("recievied;",response)

		
	}
	conn.Close()
	return response
}


