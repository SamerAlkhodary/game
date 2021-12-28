package program
import(
	"github.com/veandco/go-sdl2/sdl"
	"game/states"
	"log"
)
type Program struct{
	renderer *sdl.Renderer
	stateManager *states.StateManager
	currentState states.State
	window *sdl.Window
	width int32
	height int32
	frames uint32
}
func Init()*Program{
	blockSize:= int32(100)
	width := 16 * blockSize
	height:= 10* blockSize
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	window, err := sdl.CreateWindow("Game", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Printf("Failed to create renderer: %s\n", err)
	}
	stateManager := states.MakeStateManager()
	stateManager.Init(renderer)
	return &Program{
		renderer:renderer,
		stateManager:stateManager,
		window:window,
		width:width,
		height:height,
		frames : 30,
		currentState:stateManager.GetCurrentState(),
	}
}
func (program *Program)Run(){
	for program.stateManager.IsRunning() {
	event := eventGetter()
	start:=sdl.GetTicks()
	program.renderer.Clear()
	program.currentState.Render()
	program.currentState.Tick(event)
	program.renderer.Present()
	if 1000/program.frames > sdl.GetTicks()-start{
		sdl.Delay(1000/program.frames - (sdl.GetTicks()-start))
	}
}
	
}
func eventGetter()sdl.Event{
	var event sdl.Event
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		return event
	}
	return event
}