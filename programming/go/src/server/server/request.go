package server
import (
	"net"
	"strings"
	"strconv"
)
type Request struct{
	Address *net.UDPAddr
	PlayerId int
	GameId int
	Action string
	Data string
}
func MakeRequest(address *net.UDPAddr,data string)*Request{
	info := strings.Split(data,";")
	playerId,_ := strconv.Atoi(info[0])
	gameId,_ := strconv.Atoi(info[1])

	return &Request{
		Address : address,
		PlayerId : playerId,
		GameId : gameId,
		Action : info[2],
		Data : info[3],
	}
}