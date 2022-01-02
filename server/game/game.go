package game
import(
	"fmt"
)
type Game struct{
	Players [] *Player	
	Id string
	Name string
}
func MakeGame(id string,name string,player *Player)*Game{
	players:=make([]*Player,0)
	players = append(players,player)
	return &Game{
		Id : id,
		Name: name,
		Players : players,
	}
}
func (game *Game)EnoughPlayers() bool{
	return len(game.Players)>1
}
func (game *Game)String()string{
	return fmt.Sprintf("%d,%d",game.Id,len(game.Players))
}

