package states
import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"github.com/veandco/go-sdl2/img"
	"game/network"
	"log"
)

type GameFinderState struct{
	gameItems []*network.GameStat
	client *network.Client
	titleTexts []string
	titleRects []*sdl.Rect
	gameStatRects []*sdl.Rect
	gameStatTextures []*sdl.Texture
	titleTextures []*sdl.Texture
	stateManager *StateManager
	textRects []*sdl.Rect
	listRect *sdl.Rect
	listTexture *sdl.Texture
	buttonRects []*sdl.Rect
	buttonTextures []*sdl.Texture
	buttonTexts []string
	textTextures []*sdl.Texture
	blockSize int32
	width int32
	height int32
	selectedItem int
	renderer *sdl.Renderer
}

func MakeGameFinderState(client *network.Client,width, height, blockSize int32)*GameFinderState{
	buttonTexts :=[]string{"Create Game","Join Game","Back"}
	titleTexts := []string{"Game ID","Name","Players"}
	gameStatRects := make([]*sdl.Rect,0)
	gameStatTextures := make([]*sdl.Texture,0)
	textRects:= make([]*sdl.Rect,0)
	buttonRects:= make([]*sdl.Rect,0)
	titleRects:= make([]*sdl.Rect,0)
	buttonTextures:= make([]*sdl.Texture,0)
	titleTextures:= make([]*sdl.Texture,0)
	textTextures:= make([]*sdl.Texture,0)
	listRect := &sdl.Rect{X:blockSize,Y:blockSize,W:width/2, H:height - 2*blockSize}
	items := make([]*network.GameStat,0)
	
	
	
	
	



	ttf.Init()
	return &GameFinderState{
		buttonTexts : buttonTexts,
		gameStatRects: gameStatRects,
		gameStatTextures : gameStatTextures,
		titleTextures : titleTextures,
		titleTexts : titleTexts,
		gameItems: items,
		client:client,
		textRects : textRects,
		buttonRects : buttonRects,
		width:width,
		height:height,
		listRect : listRect,
		textTextures : textTextures,
		buttonTextures : buttonTextures, 
		blockSize : blockSize,
		selectedItem :0,
		titleRects : titleRects,
	}
}
func (gameFinderState *GameFinderState)SetStateManager(stateManager *StateManager){
	gameFinderState.stateManager=stateManager
}
func (gameFinderState *GameFinderState)Init(renderer *sdl.Renderer){
	fontSize := int(gameFinderState.blockSize)/3
	blockSize := gameFinderState.blockSize
	offSet:=int32(blockSize/12)
	textRects:= make([]*sdl.Rect,0)
	buttonRects:= make([]*sdl.Rect,0)	
	width := gameFinderState.width
	go gameFinderState.client.Listen("-1")
	resp := gameFinderState.client.GetResponse()
	response := resp.(*network.GetGameResponse)
	gameFinderState.stateManager.SetPlayerId(response.PlayerId)

	for i,text:=range(gameFinderState.buttonTexts){
		sans,_ := ttf.OpenFont("fonts/SansBold.ttf",fontSize );
		surface,_ := sans.RenderUTF8Solid(text,sdl.Color{R:105,G:105,B:105,A:255})
		texture,_ := renderer.CreateTextureFromSurface(surface)
		_,_,w,h,_ := texture.Query()
		rect := &sdl.Rect{X:width - 4*blockSize +blockSize/8 ,Y:2*blockSize*int32(i+1)+blockSize/9,W:w,H:h}
		textRects= append(textRects,rect)
		newRect := &sdl.Rect{X:width - 4*blockSize-offSet, Y:2*blockSize*int32(i+1)-offSet,W:blockSize*3+ 2*offSet,H:blockSize+offSet}
		buttonRects = append(buttonRects,newRect)
		surface.Free()
		gameFinderState.textTextures = append(gameFinderState.textTextures,texture)
	}
	gameFinderState.buttonRects = buttonRects
	gameFinderState.textRects = textRects
	titleRects:= make([]*sdl.Rect,0)
	for i,text:=range(gameFinderState.titleTexts){
		sans,_ := ttf.OpenFont("fonts/SansBold.ttf",fontSize);
		surface,_ := sans.RenderUTF8Solid(text,sdl.Color{R:0,G:0,B:0,A:255})
		texture,_ := renderer.CreateTextureFromSurface(surface)
		_,_,w,h,_ := texture.Query()
		rect := &sdl.Rect{X:blockSize + 3*blockSize * int32(i) +blockSize/5, Y: blockSize/2,W: w , H:h}
		titleRects= append(titleRects,rect)

		surface.Free()
		gameFinderState.titleTextures = append(gameFinderState.titleTextures,texture)
	}
	gameFinderState.titleRects = titleRects

	

	path := "images/menu/"
	img.Init(img.INIT_PNG)
	surface,_ :=img.Load(path +"box.png")
	texture:= textureMaker(surface,renderer)
	surface2,_ :=img.Load(path +"box_lit.png")
	texture2:= textureMaker(surface2,renderer)
	gameFinderState.buttonTextures = append(gameFinderState.buttonTextures,texture)
	gameFinderState.buttonTextures = append(gameFinderState.buttonTextures,texture2)
	gameFinderState.renderer = renderer
	surface.Free()
	surface2.Free()

}
func (gameFinderState *GameFinderState)Render(){
	gameFinderState.renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	for i,texture := range(gameFinderState.textTextures){
	
		if gameFinderState.selectedItem == i{
			gameFinderState.renderer.Copy(gameFinderState.buttonTextures[1] ,&sdl.Rect{X:0,Y:0,W:300,H:100}, gameFinderState.buttonRects[i])
		}else{
			gameFinderState.renderer.Copy(gameFinderState.buttonTextures[0] ,&sdl.Rect{X:0,Y:0,W:300,H:100}, gameFinderState.buttonRects[i])
		}
		gameFinderState.renderer.Copy(texture , nil, gameFinderState.textRects[i])	
	}
	gameFinderState.renderer.SetDrawColor(255, 255, 255, 100)
	gameFinderState.renderer.FillRect(gameFinderState.listRect)
	gameFinderState.renderer.SetDrawColor(0, 0, 0, 255)
	gameFinderState.renderer.DrawRect(gameFinderState.listRect)
	for i,texture := range(gameFinderState.titleTextures){
		gameFinderState.renderer.Copy(texture , nil, gameFinderState.titleRects[i])
	}
	for i,texture := range(gameFinderState.gameStatTextures){
		gameFinderState.renderer.Copy(texture , nil, gameFinderState.gameStatRects[i])
	}
	gameFinderState.renderer.SetDrawColor(193, 154, 107, 255)	
}
func (gameFinderState *GameFinderState)Tick(event sdl.Event){
	eventType,key,isRunning := handleEvent(event)
	if !isRunning{
		gameFinderState.stateManager.UpdateState("Exit")
	}
	if eventType == sdl.KEYDOWN{
		if key ==  sdl.K_UP{
			gameFinderState.selectedItem -=1
			if gameFinderState.selectedItem < 0{
				gameFinderState.selectedItem= len(gameFinderState.buttonTexts)-1
			}
		}else if key == sdl.K_DOWN{
			gameFinderState.selectedItem +=1
			if gameFinderState.selectedItem >len(gameFinderState.buttonTexts)-1{
				gameFinderState.selectedItem=0
			}
		}else if  key == sdl.K_RETURN{
			switch gameFinderState.selectedItem{
				case 0:
					gameFinderState.createGame()
					gameFinderState.stateManager.SetWaiting(true)
					gameFinderState.stateManager.UpdateState("GameState")	
				break
				case 2:
					gameFinderState.stateManager.UpdateState("MenuState")
				break
			}
		}
	}
	
}
func (gameFinderState *GameFinderState)addGameItem(item *network.GameStat,fontSize int){

	textures,rects:= createListRow(item,fontSize,gameFinderState.renderer,len(gameFinderState.gameItems),gameFinderState.blockSize)
	gameFinderState.gameItems = append(gameFinderState.gameItems,item)
	gameFinderState.gameStatRects = append(gameFinderState.gameStatRects, rects...)
	gameFinderState.gameStatTextures = append(gameFinderState.gameStatTextures, textures...)
}
func createListRow(item *network.GameStat,fontSize int,renderer *sdl.Renderer,index int,blockSize int32)([]*sdl.Texture,[]*sdl.Rect){
	resTextures := []*sdl.Texture{}
	resRects := []*sdl.Rect{} 

	sans,_ := ttf.OpenFont("fonts/Sans.ttf",fontSize);
	surface1,_ := sans.RenderUTF8Solid(item.GameId,sdl.Color{R:0,G:0,B:0,A:255})
	surface2,_ := sans.RenderUTF8Solid(item.GameName,sdl.Color{R:0,G:0,B:0,A:255})
	surface3,_ := sans.RenderUTF8Solid(item.NbrPlayers,sdl.Color{R:0,G:0,B:0,A:255})
	texture1,_ := renderer.CreateTextureFromSurface(surface1)

	_,_,w1,h1,_ := texture1.Query()
	texture2,_ := renderer.CreateTextureFromSurface(surface2)
	_,_,w2,h2,_ := texture2.Query()
	texture3,_ := renderer.CreateTextureFromSurface(surface3)
	_,_,w3,h3,_ := texture3.Query()
	rect1 := &sdl.Rect{X:blockSize +blockSize/5, 
		Y: blockSize+2*blockSize/3* int32(index),W: w1 , H:h1}
	rect2 := &sdl.Rect{X:blockSize + 3*blockSize  +blockSize/5, 
		Y: blockSize +2*blockSize/3* int32(index),W: w2 , H:h2}
	rect3 := &sdl.Rect{X:blockSize + 6*blockSize  +blockSize/5, 
		Y: blockSize+2*blockSize/3* int32(index),W: w3 , H:h3}
	surface1.Free()
	surface2.Free()
	surface3.Free()
	resTextures = append(resTextures,texture1)
	resTextures = append(resTextures,texture2)
	resTextures = append(resTextures,texture3)
	resRects=append(resRects,rect1)
	resRects=append(resRects,rect2)
	resRects=append(resRects,rect3)
	return resTextures,resRects
}
func  (gameFinderState *GameFinderState)GetGames(){
	gameFinderState.client.Send(
		&network.GetGameRequest{
			PlayerId: gameFinderState.stateManager.PlayerId(),
	})
	resp := gameFinderState.client.GetResponse() 
	response := resp.(*network.GetGameResponse)
	log.Println("games",response.Games)
	if len(response.Games)>0{
		log.Println("adding game",response.Games[0])
		for _,game := range(response.Games){
			gameFinderState.addGameItem(game,int(gameFinderState.blockSize/4)) 

		}

	}

}
func  (gameFinderState *GameFinderState)createGame(){
	gameFinderState.client.Send(
		&network.CreateGameRequest{
			PlayerId:gameFinderState.stateManager.PlayerId(),
			Data:"",
		})
	resp := gameFinderState.client.GetResponse() 
	response := resp.(*network.CreateGameResponse)
	gameFinderState.addGameItem(response.Game,int(gameFinderState.blockSize/4))
}
func (gameFinderState *GameFinderState)Show(){
	for _,item := range(gameFinderState.gameStatTextures){
		item.Destroy()

	}
	gameFinderState.gameStatTextures=make([]*sdl.Texture,0)
	gameFinderState.gameItems=make([]*network.GameStat,0)
	gameFinderState.gameStatRects=make([]*sdl.Rect,0)


	gameFinderState.GetGames()
}