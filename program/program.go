package program
import(
	"github.com/veandco/go-sdl2/sdl"
	"game/states"
	"log"
	"game/network"
)
type Program struct{
	renderer *sdl.Renderer
	client *network.Client
	stateManager *states.StateManager
	currentState states.State
	window *sdl.Window
	width int32
	height int32
	frames uint32 
}
func Init()*Program{
	blockSize:= int32(80)
	client := network.CreateClient(2048,"127.0.0.1","4444")
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
	stateManager := states.MakeStateManager(client,width,height,blockSize)
	stateManager.Init(renderer)
	return &Program{
		renderer:renderer,
		stateManager:stateManager,
		window:window,
		client:client,
		width:width,
		height:height,
		frames : 30,
		currentState:stateManager.GetCurrentState(),
	}
}
func (program *Program)Run(){
	for program.stateManager.IsRunning() {
	program.currentState = program.stateManager.GetCurrentState()
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
	if program.client.IsOnline(){
		program.client.Send(&network.CloseConnectionRequest{
			PlayerId: program.stateManager.PlayerId(),
			GameId: program.stateManager.GameId(),
		})

	}
	
	
}
func eventGetter()sdl.Event{
	var event sdl.Event
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		return event
	}
	return event
}