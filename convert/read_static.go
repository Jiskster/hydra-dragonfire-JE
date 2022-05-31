package convert

import (
	"github.com/anon55555/mt"
	"github.com/yuin/gopher-lua"
)

//go:generate ./read_mkauto.lua

func ReadBool(l *lua.LState, val lua.LValue, ptr *bool) {
	if val.Type() != lua.LTBool {
		panic("invalid value for bool: must be a boolean")
	}
	*ptr = bool(val.(lua.LBool))
}

func ReadString(l *lua.LState, val lua.LValue, ptr *string) {
	if val.Type() != lua.LTString {
		panic("invalid value for string: must be a string")
	}
	*ptr = string(val.(lua.LString))
}

func ReadSliceByte(l *lua.LState, val lua.LValue, ptr *[]byte) {
	if val.Type() != lua.LTString {
		panic("invalid value for []byte: must be a string")
	}
	*ptr = []byte(val.(lua.LString))
}

func ReadSliceField(l *lua.LState, val lua.LValue, ptr *[]mt.Field) {
	if val.Type() != lua.LTTable {
		panic("invalid value for []Field: must be a table")
	}
	val.(*lua.LTable).ForEach(func(k, v lua.LValue) {
		if k.Type() != lua.LTString || v.Type() != lua.LTString {
			panic("invalid value for Field: key and value must be strings")
		}
		*ptr = append(*ptr, mt.Field{Name: string(k.(lua.LString)), Value: string(v.(lua.LString))})
	})
}

func ReadPointedThing(l *lua.LState, val lua.LValue, ptr *mt.PointedThing) {
	if val.Type() != lua.LTTable {
		panic("invalid value for PointedThing: must be a table")
	}
	id := l.GetField(val, "id")

	if id != lua.LNil {
		pt := &mt.PointedAO{}
		ReadAOID(l, id, &(*pt).ID)
		*ptr = pt
	} else {
		pt := &mt.PointedNode{}
		ReadVec3Int16(l, l.GetField(val, "under"), &(*pt).Under)
		ReadVec3Int16(l, l.GetField(val, "above"), &(*pt).Above)
		*ptr = pt
	}
}
