package network
import(
	"strings"
	"strconv"
	"fmt"
)

type Response interface{
	FromString(info string)

}
type  GameStat struct{
	GameName string
	GameId string
	NbrPlayers string

}
type GetGameResponse struct{
	Games []*GameStat
	PlayerId string
	
}
func (getGameResponse *GetGameResponse)FromString(info string){
	//reposne format : GetGames;2;1&game1&2|2&game2&3
	data := strings.Split(info,";")
	playerId := data[2]
	numberOfgames ,_ := strconv.Atoi(data[1])
	games := make([]*GameStat,0)
	if numberOfgames >0{
		gamesInfo := strings.Split(data[3],"|")
		for _,game:= range(gamesInfo){
			gameStat := &GameStat{}
			info := strings.Split(game,"&")
			gameStat.GameId = info[0]
			gameStat.GameName = info[1]
			gameStat.NbrPlayers = info[2]
			games = append(games,gameStat)
		}
	}
	getGameResponse.Games = games
	getGameResponse.PlayerId= playerId
}
type Data struct{
	PlayerId string
	PlayerX string
	PlayerY string
	PlayerRotationAngle string
	TorretX string
	TorretY string
	DidFire string
	BulletName string
}
func (data *Data)String()string{
	result := fmt.Sprintf("%s&%s&%s&%s&%s&%s&%s&",data.PlayerId,data.PlayerX, data.PlayerY, data.PlayerRotationAngle,data.TorretX,data.TorretY,data.DidFire)
	if data.DidFire == "1"{
		result += data.BulletName +"&"
	}
	return result
}
type InGameResponse struct{
	GameId string
	Data []*Data
	

}
func (inGameResponse *InGameResponse)FromString(info string){
	//reposne format : InGame;2;0&10&20&15|1&10&20&15

	data := strings.Split(info,";")
	inGameResponse.GameId = data[1]

	numberOfDataRecieved,_ := strconv.Atoi(data[2])
	dataList := make([]*Data,0)
	if numberOfDataRecieved > 0 {
		gamesInfo := strings.Split(data[3],"|")

		for _,game:= range(gamesInfo){
			data:=&Data{}
			info := strings.Split(game,"&")
			data.PlayerId = info[0]
			data.PlayerX = info[1]
			data.PlayerY = info[2]
			data.PlayerRotationAngle = info[3]
			data.TorretX = info[4]
			data.TorretY = info[5]
			data.DidFire = info[6]
			if data.DidFire == "1"{
				data.BulletName = info[7]
			}
			dataList = append(dataList,data)
		}

	}
	inGameResponse.Data = dataList
}

type CreateGameResponse struct{
	Game *GameStat
	PlayerId string

}
func (createGameResponse *CreateGameResponse)FromString(info string){
	data := strings.Split(info,";")
	game := &GameStat{
		GameId:data[1],
		GameName:fmt.Sprintf("Game%s",data[1]),
		NbrPlayers: "1",

	}
	createGameResponse.PlayerId = data[2]
	createGameResponse.Game = game

}

type JoinGameResponse struct{
	Player2Id string
	Player2Number string

}
func (joinGameResponse *JoinGameResponse)FromString(info string){
	data := strings.Split(info,";")
	joinGameResponse.Player2Id = data[1]
	joinGameResponse.Player2Number = data[2]
}
type CloseConnectionResponse struct{
	PlayerId string
}
func (closeConnectionResponse *CloseConnectionResponse)FromString(info string){
	data := strings.Split(info,";")
	closeConnectionResponse.PlayerId = data[1]

}