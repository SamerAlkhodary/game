package server
import(
	"fmt"
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
	NumberOfData string
	Data []*Data
}
func (InGameRequest *InGameRequest)FromString(info string){
	//reposne format : InGame;2;0&10&20&15|1&10&20&15

	data := strings.Split(info,";")
	InGameRequest.PlayerId = data[1]
	InGameRequest.GameId = data[2]

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
			data.TorretX = info[4]
			data.TorretY = info[5]
			data.DidFire = info[6]
		if data.DidFire == "1"{
			data.BulletName = info[7]
		}	
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
	GameId string
	Name string

}
func (joinGameRequest *JoinGameRequest)FromString(info string){
	data := strings.Split(info,";")
	joinGameRequest.PlayerId = data[1]
	joinGameRequest.GameId = data[2]
	joinGameRequest.Name = data[3]
}

type CloseConnectionRequest struct{
	PlayerId string
	GameId string
}
func (closeConnectionRequest *CloseConnectionRequest) FromString(info string){
	data := strings.Split(info,";")
	closeConnectionRequest.PlayerId = data[1]
	closeConnectionRequest.GameId = data[2]
}

