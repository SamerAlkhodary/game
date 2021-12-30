package network
import(
	"strings"
)

type Response interface{
	FromString(info string)

}
type GetGameResponse struct{
	GameName string
	Action string
	GameId string
	NbrPlayers string
}
func (getGameResponse *GetGameResponse)FromString(info string){
	data := strings.Split(info,";")
	getGameResponse.Action = data[0]
	getGameResponse.GameName = data[1]
	getGameResponse.GameId = data[2]
	getGameResponse.NbrPlayers = data[3]
}

type InGameResponse struct{
	PlayerId string
	GameId string
	Action string
	Data string

}
func (inGameResponse *InGameResponse)FromString(info string){
	data := strings.Split(info,";")
	inGameResponse.Action = data[0]
	inGameResponse.PlayerId = data[1]
	inGameResponse.GameId = data[2]
	inGameResponse.Data = data[3]
}

type CreateGameResponse struct{
	GameId string
	Action string

}
func (createGameResponse *CreateGameResponse)FromString(info string){
	data := strings.Split(info,";")
	createGameResponse.Action = data[0]
	createGameResponse.GameId = data[1]
}

type JoinGameResponse struct{
	PlayerId string
	Action string
	GameId string
	Data string

}
func (joinGameResponse *JoinGameResponse)FromString(info string){
	data := strings.Split(info,";")
	joinGameResponse.Action = data[0]
	joinGameResponse.PlayerId = data[1]
	joinGameResponse.GameId = data[2]
	joinGameResponse.Data = data[3]
}
