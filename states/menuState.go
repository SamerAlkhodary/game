package states
import (
	"github.com/veandco/go-sdl2/sdl"
)

type MenuState struct{
	stateManager *StateManager
}

func MakeMenuState()*MenuState{
	return &MenuState{
		
	}
}
func (menuState *MenuState)SetStateManager(stateManager *StateManager){
	menuState.stateManager=stateManager
}
func (menuState *MenuState)Init(renderer *sdl.Renderer){

}
func (menuState *MenuState)Render(){
	
}
func (menuState *MenuState)Tick(event sdl.Event){
	
}