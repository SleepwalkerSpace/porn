package main

import (
	"fmt"
	_ "porn/common"
	"porn/common/config"
)

func main() {
	fmt.Printf("%+v", *config.Config)
}
