package model
import(
	
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

type Tile struct{
	texture *sdl.Texture
	rect *sdl.Rect
	collisionRect *sdl.Rect
	from *sdl.Rect
	isAlive bool
	isRigid bool

}
func MakeTile(i,j int32, tileType int32, renderer *sdl.Renderer,blockSize int32) *Tile{
	filePath :="images/grass/1grass1.bmp"
	var surface *sdl.Surface
	var texture *sdl.Texture
	from := &sdl.Rect{X:0,Y:0,W:100,H:100}
	collisionRect:= &sdl.Rect{X:int32(i*blockSize+blockSize*20/100), Y:j*blockSize+blockSize*20/100, W:blockSize-blockSize*40/100, H:blockSize-blockSize*40/100}
	var err error
	isRigid:=false
	switch tileType{
	case 0:
		return nil
		
	case 1:
		filePath ="images/grass/1grass1.bmp"
		isRigid = true
		collisionRect = &sdl.Rect{X:i*blockSize, Y:j*blockSize, W:blockSize, H:blockSize}
		
	break


	case 3:
		filePath ="images/earthTiles/31earth1.bmp"

		
	break
	case 4:
		filePath ="images/items/tree2.bmp"
		isRigid = true

		

	break
	case 5:
		filePath ="images/items/lake.bmp"
		isRigid =true

	
	break
	case 6:
		filePath ="images/items/dirt.bmp"

	break
	case 7:
		filePath ="images/items/wall.bmp"
		

	break
	}
	surface = spriteLoader(filePath)
	err= surface.SetColorKey(true, sdl.MapRGB(surface.Format, 0x00,0xFF,0xFF)) 
	if  err!=nil{
		log.Printf( "Unable to set Color Key!")
	}
	texture,err= renderer.CreateTextureFromSurface(surface)
	if err != nil {
		log.Fatal("[SpriteSet] NewSpriteSet(): unable to load resource",err)
	}
	
	
	return &Tile{
		rect : &sdl.Rect{X:i*blockSize, Y:j*blockSize, W:blockSize, H:blockSize},
		collisionRect : collisionRect,
		texture: texture,
		from:from,	
		isAlive:true,
		isRigid: isRigid,

	}
}
func (tile *Tile) Render(renderer *sdl.Renderer,camera *sdl.Rect){
	dest := &sdl.Rect{X:tile.rect.X,Y:tile.rect.Y,W:tile.rect.W,H:tile.rect.H}
	dest.X -=camera.X
	dest.Y -= camera.Y
	renderer.Copy(tile.texture,tile.from,dest)
	renderer.SetDrawColor(0, 255, 0, 255)
	renderer.DrawRect(tile.collisionRect)
	renderer.SetDrawColor(193, 154, 107, 255)
	
	
	
}
func (tile *Tile) Tick(eventType,key int){
	
}
func spriteLoader(filePath string) *sdl.Surface{
	surface,err := sdl.LoadBMP(filePath)
	if err != nil {
		log.Fatal("[SpriteSet] NewSpriteSet(): unable to load resource",filePath)
	}
	return surface

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
func (tile *Tile)IsAlive()bool{
	return tile.isAlive
}
func (tile *Tile)Free(){
	tile.texture.Destroy()

}
func (tile *Tile)HandleCollision(other Entity){
	

}
func(tile *Tile)IsRigid() bool{
	return tile.isRigid
}
func (tile *Tile)GetRect()*sdl.Rect{
	return tile.rect

}
func (tile *Tile)GetRotationAngle()float64{
	return 0

}