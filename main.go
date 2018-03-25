package main

import (
	"fmt"

	"github.com/astaxie/beego"
)

// HomeController is a controller that has the get and post methods
type HomeController struct {
	beego.Controller
}

// Get is reloaded
func (c *HomeController) Get() {
	c.Ctx.WriteString("Winter is coming")
}

func main() {
	fmt.Println("This is the main function of this project.")
	beego.Router("/", &HomeController{}) // 注册路由
	beego.Run()                          // 开始运行
}
