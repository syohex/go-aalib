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

func TestFlush(t *testing.T) {
	handle, _ := Init(100, 100, AA_NORMAL)
	if handle == nil {
		t.Errorf("Return nil from constructor")
	}
	handle.Flush()
}

func TestResize(t *testing.T) {
	handle, _ := Init(100, 100, AA_NORMAL)
	if handle == nil {
		t.Errorf("Return nil from constructor")
	}
	ret, _ := handle.Resize()
	if ret == -1 {
		t.Errorf("Failed: Resize()")
	}
}

// Difficult tests for Text(), Attrs(), Image()

func TestScrHeight(t *testing.T) {
	handle, _ := Init(100, 100, AA_NORMAL)
	if handle == nil {
		t.Errorf("Return nil from constructor")
	}
	scrwidth := handle.ScrHeight()
	if scrwidth == 0 {
		t.Errorf("Error")
	}
}


func TestImgWidth(t *testing.T) {
	handle, _ := Init(100, 100, AA_NORMAL)
	if handle == nil {
		t.Errorf("Return nil from constructor")
	}
	scrwidth := handle.ScrHeight()
	if scrwidth == 0 {
		t.Errorf("Error")
	}
}

func TestImgHeight(t *testing.T) {
	handle, _ := Init(100, 100, AA_NORMAL)
	if handle == nil {
		t.Errorf("Return nil from constructor")
	}
	scrwidth := handle.ImgHeight()
	if scrwidth == 0 {
		t.Errorf("Error")
	}
}
