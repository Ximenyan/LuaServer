package G_Lua

import (
	"log"
	"reflect"
	//"strings"
	"sync"

	"github.com/yuin/gopher-lua"
)

var lState *lua.LState
var initOnce = sync.Once{}

type ModFuncsTab map[string]lua.LGFunction

/*初始化*/
func init() {
	initOnce.Do(func() {
		lState = lua.NewState()
		log.Println(lua.RegistrySize, lua.CallStackSize)
		lua.RegistrySize = 1024 * 20
		lua.CallStackSize = 1024
		log.Println(lua.RegistrySize, lua.CallStackSize)
	})
}

/*获取全局lua虚拟机*/
func GetLState() *lua.LState {
	return lState
}

/**/
func GoFuncToLFunc(f interface{}, parms GoParmTypes, rets GoParmTypes) lua.LGFunction {
	fn := func(L *lua.LState) int {
		args := make([]reflect.Value, len(parms))
		index := 0
		for _, strType := range parms {
			args[index] = reflect.ValueOf(GetParms(L, strType, index+1))
			index++
		}
		ret_vals := reflect.ValueOf(f).Call(args)
		index = 0
		for _, retType := range rets {
			ToParms(L, retType, ret_vals[index])
			index++
		}
		return len(ret_vals)
	}
	return fn
}

/*设置虚拟机栈*/
func SetStack(RegistrySize, CallStackSize int) {
	lua.RegistrySize = RegistrySize
	lua.CallStackSize = CallStackSize
}

/*设置全局 function */
func SetGFunc(funcName string, f interface{}, parms GoParmTypes, rets GoParmTypes) {
	lState.SetGlobal(funcName,
		lState.NewFunction(GoFuncToLFunc(f, parms, rets)))
	return
}

/*设置全局 module */
func SetGModule(modName string, fnTab ModFuncsTab, fields GoParmsField) {
	loader := func(L *lua.LState) int {
		// register functions to the table
		mod := L.SetFuncs(L.NewTable(), fnTab)
		// register other stuff
		if fields != nil {
			//log.Println(fields)
			ltab := fields.GetLTab()
			for _, kv := range ltab {
				L.SetField(mod, kv.Name, kv.Value)
			}
		}
		// returns the module
		L.Push(mod)
		return 1
	}
	lState.PreloadModule(modName, loader)
	return
}

/*调用全局函数*/
func CallLuaGFn(fn string, ret_num int, parms GoParmsField) []lua.LValue {
	args := make([]lua.LValue, len(parms))
	rets := make([]lua.LValue, ret_num)
	if parms != nil {
		ltab := parms.GetLTab()
		idx := 0
		for _, kv := range ltab {
			args[idx] = kv.Value
			idx++
			//L.SetField(mod, kv.Name, kv.Value)
		}
	}
	lState.CallByParam(lua.P{
		Fn:      lState.GetGlobal(fn),
		NRet:    ret_num,
		Protect: true,
	}, args...)
	idx := 0
	for idx < ret_num {
		rets[idx] = lState.Get(-1 - idx)
		idx++
	}
	log.Println(rets)
	return rets
}
