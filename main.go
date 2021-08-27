package main

import (
	"fmt"
	"github.com/semyon-dev/gpn-tc-backend/config"
	"github.com/semyon-dev/gpn-tc-backend/db"
)

func main() {
	config.Load()
	db.Connect()
	fmt.Println("hello!")
}
