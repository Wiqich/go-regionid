package regionid

import (
	"testing"
)

func TestBuiltinWorld(t *testing.T) {
	if err := LoadBuiltinWorld(); err != nil {
		t.Error("load build world fail:", err.Error())
		return
	}
	shanghai := GetLocation("中国", "上海", "")
	if shanghai == nil || shanghai.Country().Iso().Numeric != 156 || shanghai.Subdivision().ID() != 310000 || shanghai.City() != nil {
		t.Error("unexpected shanghai location:", shanghai)
		return
	}
	fuzhou := GetLocation("中国", "福建", "福州")
	t.Log("fuzhou:", fuzhou)
	if fuzhou == nil || fuzhou.Country().Iso().Numeric != 156 || fuzhou.Subdivision().ID() != 350000 || fuzhou.City().ID() != 350100 {
		t.Error("unexpected fuzhou location:", fuzhou)
		return
	}
}
