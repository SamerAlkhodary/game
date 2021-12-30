package states
import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"github.com/veandco/go-sdl2/img"
	"log"


)

type MenuState struct{
	stateManager *StateManager
	textRects []*sdl.Rect
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

func MakeMenuState(width, height, blockSize int32)*MenuState{
	buttonTexts :=[]string{"New Game","Find Game","Quit Game"}
	textRects:= make([]*sdl.Rect,0)
	buttonRects:= make([]*sdl.Rect,0)
	buttonTextures:= make([]*sdl.Texture,0)
	textTextures:= make([]*sdl.Texture,0)

	offSet:=int32(blockSize/12)
	for i,_ := range(buttonTexts){
		rect := &sdl.Rect{X:width/2 - blockSize ,Y:2*blockSize*int32(i+1)+blockSize/10,W:blockSize*2,H:blockSize}
		textRects= append(textRects,rect)
		newRect := &sdl.Rect{X:width/2 - 1*blockSize*3/2 -offSet,Y:2*blockSize*int32(i+1)+offSet,W:blockSize*3+ 2*offSet,H:blockSize+offSet}
		buttonRects = append(buttonRects,newRect)
	}
	ttf.Init()

	return &MenuState{
		buttonTexts : buttonTexts,
		textRects : textRects,
		buttonRects : buttonRects,
		textTextures : textTextures,
		buttonTextures : buttonTextures, 
		blockSize : blockSize,
		selectedItem :0,
	}
}
func (menuState *MenuState)SetStateManager(stateManager *StateManager){
	menuState.stateManager=stateManager
}
func (menuState *MenuState)Init(renderer *sdl.Renderer){
	fontSize := int(menuState.blockSize)
	for _,text:=range(menuState.buttonTexts){
		sans,_ := ttf.OpenFont("fonts/SansBold.ttf",fontSize );
		surface,_ := sans.RenderUTF8Solid(text,sdl.Color{R:105,G:105,B:105,A:255})
		texture,_ := renderer.CreateTextureFromSurface(surface)
		menuState.textTextures = append(menuState.textTextures,texture)
	}
	path := "images/menu/"
	img.Init(img.INIT_PNG)
	surface,_ :=img.Load(path +"box.png")
	texture:= textureMaker(surface,renderer)
	surface2,_ :=img.Load(path +"box_lit.png")
	texture2:= textureMaker(surface2,renderer)
	menuState.buttonTextures = append(menuState.buttonTextures,texture)
	menuState.buttonTextures = append(menuState.buttonTextures,texture2)
	menuState.renderer = renderer

}
func (menuState *MenuState)Render(){
	for i,texture := range(menuState.textTextures){
		
		if menuState.selectedItem == i{
			menuState.renderer.Copy(menuState.buttonTextures[1] ,&sdl.Rect{X:0,Y:0,W:300,H:100}, menuState.buttonRects[i])
		}else{
			menuState.renderer.Copy(menuState.buttonTextures[0] ,&sdl.Rect{X:0,Y:0,W:300,H:100}, menuState.buttonRects[i])
		}
		menuState.renderer.Copy(texture , nil, menuState.textRects[i])
		menuState.renderer.SetDrawColor(193, 154, 107, 255)
	}
	
}
func (menuState *MenuState)Tick(event sdl.Event){
	eventType,key,isRunning := handleEvent(event)
	if !isRunning{
		menuState.stateManager.UpdateState("Exit")
	}
	if eventType == sdl.KEYDOWN{
		if key ==  sdl.K_UP{
			menuState.selectedItem -=1
			if menuState.selectedItem < 0{
				menuState.selectedItem= len(menuState.buttonTexts)-1
			}

		}else if key == sdl.K_DOWN{
			menuState.selectedItem +=1
			if menuState.selectedItem >len(menuState.buttonTexts)-1{
				menuState.selectedItem=0
			}
		}else if  key == sdl.K_RETURN{
			switch menuState.selectedItem{
			case 0:
				menuState.stateManager.UpdateState("GameState")
			
			break
			case 1:
				menuState.stateManager.UpdateState("GameFinder")
					
			break
			case 2:
				menuState.stateManager.UpdateState("Exit")

			break
			}
		}
	}
	
}
func handleEvent(event sdl.Event)(int,int,bool){
	eventType :=0
	key := 0	 
	switch event.(type) {		  
		case *sdl.QuitEvent:
			println("Quit")
			return 0,0,false
		case  *sdl.KeyboardEvent:
			keyEvent,_ := event.(*sdl.KeyboardEvent)
			eventType = int(event.GetType())
			key = int(keyEvent.Keysym.Sym)
		break;
	}
	
	return eventType,key,true
	
}
func textureMaker(surface *sdl.Surface,renderer *sdl.Renderer)*sdl.Texture{
	err:= surface.SetColorKey(true, sdl.MapRGB(surface.Format, 0x00,0xFF,0xFF)) 
	if  err!=nil{
		log.Printf( "Unable to set Color Key!")
	}
	texture,err:= renderer.CreateTextureFromSurface(surface)
	if err != nil {
		log.Fatal("[SpriteSet] NewSpriteSet(): unable to load resource",err)
	}
	return texture
}
func (menuState *MenuState)Show(){
	
}