package server
import(
	"fmt"
	"log"
)

type Response interface{
	String()string

}
type  GameStat struct{
	GameName string
	GameId string
	NbrPlayers string

}
func(gameStat *GameStat) String()string{
	return fmt.Sprintf("%s&%s&%s&",gameStat.GameId,gameStat.GameName,gameStat.NbrPlayers)
}
type GetGameResponse struct{
	PlayerId string
	Games []*GameStat
	
}
func (getGameResponse *GetGameResponse)String()string{
	//reposne format : GetGames;2;1&game1&2|2&game2&3
	res:= fmt.Sprintf("%s;%d;%s;","GetGame",len(getGameResponse.Games),getGameResponse.PlayerId)
	if len(getGameResponse.Games) >0{
		for i,game:= range(getGameResponse.Games){
			res += game.String()
			if i < len(getGameResponse.Games)-1{
				res +="|"
			}
		}
	}
	log.Println("sending:",res)
	return res
}


type InGameResponse struct{
	GameId string
	Data []*Data
}
func (inGameResponse *InGameResponse)String()string{
	//reposne format : InGame;2;0&10&20&15|1&10&20&15
	res:= fmt.Sprintf("%s;%s;%s;","InGame",inGameResponse.GameId,fmt.Sprintf("%d",len(inGameResponse.Data)))
	if len(inGameResponse.Data) >0{
		for i,game:= range(inGameResponse.Data){
			res += game.String()
			if i < len(inGameResponse.Data)-1{
				res +="|"
			}
		}
	}
	return res
}

type CreateGameResponse struct{
	GameId string
	PlayerId string

}
func (createGameResponse *CreateGameResponse)String()string{
	return fmt.Sprintf("CreateGame;%s;%s",createGameResponse.GameId,createGameResponse.PlayerId)

}

type JoinGameResponse struct{
	OtherPlayerId string
	OtherPlayerNumber string
}
func (joinGameResponse *JoinGameResponse)String()string{
	return fmt.Sprintf("JoinGame;%s;%s;",joinGameResponse.OtherPlayerId,joinGameResponse.OtherPlayerNumber,)

}
type CloseConnectionResponse struct{
	PlayerId string
}
func (closeConnectionResponse *CloseConnectionResponse)String()string{
	return fmt.Sprintf("ExitGame;%s;",closeConnectionResponse.PlayerId)

}