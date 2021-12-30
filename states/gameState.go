package states
import (
	"github.com/veandco/go-sdl2/sdl"
	"game/game"
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
	"game/network"
)

type GameState struct{
	game *game.Game
	stateManager *StateManager 
	blockSize int32
	width int32
	height int32
	client *network.Client
}

func MakeGameState(client *network.Client,width,height,blockSize int32)*GameState{
	game := &game.Game{}

	return &GameState{
		game:game,
		width:width,
		height:height,
		blockSize: blockSize,
		client:client,
	}
}
func (gameState *GameState)SetStateManager(stateManager *StateManager){
	gameState.stateManager=stateManager
}
func (gameState *GameState)Init(renderer *sdl.Renderer){
	tiles:=readMap()
	gameState.game = game.Init(16 * gameState.blockSize,10* gameState.blockSize,gameState.blockSize,tiles,renderer,gameState.stateManager.UpdateState)

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
func(gameState *GameState)Show(){

}