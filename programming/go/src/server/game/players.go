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
