package model
import(
	
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/img"

)

type FogOfWar struct{
	texture *sdl.Texture
	rect *sdl.Rect
	from *sdl.Rect
	isAlive bool
	isRigid bool
	coordI int
	coordJ int
	fogMatrix [][]bool
	collisionRect *sdl.Rect
}
func MakeFogOfWar(i,j int32, renderer *sdl.Renderer,blockSize int32,fogMatrix [][]bool) *FogOfWar{

	path := "images/fog/"
	img.Init(img.INIT_PNG)
	surface,_ :=img.Load(path +"FogofWar.png")
	texture:= textureMaker(surface,renderer)
	surface.Free()
	return &FogOfWar{
		rect : &sdl.Rect{X:i*blockSize, Y:j*blockSize, W:blockSize, H:blockSize},
		isAlive:true,
		isRigid: false,
		coordI: int(i),
		coordJ: int(j),
		collisionRect: &sdl.Rect{X:int32(i*blockSize+blockSize*20/100), Y:j*blockSize+blockSize*20/100, W:blockSize-blockSize*40/100, H:blockSize-blockSize*40/100},
		fogMatrix:fogMatrix,
		texture:texture,

	}
}
func (fogOfWar *FogOfWar) Render(renderer *sdl.Renderer,camera *sdl.Rect){
	renderer.SetDrawColor(0, 0, 0, 255)

		if !isFog(fogOfWar.coordI+1,fogOfWar.coordJ,fogOfWar.fogMatrix){
			renderer.CopyEx(fogOfWar.texture, &sdl.Rect{X:0,Y:35,W:30,H:30}, fogOfWar.rect, 0 , nil,sdl.FLIP_NONE);
		}else if  !isFog(fogOfWar.coordI-1,fogOfWar.coordJ,fogOfWar.fogMatrix){
			renderer.CopyEx(fogOfWar.texture, &sdl.Rect{X:70,Y:35,W:28,H:30}, fogOfWar.rect, 0 , nil,sdl.FLIP_NONE);
		}else if  !isFog(fogOfWar.coordI,fogOfWar.coordJ+1,fogOfWar.fogMatrix){
			renderer.CopyEx(fogOfWar.texture, &sdl.Rect{X:35,Y:0,W:30,H:30}, fogOfWar.rect, 0 , nil,sdl.FLIP_NONE);
		}else if  !isFog(fogOfWar.coordI,fogOfWar.coordJ-1,fogOfWar.fogMatrix){
			renderer.CopyEx(fogOfWar.texture, &sdl.Rect{X:35,Y:70,W:28,H:28}, fogOfWar.rect, 0 , nil,sdl.FLIP_NONE);
		}else if  !isFog(fogOfWar.coordI+1,fogOfWar.coordJ+1,fogOfWar.fogMatrix){
			renderer.CopyEx(fogOfWar.texture, &sdl.Rect{X:0,Y:0,W:28,H:28}, fogOfWar.rect, 0 , nil,sdl.FLIP_NONE);
		}else if  !isFog(fogOfWar.coordI-1,fogOfWar.coordJ-1,fogOfWar.fogMatrix){
			renderer.CopyEx(fogOfWar.texture, &sdl.Rect{X:70,Y:70,W:28,H:28}, fogOfWar.rect, 0 , nil,sdl.FLIP_NONE);
		}else if  !isFog(fogOfWar.coordI+1,fogOfWar.coordJ-1,fogOfWar.fogMatrix){
			renderer.CopyEx(fogOfWar.texture, &sdl.Rect{X:0,Y:70,W:28,H:28}, fogOfWar.rect, 0 , nil,sdl.FLIP_NONE);
		}else if  !isFog(fogOfWar.coordI-1,fogOfWar.coordJ+1,fogOfWar.fogMatrix){
			renderer.CopyEx(fogOfWar.texture, &sdl.Rect{X:70,Y:0,W:28,H:28}, fogOfWar.rect, 0 , nil,sdl.FLIP_NONE);
		}else{
			renderer.FillRect(fogOfWar.rect)

		}
	
	renderer.SetDrawColor(193, 154, 107, 255)	
}
func (fogOfWar *FogOfWar) Tick(eventType,key int){
	
}
func (fogOfWar *FogOfWar)Kill(){
	fogOfWar.isAlive = false
	fogOfWar.fogMatrix[fogOfWar.coordI][fogOfWar.coordJ]=false
}
func (fogOfWar *FogOfWar)IsAlive()bool{
	return fogOfWar.isAlive
}
func (fogOfWar *FogOfWar)Free(){
	//fogOfWar.texture.Destroy()

}
func (fogOfWar *FogOfWar)HandleCollision(other Entity){
	

}
func(fogOfWar *FogOfWar)IsRigid() bool{
	return fogOfWar.isRigid
}
func (fogOfWar *FogOfWar)GetRect()*sdl.Rect{
	return fogOfWar.rect

}
func (fogOfWar *FogOfWar)GetRotationAngle()float64{
	return 0

}
func isFog(i ,j int, fogMatrix [][]bool)bool{
	if i >= len(fogMatrix){
		return true
	}
	if i < 0{
		return true
	}
	if j >= len(fogMatrix[0]){
		return true
	}
	if j <0{
		return true
	}
	return fogMatrix[i][j]
}
