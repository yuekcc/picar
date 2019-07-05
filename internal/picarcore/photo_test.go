package picarcore

import (
	"testing"
)

func Test_getNameFromFilename(t *testing.T) {
	testdata := []string{
		"IMG_20151106_212111.jpg",
		"IMG-20151106-212111.jpg",
		"20151106 150215.jpg",
		"main.jpg",
		"first.second.third.jpg",
	}

	for _, data := range testdata {
		dt, _ := fromFilename(data)
		t.Log(dt)
	}
}
