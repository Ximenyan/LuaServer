package main

import (
	"LuaServer"
	"log"
	"reflect"

	"github.com/yuin/gopher-lua"
)

func Test(a, b int) int {
	return a * b
}
func main() {
	log.Println(reflect.TypeOf(Test).String())

	G_Lua.SetGFunc("test", Test, G_Lua.GoParmTypes{"int", "int"}, G_Lua.GoParmTypes{"int"})

	G_Lua.GetLState().DoString(`print(test(123,90))`)

	tab := map[string]lua.LGFunction{
		"mFunc": G_Lua.GoFuncToLFunc(Test, G_Lua.GoParmTypes{"int", "int"}, G_Lua.GoParmTypes{"int"}),
	}
	G_Lua.SetGModule("module", tab, []G_Lua.GoFields{
		G_Lua.GoFields{"Name", false},
		G_Lua.GoFields{"Name1", false},
		G_Lua.GoFields{"Name2", false},
		G_Lua.GoFields{"Name3", false},
	})
	G_Lua.GetLState().DoFile(`./test.lua`)
	G_Lua.CallLuaGFn("max", 3, []G_Lua.GoFields{
		G_Lua.GoFields{"int", 9},
		G_Lua.GoFields{"int", 800}})
	G_Lua.CallLuaGFn("max", 3, []G_Lua.GoFields{
		G_Lua.GoFields{"int", 9},
		G_Lua.GoFields{"int", 800000}})
}
