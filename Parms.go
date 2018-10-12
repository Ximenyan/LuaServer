package G_Lua

import (
	"reflect"

	"github.com/yuin/gopher-lua"
)

/*参数*/
type GoFields struct {
	Name  string
	Value interface{}
}
type LuaFields struct {
	Name  string
	Value lua.LValue
}

/*参数表*/
type GoParmsField []GoFields
type LuaParmsField []LuaFields

type GoParmTypes []string

/*函数获取参数*/
func GetParms(L *lua.LState, strType string, index int) interface{} {
	switch strType {
	case "string":
		return L.ToString(index)
	case "bool":
		return L.ToBool(index)
	case "int":
		return L.ToInt(index)
	case "int8":
		return int8(L.ToInt(index))
	case "int16":
		return int16(L.ToInt(index))
	case "int32":
		return int32(L.ToInt(index))
	case "int64":
		return L.ToInt64(index)
	case "uint":
		return uint(L.ToInt(index))
	case "uint8":
		return uint8(L.ToInt(index))
	case "uint16":
		return uint16(L.ToInt(index))
	case "uint32":
		return uint32(L.ToInt(index))
	case "uint64":
		return uint64(L.ToInt(index))
	case "float32":
		return float32(L.ToNumber(index))
	case "float64":
		return float64(L.ToNumber(index))
	default:
		return nil
	}
}

/*函数获取参数*/
func ToParms(L *lua.LState, strType string, ret_val reflect.Value) {
	switch strType {
	case "string":
		L.Push(lua.LString(ret_val.String()))
		return
	case "bool":
		L.Push(lua.LBool(ret_val.Bool()))
		return
	case "int":
		L.Push(lua.LNumber(ret_val.Int()))
		return
	case "int8":
		L.Push(lua.LNumber(ret_val.Int()))
		return
	case "int16":
		L.Push(lua.LNumber(ret_val.Int()))
		return
	case "int32":
		L.Push(lua.LNumber(ret_val.Int()))
		return
	case "int64":
		L.Push(lua.LNumber(ret_val.Int()))
		return
	case "uint":
		L.Push(lua.LNumber(ret_val.Uint()))
		return
	case "uint8":
		L.Push(lua.LNumber(ret_val.Uint()))
		return
	case "uint16":
		L.Push(lua.LNumber(ret_val.Uint()))
		return
	case "uint32":
		L.Push(lua.LNumber(ret_val.Uint()))
		return
	case "uint64":
		L.Push(lua.LNumber(ret_val.Uint()))
		return
	case "float32":
		L.Push(lua.LNumber(ret_val.Float()))
		return
	case "float64":
		L.Push(lua.LNumber(ret_val.Float()))
		return
	default:
		L.Push(lua.LNil)
		return
	}
}

/*模块获取变量*/
func (this GoParmsField) GetLTab() LuaParmsField {
	arr := []LuaFields{}
	for i := 0; i < len(this); i++ {
		field := this[i]
		if reflect.TypeOf(field.Value) == nil {
			arr = append(arr, LuaFields{field.Name, lua.LNil})
		}

		strType := reflect.TypeOf(field.Value).String()
		switch strType {
		case "string":
			arr = append(arr, LuaFields{field.Name, lua.LString(field.Value.(string))})
			continue
		case "bool":
			arr = append(arr, LuaFields{field.Name, lua.LBool(field.Value.(bool))})
			continue
		case "int":
			fNum := float64(field.Value.(int))
			arr = append(arr, LuaFields{field.Name, lua.LNumber(float64(fNum))})
			continue
		case "int8":
			fNum := float64(field.Value.(int8))
			arr = append(arr, LuaFields{field.Name, lua.LNumber(float64(fNum))})
			continue
		case "int16":
			fNum := float64(field.Value.(int16))
			arr = append(arr, LuaFields{field.Name, lua.LNumber(float64(fNum))})
			continue
		case "int32":
			fNum := float64(field.Value.(int32))
			arr = append(arr, LuaFields{field.Name, lua.LNumber(float64(fNum))})
			continue
		case "int64":
			fNum := float64(field.Value.(int64))
			arr = append(arr, LuaFields{field.Name, lua.LNumber(float64(fNum))})
			continue
		case "uint":
			fNum := float64(field.Value.(uint))
			arr = append(arr, LuaFields{field.Name, lua.LNumber(float64(fNum))})
			continue
		case "uint8":
			fNum := float64(field.Value.(uint8))
			arr = append(arr, LuaFields{field.Name, lua.LNumber(float64(fNum))})
			continue
		case "uint16":
			fNum := float64(field.Value.(uint16))
			arr = append(arr, LuaFields{field.Name, lua.LNumber(float64(fNum))})
			continue
		case "uint32":
			fNum := float64(field.Value.(uint32))
			arr = append(arr, LuaFields{field.Name, lua.LNumber(float64(fNum))})
			continue
		case "uint64":
			fNum := float64(field.Value.(uint64))
			arr = append(arr, LuaFields{field.Name, lua.LNumber(float64(fNum))})
			continue
		case "float32":
			fNum := float64(field.Value.(float32))
			arr = append(arr, LuaFields{field.Name, lua.LNumber(float64(fNum))})
			continue
		case "float64":
			fNum := float64(field.Value.(float64))
			arr = append(arr, LuaFields{field.Name, lua.LNumber(float64(fNum))})
			continue
		default: /*
				if strings.Index(strType, "int") > 0 {
				}
				if strings.Index(strType, "float") > 0 {
					arr = append(arr, LuaFields{field.Name, lua.LNumber(field.Value.(float64))})
				}*/
			continue
		}
	}
	return arr
}
