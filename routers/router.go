package routers

import (
	"reminds/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/add", &controllers.MainAddController{})
	beego.Router("/msg", &controllers.MainMsgController{})
	beego.Router("/an", &controllers.AndroidCommitController{})
	beego.Router("/all", &controllers.AndroidAllController{}) //android获取所有数据
	beego.Router("/update", &controllers.AndroidUpdateController{})
	beego.Router("/change", &controllers.AndroidChangeController{})
}
