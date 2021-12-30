package main
import(
	//"runtime"
	"server/server"
)
func main(){
	server := server.InitServer("127.0.0.1",4444,2048)
	quit := make(chan struct{})
	server.Listen(quit)
	/*for i := 0; i < runtime.NumCPU(); i++ {
			go 
	}*/
	<-quit // hang until an error

}