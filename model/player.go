package model
import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"math"
)
 type Player struct{
	name string
	id int 
	pos *Pos
	rect *sdl.Rect
	blockSize int32
	xSpeed float64
	ySpeed float64
	baseSpeed float64
	rotationSpeed int
	tankTexture *sdl.Texture
	tankRotationAngle float64
	keyController * KeyController
	fire bool
	alreadyFired bool
	torret *Torret
	isAlive bool
	renderer *sdl.Renderer
	addBullet func(Entity)


}
func MakePlayer(name string, id int,rect *sdl.Rect,renderer *sdl.Renderer,blockSize int32,keyController *KeyController, addBullets func(Entity))*Player{
	torret := MakeTorret("torret1",&sdl.Rect{X:rect.X,Y:rect.Y,W:rect.W,H:rect.H},4,500,renderer,1)
	pos := MakePos(rect.X,rect.Y)
	tankPath := "images/tank/"
	tankSurface := spriteLoader(tankPath+"tank.bmp")
	tankTexture:= textureMaker(tankSurface,renderer)


	return &Player{
		name:name,
		tankTexture: tankTexture,
		id:id,
		pos:pos,
		isAlive:true,
		fire : true,
		keyController:keyController,
		rect:rect,
		rotationSpeed:1,
		xSpeed:1,
		ySpeed:1,
		blockSize:blockSize,
		tankRotationAngle:0,
		baseSpeed:6,
		alreadyFired:false,
		torret: torret,
		addBullet:addBullets,
		renderer : renderer,
		
	}

}
func (player *Player) String() string{
	return fmt.Sprintf("{Name:%s, id: %d, pos: %s}",player.name,player.id,player.pos.String())
}
func (player *Player)Move(dx int32 , dy int32){
	newPos  := player.pos.Move(dx,dy)
	player.pos = newPos
}
func (player *Player)Render(renderer *sdl.Renderer,camera *sdl.Rect){
	renderer.CopyEx(player.tankTexture, &sdl.Rect{X:0,Y:0,W:200,H:200}, player.rect, player.tankRotationAngle , nil,sdl.FLIP_NONE);
	renderer.CopyEx(player.torret.torretTexture, &sdl.Rect{X:0,Y:0,W:200,H:200}, player.torret.torretRect, player.torret.rotationAngle , nil,sdl.FLIP_NONE);

}
func (player *Player)handleEvents(eventType,key int){
	action,isValid :=player.keyController.HandleKey(key)
	rotation,move,fire:=0,false,false
	if isValid{
		if eventType == sdl.KEYUP{
		
		}else if eventType == sdl.KEYDOWN{
			switch action{
			case "NORTH":
				rotation = -1
			
			case "SOUTH":
				rotation = 1
				break 
			case "MOVE":
				move = true
				
				break
			case "FIRE":
				fire = true
				break	
			}
		}
	}
	player.rotationSpeed = 10 * rotation
	if move{
		player.xSpeed,player.ySpeed = calculateSpeed(player.baseSpeed,player.tankRotationAngle)
	}else{
		player.xSpeed,player.ySpeed =0,0
	}

	player.fire = fire
}
func (player *Player)Tick(eventType,key int){

	player.handleEvents(eventType,key)
	player.tankRotationAngle=float64(int(player.tankRotationAngle+float64(player.rotationSpeed)) % 360)
	player.torret.rotationAngle=float64(int((player.torret.rotationAngle+float64(player.rotationSpeed)))% 360)
	
	player.rect.X += int32(player.xSpeed)
	player.rect.Y += int32(player.ySpeed)
	player.torret.torretRect.X += int32(player.xSpeed)
	player.torret.torretRect.Y += int32(player.ySpeed)
	if player.fire{
		xOffSet,yOffSet:= calculateSpeed(player.torret.baseTorretSpeed,player.torret.rotationAngle)		
		player.torret.torretRect.Y -= int32(yOffSet)
		player.torret.torretRect.X -= int32(xOffSet)
		player.torret.torretXOffset = int32(xOffSet)
		player.torret.torretYOffset = int32(yOffSet)
		player.alreadyFired =  true
		player.addBullet(MakeBullet("bullet1",&sdl.Rect{X:player.rect.X,Y:player.rect.Y,W:35,H:35},player.renderer,player.blockSize,10,20,10,player.torret.torretRange,player.tankRotationAngle))
		
	}
	if !player.fire && player.alreadyFired{
		player.torret.torretRect.Y += 	player.torret.torretYOffset
		player.torret.torretRect.X += 	player.torret.torretXOffset
		player.torret.torretXOffset = 0
		player.torret.torretYOffset = 0
		player.alreadyFired = false
	}
	
	
}
func calculateSpeed(speed float64,rotationAngle float64) (float64,float64){
	radAngle1 :=  (180-rotationAngle) * math.Pi/180
	
	ySpeed :=math.Cos(radAngle1)* speed
	xSpeed := math.Sin(radAngle1) * speed

	
	return xSpeed,ySpeed
}
 
  
func (player *Player)IsAlive()bool{
	return player.isAlive

}
func (player *Player)Free(){
	player.tankTexture.Destroy()
	player.torret.torretTexture.Destroy()

}