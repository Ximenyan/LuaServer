package G_Lua

import (
	"log"
	"sync"

	"github.com/yuin/gopher-lua"
)

var lState *lua.LState
var initOnce = sync.Once{}

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

/*设置虚拟机栈*/
func SetStack(RegistrySize, CallStackSize int) {
	lua.RegistrySize = RegistrySize
	lua.CallStackSize = CallStackSize
}

/*G function */
func SetGFunc() {

}
