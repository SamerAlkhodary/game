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
	isOnline bool
	responseChannel chan Response
}
func CreateClient(bufferSize int, ip string, port string)*Client{

	return & Client{
		bufferSize:bufferSize,
		ip : ip,
		port: port,
		isOnline:false,
		responseChannel : make(chan Response,10),
	}
}
func (client *Client) Listen(playerId string){
	if playerId != "-1"{
		return
	}
	client.isOnline=true
	log.Println("Listening to the server")
	buffer :=  make([]byte, client.bufferSize)
	conn, err := net.Dial("udp", client.ip +":"+client.port)
	if err != nil {
		fmt.Printf("Some error %v", err)
   	return 
   }
   request := &GetGameRequest{
	   PlayerId: playerId,
   }
   fmt.Fprintf(conn, request.String())
   
   for client.isOnline{
	var response Response
	_, err = bufio.NewReader(conn).Read(buffer)
	if err != nil {
	fmt.Printf("Some error %v\n", err)
	} else {
	data := strings.TrimSpace(string(buffer))
	action := strings.Split(data,";")[0]
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
	
	case "ExitGame":
		response = &CloseConnectionResponse{}
	break;
	}
	response.FromString(data)
	client.responseChannel <- response	
	}
   }
	defer conn.Close()
}
func (client *Client) Send(request Request){
	conn, err := net.Dial("udp", client.ip +":"+client.port)
	if err != nil {
   		 fmt.Printf("Some error %v", err)
    return 
	}
	fmt.Fprintf(conn, request.String())
	defer conn.Close()
}
func (client *Client)GetResponse()Response{
	return <-client.responseChannel

}
func(client *Client)IsOnline()bool{
	return client.isOnline
}


