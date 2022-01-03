package model
import(
	
	"github.com/veandco/go-sdl2/sdl"
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


	return &FogOfWar{
		rect : &sdl.Rect{X:i*blockSize, Y:j*blockSize, W:blockSize, H:blockSize},
		isAlive:true,
		isRigid: false,
		coordI: int(i),
		coordJ: int(j),
		collisionRect: &sdl.Rect{X:int32(i*blockSize+blockSize*20/100), Y:j*blockSize+blockSize*20/100, W:blockSize-blockSize*40/100, H:blockSize-blockSize*40/100},
		fogMatrix:fogMatrix,

	}
}
func (fogOfWar *FogOfWar) Render(renderer *sdl.Renderer){
	renderer.SetDrawColor(0, 0, 0, 255)
	if fogOfWar.coordI < len(fogOfWar.fogMatrix)-1{
		if !fogOfWar.fogMatrix[fogOfWar.coordI+1][fogOfWar.coordJ]{
			renderer.SetDrawColor(0, 100, 0, 255)
		}
	}
	renderer.FillRect(fogOfWar.rect)
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
