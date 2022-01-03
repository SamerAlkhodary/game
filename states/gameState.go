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
	"fmt"
)

type GameState struct{
	game *Game
	stateManager *StateManager 
	blockSize int32
	width int32
	height int32
	client *network.Client
	renderer *sdl.Renderer
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
	gameState.renderer = renderer
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
	tiles:=readMap()
	gameState.game = InitGame(
		gameState.client,
		16 * gameState.blockSize,10* gameState.blockSize,
		gameState.blockSize,tiles,gameState.renderer,gameState.stateManager)
	gameState.game.initEntities()
}
type Game struct{
	gameId string
	client  *network.Client
	playerId string
	stateManager *StateManager
	entities []model.Entity
	bullets []model.Entity
	players []*model.Player
	explosions []model.Entity
	fogEntities []model.Entity
	fogMatrix [][]bool
	window *sdl.Window
	renderer *sdl.Renderer
	width int32
	height int32
	absolutePos *model.Pos
	xSpeed,ySpeed int32
	frames uint32
	mapTiles [][]int32
	blockSize int32
	running bool
	mainPlayer *model.Player
	hasFired bool
	bulletName string
}

func InitGame(client*network.Client,width,height int32,blockSize int32,tiles [][]int32,renderer *sdl.Renderer,stateManager *StateManager) *Game{
	game:= &Game{
		entities:make([]model.Entity, 0),
		fogEntities:make([]model.Entity, 0),
		bullets: make([]model.Entity, 0),
		explosions :make([]model.Entity, 0),
		fogMatrix : make([][]bool,len(tiles[0])),
		players:make([]*model.Player, 0),
		renderer:renderer,
		stateManager:stateManager,
		width: width,
		height:height,
		playerId: stateManager.PlayerId(),
		gameId: stateManager.GameId(),
		client:client,
		absolutePos: model.MakePos(0,0),
		frames : 30,
		blockSize:blockSize,
		mapTiles:tiles,

		running:true,
	} 
	
	return game

}
func (game *Game)initEntities(){
	blockSize := game.blockSize
	background := model.MakeBackground("game.mapTiles",game.width,game.height,game.renderer)
	player1Rect := &sdl.Rect{ X:2*blockSize,Y:2*blockSize,W:1*blockSize,H:1*blockSize}
	player1KeyControl := model.MakeKeyController('w','e',sdl.K_SPACE,sdl.K_RSHIFT)
	player1 := model.MakePlayer(
		"Player"+game.stateManager.playerNumber,
		game.stateManager.PlayerId(),
		game.stateManager.playerNumber,
		player1Rect,	
		game.renderer,
		game.blockSize,
		player1KeyControl,
		game.AddBullet)
	game.mainPlayer = player1
	game.players = append(game.players,player1)

	if game.stateManager.playerNumber == "2"{
		player2Rect := &sdl.Rect{ X:2*blockSize,Y:2*blockSize,W:1*blockSize,H:1*blockSize}
		player := model.MakePlayer(
			"Player"+game.stateManager.otherPlayerNumber,
			game.stateManager.otherPlayerId,
			game.stateManager.otherPlayerNumber,
			player2Rect	,
			game.renderer,
			game.blockSize,
			nil,
			game.AddBullet)
		game.players = append(game.players,player)

	}
	if game.stateManager.IsWaiting(){
		resp := game.client.GetResponse() 
		response := resp.(*network.JoinGameResponse)
		game.stateManager.otherPlayerNumber = response.Player2Number
		game.stateManager.otherPlayerId = response.Player2Id
		player2Rect := &sdl.Rect{ X:4*blockSize,Y:2*blockSize,W:1*blockSize,H:1*blockSize}		

		player := model.MakePlayer(
			"Player"+game.stateManager.otherPlayerNumber,
			game.stateManager.otherPlayerId,
			game.stateManager.otherPlayerNumber,
			player2Rect	,
			game.renderer,
			game.blockSize,
			nil,
			game.AddBullet)
		game.players = append(game.players,player)
		game.stateManager.SetWaiting(false)
	}
	game.AddEntity(background)	
	for i := 0; i < len(game.mapTiles[0]); i++{
		game.fogMatrix[i] = make([]bool,len(game.mapTiles))
		for j := 0; j < len(game.mapTiles); j++{
			game.fogMatrix[i][j] = true
		}
	}
	for i ,_:= range(game.mapTiles){
		for j,_ := range(game.mapTiles[i]){
			game.makeTile(int32(i),int32(j))
			game.makeFog(int32(i),int32(j),game.fogMatrix)
		}
	}

}
func (game *Game)makeFog(i, j int32 ,fogMatrix[][]bool){
	fogTile:=model.MakeFogOfWar(j,i,game.renderer,game.blockSize,fogMatrix)
	game.fogEntities = append(game.fogEntities,fogTile)
	
	
}
func (game *Game)makeTile(i, j int32){
	value := game.mapTiles[i][j]
	tile:=model. MakeTile(j,i,value,game.renderer,game.blockSize)
	if tile!=nil{
		game.AddEntity(tile)
	}
}

func (game *Game)AddBullet(e model.Entity){
	game.hasFired = true
	game.bulletName = "bullet1"

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
	for _,player := range(game.players){
		player.Render(game.renderer)
	}
	for _,explosion := range(game.explosions){
		explosion.Render(game.renderer)
		
	}
	for _,fog := range(game.fogEntities){
		fog.Render(game.renderer)
		
	}
}
func  (game *Game) Tick(event sdl.Event){
	game.hasFired = false
	eventType,key,running:=handleEvent(event)
	if running ==false{
		game.stateManager.UpdateState("Exit")
	}
	
	for _,entity := range(game.entities){
			entity.Tick(eventType,key)
	}
	for _,explosion := range(game.explosions){
		explosion.Tick(eventType,key)
		
	}
	for _,bullet := range(game.bullets){
		bullet.Tick(eventType,key)
		for _,player := range(game.players){
			bullet.HandleCollision(player)
		}
		for _,entity := range(game.entities){
			bullet.HandleCollision(entity)
		}
	}
	for _,player := range(game.players){
		player.Tick(eventType,key)
		for _,entity := range(game.entities){
			player.HandleCollision(entity)
		}
		for _,otherplayer := range(game.players){
			player.HandleCollision(otherplayer)
		}
		for _,fog := range(game.fogEntities){
			game.mainPlayer.HandleCollision(fog)
		}
				

	}
	didFire :="0"
	if game.hasFired{
		didFire = "1"
	}
	game.client.Send(
		&network.InGameRequest{
			PlayerId: game.stateManager.PlayerId(),
			GameId :  game.stateManager.GameId(),
			NumberOfData :"1",
			Data:&network.Data{
				PlayerId : fmt.Sprintf("%s",game.stateManager.PlayerId()),
				PlayerX :fmt.Sprintf("%d",game.mainPlayer.GetRect().X),
				PlayerY : fmt.Sprintf("%d",game.mainPlayer.GetRect().Y),
				PlayerRotationAngle : fmt.Sprintf("%d",int(game.mainPlayer.GetRotationAngle())),
				TorretX: fmt.Sprintf("%d",game.mainPlayer.TorretRect().X),
				TorretY:fmt.Sprintf("%d",game.mainPlayer.TorretRect().Y),
				DidFire: didFire,
				BulletName : game.bulletName,
			},

	})
	if game.stateManager.isMultiPlayer{
		resp := game.client.GetResponse()
		_,ok := resp.(*network.CloseConnectionResponse)
		if ok{
			game.stateManager.isMultiPlayer = false
			return
		}
		response,_ := resp.(*network.InGameResponse)
		otherPlayer := game.players[1]
		otherPlayer.Update(response.Data[0])
	}

	var deadBullets []model.Entity
	game.entities,_ = filterAlive(game.entities)
	game.bullets,deadBullets = filterAlive(game.bullets)
	game.fogEntities,_ = filterAlive(game.fogEntities)

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