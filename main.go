package main
import (
	game"game/game"
	"bufio"
    "log"
    "os"
	"strings"
	"strconv"

)
func main(){
	tiles:=readMap()
	game:= game.Init(16*100,10*100,100,tiles)
	game.Run()
}
func readMap()[][]int32{
	file, err := os.Open("map.txt")
  
    if err != nil {
        log.Fatalf("failed to open")
  
    }
	var tiles [][]int32
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
 
	defer  file.Close()
	for _,line := range(lines){
		line = strings.TrimSpace(line)
		var elems []int32 
		row:= strings.Split(line," ")
		for _,elem := range(row){
			intElem,_:=strconv.Atoi(elem)
			elems = append(elems,int32(intElem))
		}
		tiles = append(tiles,elems)
	}
	return tiles

    

}