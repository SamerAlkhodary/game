package model
import "fmt"

type Pos struct{
	X int32
	Y int32
}

func MakePos(x int32 , y int32) *Pos{
	return &Pos{
		X:x,
		Y:y,
	}
}
func (p *Pos) String() string{
	return fmt.Sprintf("(x:%d, y:%d)",p.X,p.Y)
}

func (pos *Pos) Move(dx int32 ,dy int32) *Pos{
	
	return &Pos{
		pos.X+dx,
		pos.Y+dy,
	}
}

