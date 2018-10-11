package main

import (
	"LuaServer"
	"log"
)

func main() {
	defer log.Println("1")
	defer log.Println("sadqwe")
	G_Lua.GetLState().DoString(`print("hello)`)
	//G_Lua.GetLState().
}
