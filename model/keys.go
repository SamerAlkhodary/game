package model

type KeyController struct{
	keys map[int]string
 

}
func MakeKeyController(north,south,fire,move int) *KeyController{
	keys:= make(map[int]string)
	keys[north] = "NORTH"
	keys[south] = "SOUTH"
	keys[fire] = "FIRE"
	keys[move] = "MOVE"
	return &KeyController{
		keys:keys,
	}
}
func (keyController *KeyController) HandleKey(key int)(string,bool){
	if action, ok := keyController.keys[key]; ok {
		return action,true
	}
	return "UNDEFIEND",false
}