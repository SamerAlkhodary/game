package model
import(
	"github.com/veandco/go-sdl2/sdl"

)
type Entity interface{
	Render(renderer *sdl.Renderer,camera *sdl.Rect)
	Tick(int,int)
	IsAlive()bool
	Free()

}
