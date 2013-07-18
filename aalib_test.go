package aalib

import (
	"testing"
)

func TestInit(t *testing.T) {
	handle, _ := Init(100, 100, AA_NORMAL)
	if handle == nil {
		t.Errorf("Return nil from constructor")
	}
}

func TestClose(t *testing.T) {
	handle, _ := Init(100, 100, AA_NORMAL)
	if handle == nil {
		t.Errorf("Return nil from constructor")
	}
	handle.Close()
}

func TestImgWidth(t *testing.T) {
	handle, _ := Init(100, 100, AA_NORMAL)
	if handle == nil {
		t.Errorf("Return nil from constructor")
	}
	scrwidth := handle.ImgWidth()
	if scrwidth == 0 {
		t.Errorf("Error")
	}
}
