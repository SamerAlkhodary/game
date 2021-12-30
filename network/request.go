package network
import(
	"fmt"
)

type Request interface{
	String()string
	
}
type GetGameRequest struct{
	PlayerId string
	Action string
}
func (getGameRequest *GetGameRequest) String()string{
	return fmt.Sprintf("%s;%s",getGameRequest.Action,getGameRequest.PlayerId)
}

type InGameRequest struct{
	PlayerId string
	GameId string
	Action string
	Data string

}
func (inGameRequest *InGameRequest) String()string{
	return fmt.Sprintf("%s;%s;%s;%s",inGameRequest.Action,inGameRequest.PlayerId,inGameRequest.GameId,inGameRequest.Data)
}

type CreateGameRequest struct{
	PlayerId string
	Action string
	Data string

}
func (createGameRequest *CreateGameRequest) String()string{
	return fmt.Sprintf("%s;%s;%s",createGameRequest.Action,createGameRequest.PlayerId,createGameRequest.Data)
}

type JoinGameRequest struct{
	PlayerId string
	Action string
	GameId string
	Data string

}
func (joinGameRequest *JoinGameRequest) String()string{
	return fmt.Sprintf("%s;%s;%s;%s",joinGameRequest.Action,joinGameRequest.PlayerId,joinGameRequest.GameId,joinGameRequest.Data)
}

