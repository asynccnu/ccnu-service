package service

import (
	"fmt"
	"testing"
)

func TestGetModulusAndExpoent(t *testing.T) {
	got, got1 := GetModulusAndExpoent()
	fmt.Println(got)
	fmt.Println(got1)
}

func TestEncodeMM(t *testing.T) {
	p := "zhao1638678192%"
	got := EncodeMM(p)
	fmt.Println(got)
}
