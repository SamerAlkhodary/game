package states

import(
	"github.com/veandco/go-sdl2/sdl"
	"game/network"
	
	
)


type StateManager struct{
	currentState State
	states map[string] State
	isRunning bool

}
func MakeStateManager(client *network.Client, width, height, blockSize int32)*StateManager{
	states:= make(map[string]State)
	states["MenuState"]= MakeMenuState(width, height, blockSize)
	states["GameState"] = MakeGameState(client,width, height, blockSize)
	states["GameFinder"] = MakeGameFinderState(client,width, height, blockSize)
	
	return &StateManager{
		states:states,
		currentState: states["MenuState"], 
		isRunning : true,
	}
}
func (stateManager *StateManager)Init(renderer *sdl.Renderer){
	stateManager.states["MenuState"].SetStateManager(stateManager)
	stateManager.states["GameState"].SetStateManager(stateManager)
	stateManager.states["GameFinder"].SetStateManager(stateManager)
	stateManager.states["MenuState"].Init(renderer)
	stateManager.states["GameState"].Init(renderer)
	stateManager.states["GameFinder"].Init(renderer)
}
func (stateManager *StateManager) UpdateState(stateName string){
	if stateName == "Exit"{
		stateManager.isRunning = false
		return
	}
	stateManager.currentState = stateManager.states[stateName]
	stateManager.currentState.Show()
}
func (stateManager *StateManager) GetCurrentState()State{
	return stateManager.currentState
}
func (stateManager *StateManager)IsRunning() bool{
	return stateManager.isRunning
}
