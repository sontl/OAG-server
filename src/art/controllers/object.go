package controllers

import (
	"art/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

// Operations about object
type ObjectController struct {
	beego.Controller
}

// @Title create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (this *ObjectController) Post() {
	beego.Info("Post")
	var ob models.Object
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	objectid := models.AddOne(ob)
	this.Data["json"] = map[string]string{"ObjectId": objectid}
	this.ServeJson()
}

// @Title create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [options]
func (this *ObjectController) Options() {
	beego.SetLogger("file", `{"filename":"//Users/sontl/go/logs/controller.log"}`)

	obs := models.GetAll()
	this.Data["json"] = obs

	beego.Warn("Warn")
	file, fileHeader, err := this.GetFile("file")
	beego.Info("file Uploaded: ", fileHeader.Filename)
	if file != nil {
		beego.Info("file is not null", fileHeader.Filename)
	}
	if err != nil {
		beego.Info("err happen ", err.Error())
	}
	this.SaveToFile("file", "/Users/sontl/go/uploaded/test.jpg")
	// allow cross domain AJAX requests
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	this.ServeJson()
}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (this *ObjectController) Get() {
	beego.Info("Get")
	objectId := this.Ctx.Input.Params[":objectId"]
	if objectId != "" {
		ob, err := models.GetOne(objectId)
		if err != nil {
			this.Data["json"] = err
		} else {
			this.Data["json"] = ob
		}
	}
	this.ServeJson()
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (this *ObjectController) GetAll() {
	obs := models.GetAll()
	this.Data["json"] = obs
	beego.Trace("GetAll")
	this.ServeJson()
}

// @Title update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (this *ObjectController) Put() {
	objectId := this.Ctx.Input.Params[":objectId"]
	var ob models.Object
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)

	err := models.Update(objectId, ob.Score)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = "update success!"
	}
	this.ServeJson()
}

// @Title delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (this *ObjectController) Delete() {
	objectId := this.Ctx.Input.Params[":objectId"]
	models.Delete(objectId)
	this.Data["json"] = "delete success!"
	this.ServeJson()
}
