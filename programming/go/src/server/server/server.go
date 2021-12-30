package server
import (
	"log"
	"net"
	"strings"
	"fmt"
	"server/game"

)
type Server struct{
	maxBufferSize int
	port int
	ip string
	connection *net.UDPConn
	games []*game.Game

}

func InitServer(ip string, port int, maxBufferSize int)*Server{
	conn,err := net.ListenUDP("udp",&net.UDPAddr{
		Port:port,
		IP:net.ParseIP(ip),
	})
	if err !=nil{
		log.Println(err)
	}
	return &Server{
		maxBufferSize:maxBufferSize,
		port:port,
		ip:ip,
		connection:conn,
	}
}
func (server *Server) Listen(quit chan struct{}){
	
	log.Println("Server is running on address:",server.connection.LocalAddr().String())
	buffer := make([]byte, server.maxBufferSize)
	 
        n, remoteAddr, err := 0, new(net.UDPAddr), error(nil)
        for err == nil {
                n, remoteAddr, err = server.connection.ReadFromUDP(buffer)
				data := strings.TrimSpace(string(buffer[:n]))
				server.handleConnection(MakeRequest(remoteAddr,data),server.connection)
        }
        log.Println("listener failed - ", err)
        quit <- struct{}{}
}



func (server *Server)handleConnection(request *Request,connection *net.UDPConn){
	log.Println(request.Address.String(),request.PlayerId,request.GameId,request.Data)
	switch request.Action{
	case "GetGames":
		res :="games;"+ string(len(server.games))+";"
		if len(server.games)>0{
			for _,game := range(server.games){
				res += game.String()+"%"
			}
		}
		_, err := connection.WriteToUDP([]byte(res),request.Address)
		if err != nil {
			fmt.Printf("Some error %v\n", err)
		}
		break
	case "InGames":
		game := server.games[request.GameId]
		for _,player := range(game.Players){
			_, err := connection.WriteToUDP([]byte(fmt.Sprintf("%d;%d;%s,%s",request.PlayerId,request.GameId,"InGames",request.Data)),player.Address)
			if err != nil {
				fmt.Printf("Some error %v\n", err)
			}
		}
	break
	case "CreateGame":
		server.games = append(server.games, game.MakeGame(len(server.games),game.MakePlayer(request.Address,0)))
		_, err := connection.WriteToUDP([]byte(fmt.Sprintf("%d;%d;%s",0,len(server.games)-1,"CreateGame")),request.Address)
		if err != nil {
			fmt.Printf("Some error %v\n", err)
		}


	break
	case "JoinGame":
		chosenGame := server.games[request.GameId]
		chosenGame.Players = append(chosenGame.Players,game.MakePlayer(request.Address,request.PlayerId))
		for _,player := range(chosenGame.Players){
			_, err := connection.WriteToUDP([]byte(fmt.Sprintf("%d;%d;%s;%s",request.PlayerId,request.GameId,"JoinGame",request.Data)),player.Address)
			if err != nil {
				fmt.Printf("Some error %v\n", err)
			}
		}
		
	break


	}	
}
 
 
 