package wraper

import (
	"fmt"
	"testing"
)

func TestWrape_WrapStringMacToInt64(t *testing.T) {
	mac := `01-23-45-67-89-ab`
	Init()
	fmt.Println(Wraper.WrapStringMacToInt64(mac))
}
