package states
import (
	"github.com/veandco/go-sdl2/sdl"
	"game/game"
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
)

type GameState struct{
	game *game.Game
	stateManager *StateManager 

}

func MakeGameState()*GameState{
	game := &game.Game{}

	return &GameState{
		game:game,

	}
}
func (gameState *GameState)SetStateManager(stateManager *StateManager){
	gameState.stateManager=stateManager
}
func (gameState *GameState)Init(renderer *sdl.Renderer){
	tiles:=readMap()
	blockSize:= int32(100)
	gameState.game = game.Init(16 * blockSize,10* blockSize,blockSize,tiles,renderer,gameState.stateManager.UpdateState)

}
func (gameState *GameState)Render(){
	gameState.game.Render()
	
}
func (gameState *GameState)Tick(event sdl.Event){
	gameState.game.Tick(event)
	
}
func readMap()[][]int32{
	file, err := os.Open("map.txt")
  
    if err != nil {
        log.Fatalf("failed to open")
  
    }
	var tiles [][]int32
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
 
	defer  file.Close()
	for _,line := range(lines){
		line = strings.TrimSpace(line)
		var elems []int32 
		row:= strings.Split(line," ")
		for _,elem := range(row){
			intElem,_:=strconv.Atoi(elem)
			elems = append(elems,int32(intElem))
		}
		tiles = append(tiles,elems)
	}
	return tiles

    

}