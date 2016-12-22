package controllers

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var o orm.Ormer

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:qq5566@/remind?charset=utf8")
	orm.RegisterModel(new(anData))
	o = orm.NewOrm()
	o.Using("remind")
}

type MainController struct {
	beego.Controller
}
type AndroidCommitController struct {
	beego.Controller
}
type AndroidAllController struct {
	beego.Controller
}
type AndroidDelController struct {
	beego.Controller
}
type MainAddController struct {
	beego.Controller
}
type MainMsgController struct {
	beego.Controller
}

func (c *MainMsgController) Get() {
	c.Data["Msg"] = c.Input().Get("msg")
	c.TplName = "msg.html"
}
func (c *MainAddController) Get() {
	c.TplName = "add.html"
}

// 處理網頁提交過來的數據
func (c *MainAddController) Post() {
	d := anData{}
	if err := c.ParseForm(&d); err != nil {
		log.Println(err.Error())
		c.Redirect("/msg?msg="+"服務器出錯啦", 302)
	} else {
		d.Time = getCurTime()
		_, e := o.Insert(&d)
		if e != nil {
			log.Println(e.Error())
			c.Redirect("/msg?msg="+"服務器出錯啦", 302)
		}
		c.Redirect("/msg?msg="+"添加成功", 302)
	}
}

func (c *MainController) Get() {
	delId := c.Input().Get("del")
	if delId == "" {
		c.Data["Datas"] = getAll()
		c.TplName = "index.html"
	} else {
		err := delById(delId)
		if err != nil {
			c.Redirect("/msg?msg="+"刪除失敗", 302)
		}
		c.Redirect("/msg?msg="+"刪除成功", 302)
	}
}

type anData struct {
	Id      int
	Title   string `form:"title"`
	Content string `form:"content"`
	Time    string
}

// 处理android客户端提交过来的数据
func (c *AndroidCommitController) Post() {
	d := anData{}
	if err := c.ParseForm(&d); err != nil {
		log.Println(err.Error())
	} else {
		d.Time = getCurTime()
		pk, e := o.Insert(&d)
		if e != nil {
			log.Println(e.Error())
			return
		}
		c.Ctx.WriteString(strconv.Itoa(int(pk)))
	}
}
func (c *AndroidAllController) Get() {
	c.Data["json"] = getAll()
	c.ServeJSON()
}

func (c *AndroidDelController) Get() {
	err := delById(c.Input().Get("id"))
	if err != nil {
		c.Ctx.WriteString("0")
	} else {
		c.Ctx.WriteString("1")
	}
}

/* ---------------------- ------------------------*/
// 獲取所有
func getAll() []*anData {
	var datas []*anData
	_, err := o.QueryTable("an_data").All(&datas)
	if err != nil {
		log.Println(err.Error())
	}
	return datas
}

//根據id刪除
func delById(id string) error {
	i, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error)
		return err
	}
	var num int64
	if num, err = o.Delete(&anData{Id: i}); err != nil {
		log.Println(err.Error)
		return err
	}
	if num == 0 {
		return errors.New("沒有對應的行")
	}
	return nil
}
func getCurTime() string {
	t := time.Now().UTC()
	timestamp := t.Unix()
	_, offset := t.Zone()
	currenttime := time.Unix(timestamp+int64(offset), 0)
	return currenttime.Format("2006-01-02 15:04")
}
