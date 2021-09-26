package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"go-fastdfs-web-go/commons"
	"go-fastdfs-web-go/form"
	"go-fastdfs-web-go/models"
)

type PeersController struct {
	BaseController
}

// Index 集群管理首页
func (c *PeersController) Index() {
	c.TplName = "peers/list.tpl"
}

// PageList 获取集群分页列表
func (c *PeersController) PageList() {
	page := c.Ctx.Input.Query("page")
	limit := c.Ctx.Input.Query("limit")
	peers := models.Peers{}
	pager := peers.PageList(page, limit)
	c.Data["json"] = pager
	c.ServeJSON()
	c.StopRun()
}

// AddPage 集群添加页
func (c *PeersController) AddPage() {
	c.TplName = "peers/add.tpl"
}

// DoAdd 添加集群
func (c *PeersController) DoAdd() {
	var peersForm form.PeersForm
	err := c.ParseForm(&peersForm)
	if err != nil {
		c.ErrorJson(500, "param error", nil)
	}

	c.ValidParam(&peersForm, "添加失败")

	peers, err := peersForm.CheckPeers()
	if err == nil {
		c.ErrorJson(500, "该集群已存在!", nil)
	}

	// 拼装url
	path := peers.ServerAddress
	if peers.GroupName != "" {
		path += "/" + peers.GroupName
	}
	path += commons.ApiStatus
	logs.Info("DoAdd url -> ", path)

	// 测试连接GoFastDfs
	result, err := httpUtil.PostForm(path, nil)
	if err != nil {
		logs.Error("DoAdd -> ", err)
		c.ErrorJson(500, "连接GoFastDfs服务失败!请检查服务地址是否正确,以及是否配置白名单!", nil)
	}

	var resultMap map[string]interface{}
	err = json.Unmarshal([]byte(result), &resultMap)
	if err != nil || resultMap["status"] != commons.ApiStatusSuccess {
		logs.Error("DoAdd json exception -> ", err)
		c.ErrorJson(500, "连接GoFastDfs服务失败!请检查服务地址是否正确,以及是否配置白名单!", nil)
	}
	logs.Info("DoAdd result -> ", resultMap)

	_, err = peersForm.Save()

	if err != nil {
		c.ErrorJson(500, "添加失败!", nil)
	}
	c.SuccessJson(peersForm)
}
