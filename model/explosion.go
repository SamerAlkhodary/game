package model
import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/img"

)
type  Explosion struct{
	texture *sdl.Texture
	rect *sdl.Rect
	collisionRect *sdl.Rect
	isAlive bool
	isRigid bool
	blockSize int32
	currentFrame int32

}
func MakeExplosion(name string,x , y,blockSize int32,renderer *sdl.Renderer) *Explosion{
	path := "images/explosion/"
	img.Init(img.INIT_PNG)
	surface,_ :=img.Load(path +name+".png")
	//surface := spriteLoader(path+name+".bmp")
	texture:= textureMaker(surface,renderer)
	surface.Free()

	return &Explosion{
		texture:texture,
		rect: &sdl.Rect{X:x,Y:y,W:50,H:50},
		isAlive:true,
		isRigid:true,
		currentFrame:0,
		collisionRect: &sdl.Rect{X:x+10*100/blockSize, Y:y+ 10*100/blockSize,W:blockSize -10*100/blockSize,H:blockSize-10*100/blockSize},
	}

}
func (explosion *Explosion) IsRigid()bool{
	return explosion.isRigid
}
func (explosion *Explosion) IsAlive()bool{
	return explosion.isAlive
}

func (explosion *Explosion) GetRect() *sdl.Rect{
	return explosion.rect
}
func (explosion *Explosion) HandleCollision(other Entity) {


}
func (explosion *Explosion)Render(renderer *sdl.Renderer){
	renderer.CopyEx(explosion.texture, &sdl.Rect{X:0+ 220* explosion.currentFrame,Y:0,W:200,H:200}, explosion.rect, 0 , nil,sdl.FLIP_NONE);
	
	
	

}
func (explosion *Explosion)Tick(eventType,key int){
	explosion.currentFrame += 1
	if explosion.currentFrame > 7{
		explosion.isAlive = false
	}
}
func  (explosion *Explosion) Free(){
	explosion.texture.Destroy()
}

func (explosion *Explosion) GetRotationAngle()float64{
	return 0;
}