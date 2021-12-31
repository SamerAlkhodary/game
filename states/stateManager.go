package states

import(
	"github.com/veandco/go-sdl2/sdl"
	"game/network"	
)
type StateManager struct{
	currentState State
	states map[string] State
	isRunning bool
	playerId string
	isWaiting bool
	gameId string

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
		isWaiting :false,
		playerId : "-1",
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
func (stateManager *StateManager)PlayerId()string{
	return stateManager.playerId
}

func (stateManager *StateManager)IsWaiting() bool{
	return stateManager.isWaiting
}
func (stateManager *StateManager)SetWaiting(value bool){
	stateManager.isWaiting = value
}
func (stateManager *StateManager)SetPlayerId(value string){
	stateManager.playerId = value
}

func (stateManager *StateManager)SetGameId(value string){
	stateManager.gameId = value
}
func (stateManager *StateManager)GameId()string{
	return stateManager.gameId
}

