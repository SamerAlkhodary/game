package model
import(
	
	"github.com/veandco/go-sdl2/sdl"
)


type Torret struct{
	name string
	torretRect *sdl.Rect
	rotationAngle float64
	baseTorretSpeed float64
	torretXOffset int32
	torretYOffset int32
	torretTexture *sdl.Texture
	torretRange int32
	cooldown int32
}
func MakeTorret(name string, torretRect *sdl.Rect,baseTorretSpeed float64 ,torretRange int32,renderer *sdl.Renderer,cooldown int32)*Torret{
	path := "images/torrets/"
	torretSurface := spriteLoader(path+name+".bmp")
	torretTexture:= textureMaker(torretSurface,renderer)
	return &Torret{
		name:name,
		torretRect:torretRect,
		torretRange: torretRange,
		torretTexture: torretTexture,
		torretXOffset:0,
		torretYOffset:0,
		rotationAngle:0,
		baseTorretSpeed:baseTorretSpeed,
		cooldown:cooldown,
	}
}
