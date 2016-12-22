package main

import (
	_ "reminds/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/bootstrap", "static")
	beego.Run()
}
