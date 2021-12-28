package model
import(
	"github.com/veandco/go-sdl2/sdl"

)
type Entity interface{
	Render(renderer *sdl.Renderer)
	Tick(int,int)
	IsAlive()bool
	Free()
	HandleCollision(other Entity)
	IsRigid()bool
	GetRect() *sdl.Rect

}
