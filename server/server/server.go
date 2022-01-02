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
	onlinePlayers map[string]*game.Player
	connection *net.UDPConn
	games map[string]*game.Game
	idCounter int
	gameCounter int
}

func InitServer(ip string, port int, maxBufferSize int)*Server{
	players := make(map[string]*game.Player)
	games := make(map[string]*game.Game)
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
		games:games,
		connection:conn,
		onlinePlayers:players,
		idCounter :0,
		gameCounter:0,

	}
}
func (server *Server) addPlayer(playerId string, address  *net.UDPAddr )string{
	if _,ok := server.onlinePlayers[playerId];ok{
		return ""
	}
	id:=server.idCounter
	idStr := fmt.Sprintf("%d",id)
	server.onlinePlayers[idStr] = game.MakePlayer(address,idStr)
	server.idCounter+=1
	return idStr
}
func (server *Server) Listen(quit chan struct{}){
	
	log.Println("Server is running on address:",server.connection.LocalAddr().String())
	buffer := make([]byte, server.maxBufferSize)
	 
        n, remoteAddr, err := 0, new(net.UDPAddr), error(nil)
        for err == nil {
                n, remoteAddr, err = server.connection.ReadFromUDP(buffer)
				data := strings.TrimSpace(string(buffer[:n]))
				server.handleConnection(data,remoteAddr,server.connection)
        }
        log.Println("listener failed - ", err)
        quit <- struct{}{}
}


func (server *Server)handleConnection(info string,address *net.UDPAddr, connection *net.UDPConn){

	data := strings.Split(info,";")
	var request Request
	switch data[0]{
	case "GetGame":
		request = &GetGamesRequest{}
		request.FromString(info)
		req := request.(*GetGamesRequest)
		id := server.addPlayer(req.PlayerId,address)
		responseId :=""
		if id == ""{
			responseId= req.PlayerId
		}else{
			responseId = id
		}
		gameStats := make([]*GameStat,0)
		for _,game := range(server.games){
			gameStat:= &GameStat{
				GameName: game.Name,
				GameId: game.Id,
				NbrPlayers : fmt.Sprintf("%d",len(game.Players)),
			}
			gameStats = append(gameStats,gameStat)
		}
		
		
		log.Println("Online player:",server.idCounter)
		response := &GetGameResponse{
			Games: gameStats,
			PlayerId:responseId,
		}
		_, err := connection.WriteToUDP([]byte(response.String()),server.onlinePlayers[responseId].Address)
		if err != nil {
			fmt.Printf("Some error %v\n", err)
		}
		break
	case "InGame":
		request := &InGameRequest{}
		request.FromString(info)
		log.Println(request.GameId)
		game := server.games[request.GameId]
		for _,player := range(game.Players){
			if player.Id != request.PlayerId{
				log.Println("Sending ingame data to player:",player.Id)
				response := &InGameResponse{
					GameId: request.GameId,
					Data: request.Data,
				}
				_, err := connection.WriteToUDP([]byte(response.String()),player.Address)
				if err != nil {
					fmt.Printf("Some error %v\n", err)
				}
			}
		}
	break
	case "CreateGame":
		request = &CreateGameRequest{}
		request.FromString(info)
		req := request.(*CreateGameRequest)
		player := server.onlinePlayers[req.PlayerId]
		strGameId :=fmt.Sprintf("%d",server.gameCounter)
		game := game.MakeGame(strGameId,fmt.Sprintf("Game%s",strGameId),player)
		server.gameCounter+=1
		server.games[strGameId] = game
		response := &CreateGameResponse{
			GameId : strGameId,
			PlayerId : req.PlayerId,

		}
		_, err := connection.WriteToUDP([]byte(response.String()),player.Address)
		if err != nil {
			fmt.Printf("Some error %v\n", err)
		}


	break
	case "JoinGame":
		request := &JoinGameRequest{}
		request.FromString(info)
		chosenGame := server.games[request.GameId]
		player1 :=chosenGame.Players[0]
		player2 := server.onlinePlayers[request.PlayerId]
		chosenGame.Players = append(chosenGame.Players,player2)
		response := &JoinGameResponse{
			OtherPlayerId : player1.Id,
			OtherPlayerNumber : "1",
		}
		_, err := connection.WriteToUDP([]byte(response.String()),player2.Address)
		if err != nil {
			fmt.Printf("Some error %v\n", err)
		}
		response = &JoinGameResponse{
			OtherPlayerId : player2.Id,
			OtherPlayerNumber : "2",
		}
		_, err = connection.WriteToUDP([]byte(response.String()),player1.Address)
		if err != nil {
			fmt.Printf("Some error %v\n", err)
		}
		break
	case "ExitGame":
		request := &CloseConnectionRequest{}
		request.FromString(info)
		chosenGame := server.games[request.GameId]
		exitingPlayer := server.onlinePlayers[request.PlayerId]
		delete(server.onlinePlayers,request.PlayerId)
		players:= make([]*game.Player,0)
		for _,player := range(chosenGame.Players){
			if player.Id != exitingPlayer.Id{
				players = append(players,player)
				response := &CloseConnectionResponse{
					PlayerId : player.Id,
				}
				_, err := connection.WriteToUDP([]byte(response.String()),player.Address)
				if err != nil {
					fmt.Printf("Some error %v\n", err)
				}
			}
		}
		chosenGame.Players = players
		if len(chosenGame.Players) == 0{
			delete( server.games,request.GameId)
		}
	}	
}
 
 
 