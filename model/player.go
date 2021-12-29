package model
import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/ttf"
	"math"
	"log"
)
 type Player struct{
	name string
	id int 
	pos *Pos
	rect *sdl.Rect
	healthBarRect *sdl.Rect
	healthBarBackgroundRect *sdl.Rect
	blockSize int32
	xSpeed float64
	ySpeed float64
	baseSpeed float64
	nameTexture *sdl.Texture
	nameRect *sdl.Rect
	health int32
	rotationSpeed int
	rotation int
	tankTexture *sdl.Texture
	tankRotationAngle float64
	keyController * KeyController
	fire bool
	alreadyFired bool
	torret *Torret
	isAlive bool
	renderer *sdl.Renderer
	addBullet func(Entity)
	isRigid bool
	collisionRect *sdl.Rect
	chunks []*mix.Chunk
	hasPlayedSounds []bool
	move bool


}
func MakePlayer(name string, id int,rect *sdl.Rect,renderer *sdl.Renderer,blockSize int32,keyController *KeyController, addBullets func(Entity))*Player{
	torret := MakeTorret("torret1",&sdl.Rect{X:rect.X,Y:rect.Y,W:rect.W,H:rect.H},4,500,renderer,1)
	pos := MakePos(rect.X,rect.Y)
	tankPath := "images/tank/"
	tankSurface := spriteLoader(tankPath+"tank.bmp")
	tankTexture:= textureMaker(tankSurface,renderer)
	ttf.Init()
	sans,_ := ttf.OpenFont("fonts/SansBold.ttf", 28);
	surface,_ := sans.RenderUTF8Solid(name,sdl.Color{R:0,G:0,B:0,A:255})
	texture,_:=renderer.CreateTextureFromSurface( surface);
	healthBarBackgroundRect:= &sdl.Rect{X: blockSize, Y:blockSize/4,W:5*blockSize,H:blockSize/2}
	mix.Init(mix.INIT_FLAC)
	mix.OpenAudio(mix.DEFAULT_FREQUENCY,mix.DEFAULT_FORMAT,mix.DEFAULT_CHANNELS,mix.DEFAULT_CHUNKSIZE)
	chunk1,err:= mix.LoadWAV("audio/movingTank.wav")
	chunk2,err:= mix.LoadWAV("audio/engine.wav")
	chunk3,err:= mix.LoadWAV("audio/fire.wav")
	
	chunks:= make([]*mix.Chunk,0)
	if err!=nil{
		log.Println(err)
	}
	chunks = append(chunks,chunk1)
	chunks = append(chunks,chunk2)
	chunks = append(chunks,chunk3)
	mix.Volume(1,10)
	mix.Volume(0,50)
	mix.Volume(2,50)
	return &Player{
		name:name,
		tankTexture: tankTexture,
		id:id,
		pos:pos,
		rotation:0,
		isAlive:true,
		fire : false,
		health: 100,
		move:false,
		keyController:keyController,
		rect:rect,
		collisionRect: &sdl.Rect{X:rect.X+blockSize*10/100,Y:rect.Y+blockSize*10/100,W:rect.W-blockSize*25/100,H:rect.H-blockSize*25/100},
		healthBarRect: &sdl.Rect{X: blockSize, Y:blockSize/4,W:5*blockSize,H:blockSize/2},
		healthBarBackgroundRect:healthBarBackgroundRect,
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
		isRigid:true,
		hasPlayedSounds: []bool{false,false,false},
		chunks:chunks,
		nameTexture:texture,
		nameRect: &sdl.Rect{X:(healthBarBackgroundRect.X+healthBarBackgroundRect.W)/2,Y:(healthBarBackgroundRect.Y+healthBarBackgroundRect.H)/3,W:blockSize,H:blockSize/2},
		
	}

}

func (player *Player) String() string{
	return fmt.Sprintf("{Name:%s, id: %d, pos: %s}",player.name,player.id,player.pos.String())
}

func (player *Player)Render(renderer *sdl.Renderer){
	renderer.CopyEx(player.tankTexture, &sdl.Rect{X:0,Y:0,W:200,H:200}, player.rect, player.tankRotationAngle , nil,sdl.FLIP_NONE);
	renderer.CopyEx(player.torret.torretTexture, &sdl.Rect{X:0,Y:0,W:200,H:200}, player.torret.torretRect, player.torret.rotationAngle , nil,sdl.FLIP_NONE);
	renderer.SetDrawColor(255, 0, 0, 255)
	renderer.FillRect(player.healthBarBackgroundRect)
	renderer.SetDrawColor(0, 255, 0, 255)
	renderer.FillRect(player.healthBarRect)
	renderer.SetDrawColor(193, 154, 107, 255)
	renderer.SetDrawColor(0, 255, 0, 255)
	renderer.DrawRect(player.collisionRect)
	renderer.SetDrawColor(193, 154, 107, 255)
	renderer.Copy(player.nameTexture,nil,player.nameRect)

}
func (player *Player)handleEvents(eventType,key int){
	action,isValid :=player.keyController.HandleKey(key)
	if isValid{
		if eventType == sdl.KEYUP{
			switch action{
				case "MOVE":
					player.move = false
					break
				case "NORTH":
					player.rotation = 0
							
				case "SOUTH":
					player.rotation = 0
					break 
				case "FIRE":
					player.fire = false
					break	
				
			}
		
		}else if eventType == sdl.KEYDOWN{
			switch action{
			case "NORTH":
				player.rotation = -1
			
			case "SOUTH":
				player.rotation = 1
				break 
			case "MOVE":
				player.move = true
				
				break
			case "FIRE":
				player.fire = true
				player.chunks[2].Play(2,0)				

				break	
			}
		}
	}
	player.rotationSpeed = 10 * player.rotation
	if player.move{
		if !player.hasPlayedSounds[0]{
			player.chunks[0].Play(0,10)
			player.hasPlayedSounds[0]=true
		}
		player.xSpeed,player.ySpeed = calculateSpeed(player.baseSpeed,player.tankRotationAngle)
		mix.Pause(1)
		mix.Resume(0)
	}else{
		mix.Pause(0)
		mix.Resume(1)
		if !player.hasPlayedSounds[1]{
			player.chunks[1].Play(1,10)
			player.hasPlayedSounds[1]=true
		}
		player.xSpeed,player.ySpeed =0,0
	}

 
}
func(player *Player) Move(){
	player.rect.X += int32(player.xSpeed)
	player.rect.Y += int32(player.ySpeed)
	player.collisionRect.X += int32(player.xSpeed)
	player.collisionRect.Y += int32(player.ySpeed)
	player.torret.torretRect.X += int32(player.xSpeed)
	player.torret.torretRect.Y += int32(player.ySpeed)
}
func (player *Player) Fire(){

	if player.fire &&  !player.alreadyFired{
		

		xOffSet,yOffSet:= calculateSpeed(player.torret.baseTorretSpeed,player.torret.rotationAngle)		
			player.torret.torretRect.Y -= int32(yOffSet)
			player.torret.torretRect.X -= int32(xOffSet)
			player.torret.torretXOffset = int32(xOffSet)
			player.torret.torretYOffset = int32(yOffSet)
			player.alreadyFired =  true
			player.addBullet(MakeBullet("bullet1",&sdl.Rect{X:player.rect.X,Y:player.rect.Y,W:1 *player.blockSize/3,H:1*player.blockSize/3},
			player.renderer,player.blockSize,10,20,10,player.torret.torretRange,player.tankRotationAngle))
		}
		if !player.fire && player.alreadyFired{
			player.torret.torretRect.Y += 	player.torret.torretYOffset
			player.torret.torretRect.X += 	player.torret.torretXOffset
			player.torret.torretXOffset = 0
			player.torret.torretYOffset = 0
			player.alreadyFired = false

		}
}
func (player *Player)Tick(eventType,key int){

	player.handleEvents(eventType,key)
	player.tankRotationAngle=float64(int(player.tankRotationAngle+float64(player.rotationSpeed)) % 360)
	player.torret.rotationAngle=float64(int((player.torret.rotationAngle+float64(player.rotationSpeed)))% 360)
	player.Move()

	player.Fire()
	
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
	player.nameTexture.Destroy()
	for _,chunk := range(player.chunks){
		chunk.Free()

	}

}
func (player *Player)HandleCollision(other Entity){
	
	if other.IsRigid(){
		tile,done := other.(*Tile)
		if done{
			collision:=tile.collisionRect.HasIntersection(player.collisionRect)
			if collision{
				player.xSpeed *= -1
				player.ySpeed *= -1
				player.health -= 5
				if player.health <=0{
					player.health = 0
				}
				player.healthBarRect.W -= player.healthBarRect.W *(100- (player.health*100/100))/100
				player.Move()
			}
		}

	}	

}
func(player *Player)IsRigid()bool{
	return player.isRigid
}
func (player *Player)GetRect()*sdl.Rect{
	return player.rect

}