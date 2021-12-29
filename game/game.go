package game
import (
		"game/model"
		"github.com/veandco/go-sdl2/sdl"
	)
type Game struct{
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
	player model.Entity
}

func Init(width,height int32,blockSize int32,tiles [][]int32,renderer *sdl.Renderer,stateChanger func (string) ) *Game{
	

	game:= &Game{
		entities:make([]model.Entity, 0),
		bullets: make([]model.Entity, 0),
		explosions :make([]model.Entity, 0),
		players:make([]model.Entity, 0),
		renderer:renderer,
		width: width,
		height:height,
		absolutePos: model.MakePos(0,0),
		frames : 30,
		blockSize:blockSize,
		mapTiles:tiles,
		stateChanger: stateChanger,
		running:true,
	} 
	game.initEntities()
	return game


}
func (game *Game)initEntities(){
	blockSize := game.blockSize
	background := model.MakeBackground("game.mapTiles",game.width,game.height,game.renderer)
	player1Rect := &sdl.Rect{ X:1*blockSize,Y:2*blockSize,W:1*blockSize,H:1*blockSize}
	player1KeyControl := model.MakeKeyController('w','e',sdl.K_SPACE,sdl.K_RSHIFT)
	player1 := model.MakePlayer("Samer",1,player1Rect,game.renderer,game.blockSize,player1KeyControl,game.AddBullet)
	game.player = player1
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
	for _,player := range(game.players){
		player.Tick(eventType,key)
		for _,entity := range(game.entities){
			player.HandleCollision(entity)
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
func handleEvent(event sdl.Event)(int,int,bool){
		eventType :=0
		key := 0	 
		switch event.(type) {
			case  *sdl.KeyboardEvent:
				keyEvent,_ := event.(*sdl.KeyboardEvent)
				eventType = int(event.GetType())
				key = int(keyEvent.Keysym.Sym)
			break;
			  
			case *sdl.QuitEvent:
					println("Quit")
					return 0,0,false
		}
		
		return eventType,key,true
		
}