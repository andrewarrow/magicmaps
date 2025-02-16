package magic

import (
	"strings"
	"time"
)

type Map struct {
	Item map[string]any
}

func NewMap(item map[string]any) *Map {
	b := Map{}
	b.Item = item
	return &b
}

func (b *Map) GetBytes(name string) []byte {
	v, _ := b.Item[name].(string)
	return []byte(v)
}
func (b *Map) GetFloat(name string) float64 {
	v, _ := b.Item[name].(float64)
	return v
}
func (b *Map) GetFloatAsInt(name string) int64 {
	v, _ := b.Item[name].(float64)
	return int64(v)
}
func (b *Map) GetInt(name string) int64 {
	v, _ := b.Item[name].(int64)
	return v
}
func (b *Map) GetString(name string) string {
	v, _ := b.Item[name].(string)
	return strings.TrimSpace(v)
}
func (b *Map) GetStringOk(name string) (string, bool) {
	if b.Item[name] == nil {
		return "", false
	}
	v, _ := b.Item[name].(string)
	return v, true
}
func (b *Map) GetMap(name string) map[string]any {
	v, _ := b.Item[name].(map[string]any)
	return v
}
func (b *Map) GetBool(name string) (bool, bool) {
	if b.Item[name] == nil {
		return false, false
	}
	v := b.Item[name].(bool)
	if v == false {
		return false, true
	}
	return true, true
}
func (b *Map) GetSimpleBool(name string) bool {
	v, _ := b.Item[name].(bool)
	return v
}
func (b *Map) GetList(name string) []any {
	v, _ := b.Item[name].([]any)
	return v
}
func (b *Map) GetTime(name string) time.Time {
	v, ok := b.Item[name].(time.Time)
	if ok {
		return v
	}
	v2, ok2 := b.Item[name].(int64)
	if ok2 {
		return time.Unix(v2, 0)
	}
	return time.Now()
}
