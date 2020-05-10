package model

import (
	"fmt"
	"testing"
)

func TestBase(t *testing.T) {
	c := new(Category)
	c.Id = "5e9cffc2d401ca38ac109fe1"
	c.SetObjectId()
	fmt.Printf("%+v", c)
}

func TestDrop(t *testing.T) {

}
