package model
import(
	
	"github.com/veandco/go-sdl2/sdl"
)

type Background struct{
	
	rect *sdl.Rect

}
func MakeBackground(image string,width int32, height int32,renderer *sdl.Renderer) *Background{
	/*surface,err := sdl.LoadBMP(image)
	if err != nil {
		log.Fatal("[SpriteSet] NewSpriteSet(): unable to load resource")
	}
	texture,err:= renderer.CreateTextureFromSurface(surface)
	if err != nil {
		log.Fatal("[SpriteSet] NewSpriteSet(): unable to load resource")
	}
	return &Background{
		rect : &sdl.Rect{0, 0, width, height},
		texture: texture,

	}*/
	return &Background{
		rect : &sdl.Rect{0, 0, width, height},


	}
	
}
func (bg *Background)Render(renderer *sdl.Renderer){
	renderer.SetDrawColor(193, 154, 107, 255)
	renderer.DrawRect(bg.rect)

	/*if camera.X> (3997-camera.W){
		camera.X=0
		

	}
	
	renderer.Copy(bg.texture,camera,bg.rect)*/

}
func (bg *Background) Tick(eventType,key int){
		

}
func (bg *Background)IsAlive()bool{
	return true

}
func (bg *Background)Free(){
	

}
func (bg *Background)HandleCollision(other Entity){
	

}

func (bg *Background)IsRigid()bool{
	return false

}
func (bg *Background)GetRect()*sdl.Rect{
	return nil

}
func (bg *Background)GetRotationAngle()float64{
	return 0

}