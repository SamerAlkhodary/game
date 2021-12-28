package states
import (
	"github.com/veandco/go-sdl2/sdl"
)
type State interface{
	Render()
	Tick(sdl.Event)
	Init(*sdl.Renderer)
	SetStateManager(stateManager *StateManager)

}