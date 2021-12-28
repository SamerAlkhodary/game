package states
import (
	"github.com/veandco/go-sdl2/sdl"
)

type GameFinderState struct{
	stateManager *StateManager
}

func MakeGameFinderState()*GameFinderState{
	return &GameFinderState{
		

	}
}
func (gameFinderState *GameFinderState)SetStateManager(stateManager *StateManager){
	gameFinderState.stateManager=stateManager
}
func (gameFinderState *GameFinderState)Init(renderer *sdl.Renderer){

}
func (gameFinderState *GameFinderState)Render(){
	
}
func (gameFinderState *GameFinderState)Tick(event sdl.Event){
	
}