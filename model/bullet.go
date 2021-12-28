package model
import(
	
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"fmt"
	"log"
)

type Bullet struct{
	speed float64
	name string
	damage int32
	initialRect *sdl.Rect
	effectRadius int32
	rect *sdl.Rect
	texture *sdl.Texture
	blockSize int32
	isAlive bool
	bulletRange int32
	rotationAngle float64
	xSpeed float64
	ySpeed float64
	isRigid bool
	collisionRect *sdl.Rect
}
func MakeBullet(name string,rect *sdl.Rect,renderer *sdl.Renderer,blockSize int32,damage int32, speed float64,effectRadius int32,bulletRange int32,rotationAngle float64)*Bullet{
	bulletPath := "images/bullets/"
	bulletSurface := spriteLoader(bulletPath+name+".bmp")
	bulletTexture:= textureMaker(bulletSurface,renderer)
	xNew, yNew := calculateCoords(float64(blockSize),rotationAngle,float64(rect.X),float64(rect.Y))
	rect.X= xNew
	rect.Y= yNew

	return &Bullet{
		speed : speed,
		xSpeed : 0,
		ySpeed : 0,
		name : name,
		damage : damage,
		effectRadius : effectRadius,
		texture : bulletTexture,
		rect : rect,
		blockSize : blockSize,
		isAlive:true,
		bulletRange : bulletRange,
		rotationAngle : rotationAngle,
		initialRect: &sdl.Rect{X: rect.X, Y: rect.Y, W:rect.W,H:rect.H},
		collisionRect: &sdl.Rect{X:rect.X+blockSize*10/100,Y:rect.Y+blockSize*10/100,W:rect.W-blockSize*20/100,H:rect.H-blockSize*20/100},
		isRigid: true,
	}

}
func (bullet *Bullet)Render(renderer *sdl.Renderer){
	renderer.CopyEx(bullet.texture, &sdl.Rect{X:0,Y:0,W:50,H:50}, bullet.rect, bullet.rotationAngle , nil,sdl.FLIP_NONE);
	
	
	

}
func outOfRange(initialRect, currentRect *sdl.Rect, bulletRange int32) bool{
	distance := math.Sqrt(math.Pow(float64(initialRect.X) -float64(currentRect.X),2) + math.Pow(float64(initialRect.Y) -float64(currentRect.Y),2))
	return distance> float64(bulletRange)

}
func (bullet *Bullet)Tick(eventType,key int){
	bullet.xSpeed,bullet.ySpeed =calculateSpeed(bullet.speed,bullet.rotationAngle)
	bullet.rect.X += int32(bullet.xSpeed)
	bullet.rect.Y += int32(bullet.ySpeed)
	bullet.collisionRect.X += int32(bullet.xSpeed)
	bullet.collisionRect.Y += int32(bullet.ySpeed)
	bullet.isAlive = !outOfRange(bullet.initialRect,bullet.rect,bullet.bulletRange)
	
}

func (bullet *Bullet)IsAlive()bool{
	return bullet.isAlive

}
func calculateCoords(length float64 ,rotationAngle float64,x1,y1 float64)(int32,int32){
	radAngle1 :=  float64(90- (int(rotationAngle)%360))* math.Pi/180
	
	x := x1+length/2 - 20 + math.Cos(radAngle1) *length/2
	y := y1+length/2 - 20  - math.Sin(radAngle1) *length/2
	fmt.Println(x,y ,(int(rotationAngle)%360),x1,y1)

	
	return int32(x),int32(y)
}

func (bullet *Bullet)Free(){
	bullet.texture.Destroy()
}
func (bullet *Bullet)HandleCollision(other Entity){
	
	if other.IsRigid(){
		tile,done := other.(*Tile)
	if done{
		collision:=tile.collisionRect.HasIntersection(bullet.collisionRect)
		if collision{
			log.Println("coll")
			bullet.isAlive = false
		}
	}
		
	}

}
func(bullet *Bullet)IsRigid()bool{
	return bullet.isRigid
}
func (bullet *Bullet)GetRect()*sdl.Rect{
	return bullet.rect

}