package game
import(
	"net"

)
type Player struct{
	Address  *net.UDPAddr
	Id string
}
func MakePlayer(address *net.UDPAddr, id string)*Player{
	return &Player{
		Address:address,
		Id:id,
	}

}