package main

import (
	"github.com/yuin/gopher-lua"
)

func Double(L *lua.LState) int {
	lv := L.ToInt(1)             /* get argument */
	ls := L.ToInt(2)             /* get argument */
	L.Push(lua.LNumber(lv * ls)) /* push result */
	return 1                     /* number of results */
}
func main() {
	L := lua.NewState()
	defer L.Close()
	L.SetGlobal("double", L.NewFunction(Double)) /* Original lua_setglobal uses stack... */
	if err := L.DoString(`a=50`); err != nil {
		panic(err)
	}
	if err := L.DoString(`a=1`); err != nil {
		panic(err)
	}
	if err := L.DoString(`print(double(a,30))`); err != nil {
		panic(err)
	}
}
