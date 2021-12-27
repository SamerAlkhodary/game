package model
import(
	
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

type Tile struct{
	texture *sdl.Texture
	rect *sdl.Rect
	from *sdl.Rect
	isAlive bool

}
func MakeTile(i,j int32, tileType int32, renderer *sdl.Renderer,blockSize int32) *Tile{
	filePath :="images/grass/1grass1.bmp"
	var surface *sdl.Surface
	var texture *sdl.Texture
	from := &sdl.Rect{X:0,Y:0,W:100,H:100}
	var err error
	switch tileType{
	case 0:
		return nil
		
	case 1:
		filePath ="images/grass/1grass1.bmp"
		
	break
	case 2:
		filePath ="images/waterTiles/11water1.bmp"


	break
	case 3:
		filePath ="images/earthTiles/31earth1.bmp"

		
	break
	case 4:
		filePath ="images/items/tree2.bmp"

		

	break
	case 5:
		filePath ="images/items/lake.bmp"

	
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
		texture: texture,
		from:from,	
		isAlive:true,

	}
}
func (tile *Tile) Render(renderer *sdl.Renderer,camera *sdl.Rect){
	renderer.Copy(tile.texture,tile.from,tile.rect)

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