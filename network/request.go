package network
import(
	"fmt"
)

type Request interface{
	String()string
	
}
type GetGameRequest struct{
	PlayerId string
}
func (getGameRequest *GetGameRequest) String()string{
	return fmt.Sprintf("%s;%s","GetGame",getGameRequest.PlayerId)
}

type InGameRequest struct{
	PlayerId string
	GameId string
	NumberOfData string
	Action string
	Data *Data
}
func (inGameRequest *InGameRequest) String()string{
	fmt.Println(inGameRequest.Data.String())
	return fmt.Sprintf("%s;%s;%s;%s;%s","InGame",inGameRequest.PlayerId,inGameRequest.GameId,"1",inGameRequest.Data.String())
}

type CreateGameRequest struct{
	PlayerId string
	Data string

}
func (createGameRequest *CreateGameRequest) String()string{
	return fmt.Sprintf("%s;%s;%s","CreateGame",createGameRequest.PlayerId,createGameRequest.Data)
}

type JoinGameRequest struct{
	PlayerId string
	GameId string
	Name string

}
func (joinGameRequest *JoinGameRequest) String()string{
	return fmt.Sprintf("%s;%s;%s;%s","JoinGame",joinGameRequest.PlayerId,joinGameRequest.GameId,joinGameRequest.Name)
}

