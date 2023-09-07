package main

import (
	initialize "github.com/WuLianN/go-toy/initialize"
)

func init() {
	initialize.SetupInit()
}

// @title github.com/WuLianN/go-toy
// @version 1.0
// @description 玩起来！！！
func main() {
	initialize.SetupRouter()
}