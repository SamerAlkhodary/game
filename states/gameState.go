package states
import (
	"github.com/veandco/go-sdl2/sdl"
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
	"game/network"
	"game/model"
)

type GameState struct{
	game *Game
	stateManager *StateManager 
	blockSize int32
	width int32
	height int32
	client *network.Client
}

func MakeGameState(client *network.Client,width,height,blockSize int32)*GameState{
	game := &Game{}

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
	gameState.game = InitGame(gameState.stateManager.GameId(),gameState.stateManager.PlayerId(),
	16 * gameState.blockSize,10* gameState.blockSize,
	gameState.blockSize,tiles,renderer,gameState.stateManager)

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
type Game struct{
	gameId string
	playerId string
	stateManager *StateManager
	entities []model.Entity
	bullets []model.Entity
	players []model.Entity
	explosions []model.Entity
	window *sdl.Window
	renderer *sdl.Renderer
	width int32
	height int32
	stateChanger func (string)
	absolutePos *model.Pos
	xSpeed,ySpeed int32
	frames uint32
	mapTiles [][]int32
	blockSize int32
	running bool
	mainPlayer model.Entity
}

func InitGame(gameId, playerId string,width,height int32,blockSize int32,tiles [][]int32,renderer *sdl.Renderer,stateManager *StateManager) *Game{

	game:= &Game{
		entities:make([]model.Entity, 0),
		bullets: make([]model.Entity, 0),
		explosions :make([]model.Entity, 0),
		players:make([]model.Entity, 0),
		renderer:renderer,
		stateManager:stateManager,
		width: width,
		height:height,
		playerId:playerId,
		gameId:gameId,
		absolutePos: model.MakePos(0,0),
		frames : 30,
		blockSize:blockSize,
		mapTiles:tiles,
		running:true,
	} 
	game.initEntities()
	return game


}
func (game *Game)initEntities(){
	blockSize := game.blockSize
	background := model.MakeBackground("game.mapTiles",game.width,game.height,game.renderer)
	player1Rect := &sdl.Rect{ X:2*blockSize,Y:2*blockSize,W:1*blockSize,H:1*blockSize}
	player1KeyControl := model.MakeKeyController('w','e',sdl.K_SPACE,sdl.K_RSHIFT)
	player1 := model.MakePlayer("Player"+game.playerId,game.playerId,player1Rect,game.renderer,game.blockSize,player1KeyControl,game.AddBullet)
	game.mainPlayer = player1
	game.players = append(game.players,player1)

	game.AddEntity(background)	
	for i ,_:= range(game.mapTiles){
		for j,_ := range(game.mapTiles[i]){
			game.makeTile(int32(i),int32(j))
		}
	}
}
func (game *Game)makeTile(i, j int32){
	value := game.mapTiles[i][j]
	tile:=model. MakeTile(j,i,value,game.renderer,game.blockSize)
	if tile!=nil{
		game.AddEntity(tile)
	}
}

func (game *Game)AddBullet(e model.Entity){

	game.bullets = append(game.bullets, e)
	
}
func (game *Game)AddExplosion(e model.Entity){

	game.explosions = append(game.explosions, e)
	
}
func (game *Game)modifyTile(i, j int32 , tileType int32){
	if game.mapTiles[i][j]==0{
		tile:=model. MakeTile(j,i,tileType,game.renderer,game.blockSize)
		game.AddEntity(tile)
	}

}
func (game *Game)AddEntity(e model.Entity){
	game.entities = append(game.entities, e)
	
}

func  (game *Game) Render(){
	
	for _,entity := range(game.entities){
			entity.Render(game.renderer)
		
	}
	for _,bullet := range(game.bullets){
		bullet.Render(game.renderer)
	
	}
	for _,explosion := range(game.explosions){
		explosion.Render(game.renderer)
		
	}
	for _,player := range(game.players){
		player.Render(game.renderer)
		
	}
}
func  (game *Game) Tick(event sdl.Event){
	eventType,key,running:=handleEvent(event)
	if running ==false{
		game.stateChanger("Exit")
	}
	
	for _,entity := range(game.entities){
			entity.Tick(eventType,key)
	}
	for _,explosion := range(game.explosions){
		explosion.Tick(eventType,key)
		
	}
	for _,bullet := range(game.bullets){
		bullet.Tick(eventType,key)
		for _,entity := range(game.entities){
			bullet.HandleCollision(entity)

		}
	}
	if !game.stateManager.IsWaiting(){
		for _,player := range(game.players){
			player.Tick(eventType,key)
			for _,entity := range(game.entities){
				player.HandleCollision(entity)
			}
			
		}
		
	}
	
	var deadBullets []model.Entity
	game.entities,_ = filterAlive(game.entities)
	game.bullets,deadBullets = filterAlive(game.bullets)
	for _,bulletEntity := range(deadBullets){
		coords:= bulletEntity.GetRect()
		game.AddExplosion(model.MakeExplosion("exp",coords.X,coords.Y,game.blockSize,game.renderer))
		game.modifyTile(coords.Y/game.blockSize,coords.X/game.blockSize,6)

	}
	

}

func filterAlive(entities []model.Entity) ([]model.Entity,[]model.Entity){
	res := []model.Entity{}
	dead := []model.Entity{}
	for _,entity := range(entities){
		if entity.IsAlive(){
			res = append(res,entity)
		}else{
			dead = append(dead,entity)
			entity.Free()
		}
	}
	return res,dead

}