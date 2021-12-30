package game
import(
	"fmt"
)
type Game struct{
	Players [] *Player
	Id int
}
func MakeGame(id int,player *Player)*Game{
	players:=make([]*Player,0)
	players = append(players,player)
	return &Game{
		Id : id,
		Players : players,
	}
}
func (game *Game)EnoughPlayers() bool{
	return len(game.Players)>1
}
func (game *Game)String()string{
	return fmt.Sprintf("%d,%d",game.Id,len(game.Players))
}


package game
import(
	"net"

)
type Player struct{
	Address  *net.UDPAddr
	id int
}
func MakePlayer(address *net.UDPAddr, id int)*Player{
	return &Player{
		Address:address,
		id:id,
	}

}

package server
import(
	//"fmt"
	"strings"
	"strconv"
)

type Request interface{
	FromString(info string)
	
}
type Data struct{
	PlayerId string
	PlayerX string
	PlayerY string
	PlayerRotationAngle string
}
type GetGamesRequest struct{
	PlayerId string
}
func (getGameRequest *GetGamesRequest) FromString(info string){
	data := strings.Split(info,";")
	getGameRequest.PlayerId = data[1]
}

type InGameRequest struct{
	PlayerId string
	GameId string
	Action string
	NumberOfData string
	Data []*Data
}
func (InGameRequest *InGameRequest)FromString(info string){
	//reposne format : InGame;2;0&10&20&15|1&10&20&15

	data := strings.Split(info,";")
	InGameRequest.GameId = data[1]

	numberOfDataRecieved,_ := strconv.Atoi(data[3])
	dataList := make([]*Data,0)
	if numberOfDataRecieved > 0 {
		gamesInfo := strings.Split(data[4],"|")

		for _,game:= range(gamesInfo){
			data:=&Data{}
			info := strings.Split(game,"&")
			data.PlayerId = info[0]
			data.PlayerX = info[1]
			data.PlayerY = info[2]
			data.PlayerRotationAngle = info[3]
			dataList = append(dataList,data)
		}

	}
	InGameRequest.Data = dataList
}

type CreateGameRequest struct{
	PlayerId string

}
func (createGameRequest *CreateGameRequest) FromString(info string){
	data := strings.Split(info,";")
	createGameRequest.PlayerId = data[1]
}

type JoinGameRequest struct{
	PlayerId string
	Action string
	GameId string
	Name string

}
func (joinGameRequest *JoinGameRequest)FromString(info string){
	data := strings.Split(info,";")
	joinGameRequest.PlayerId = data[1]
	joinGameRequest.GameId = data[2]
	joinGameRequest.Name = data[3]
}




package server
import(
	//"fmt"
	"strings"
	"strconv"
)

type Request interface{
	FromString(info string)
	
}
type Data struct{
	PlayerId string
	PlayerX string
	PlayerY string
	PlayerRotationAngle string
}
type GetGamesRequest struct{
	PlayerId string
}
func (getGameRequest *GetGamesRequest) FromString(info string){
	data := strings.Split(info,";")
	getGameRequest.PlayerId = data[1]
}

type InGameRequest struct{
	PlayerId string
	GameId string
	Action string
	NumberOfData string
	Data []*Data
}
func (InGameRequest *InGameRequest)FromString(info string){
	//reposne format : InGame;2;0&10&20&15|1&10&20&15

	data := strings.Split(info,";")
	InGameRequest.GameId = data[1]

	numberOfDataRecieved,_ := strconv.Atoi(data[3])
	dataList := make([]*Data,0)
	if numberOfDataRecieved > 0 {
		gamesInfo := strings.Split(data[4],"|")

		for _,game:= range(gamesInfo){
			data:=&Data{}
			info := strings.Split(game,"&")
			data.PlayerId = info[0]
			data.PlayerX = info[1]
			data.PlayerY = info[2]
			data.PlayerRotationAngle = info[3]
			dataList = append(dataList,data)
		}

	}
	InGameRequest.Data = dataList
}

type CreateGameRequest struct{
	PlayerId string

}
func (createGameRequest *CreateGameRequest) FromString(info string){
	data := strings.Split(info,";")
	createGameRequest.PlayerId = data[1]
}

type JoinGameRequest struct{
	PlayerId string
	Action string
	GameId string
	Name string

}
func (joinGameRequest *JoinGameRequest)FromString(info string){
	data := strings.Split(info,";")
	joinGameRequest.PlayerId = data[1]
	joinGameRequest.GameId = data[2]
	joinGameRequest.Name = data[3]
}


