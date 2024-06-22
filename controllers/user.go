package controllers

import (
	"chat-im/global"
	"chat-im/models"
	"chat-im/pkg"
	"github.com/astaxie/beego"
	"gorm.io/gorm"
	"strconv"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Register() {
	username := c.GetString("username")
	if username == "" {
		c.Data["json"] = global.ResponseData{
			Code:    -500,
			Message: "用户名不能为空",
		}
		c.ServeJSON(true)
		return
	}
	password := c.GetString("password")
	if password == "" {
		c.Data["json"] = global.ResponseData{
			Code:    -500,
			Message: "密码不能为空",
		}
		c.ServeJSON(true)
		return
	}
	err := pkg.VerifyPwd(password)
	if err != nil {
		c.Data["json"] = global.ResponseData{
			Code:    -500,
			Message: err.Error(),
		}
		c.ServeJSON(true)
		return
	}
	// 开始注册
	userInfo := models.User{
		Account:  username,
		Nickname: "",
		Password: password,
		Salt:     "",
		Mobile:   "",
		Email:    "",
		Avatar:   "",
		Status:   1,
	}
	user := userInfo.CreateUser()
	if user.Error != nil {
		c.Data["json"] = global.ResponseData{
			Code:    -500,
			Message: "用户创建失败",
		}
		c.ServeJSON(true)
		return
	}
	c.Data["json"] = global.ResponseData{
		Code:    0,
		Message: "用户创建成功",
	}
	c.ServeJSON(true)
	return
}

func (c *UserController) PwdStrength() {
	password := c.GetString("password")
	verify := pkg.PwdStrongVerify(password)
	c.Data["json"] = global.ResponseData{
		Code:    200,
		Message: "密码强度为" + strconv.Itoa(verify) + "级",
	}
	c.ServeJSON(true)
}

func (c *UserController) Login() {
	username := c.GetString("username")
	if username == "" {
		c.Data["json"] = global.ResponseData{
			Code:    -500,
			Message: "用户名不能为空",
		}
		c.ServeJSON(true)
		return
	}
	password := c.GetString("password")
	if password == "" {
		c.Data["json"] = global.ResponseData{
			Code:    -500,
			Message: "密码不能为空",
		}
		c.ServeJSON(true)
		return
	}
	err := pkg.VerifyPwd(password)
	if err != nil {
		c.Data["json"] = global.ResponseData{
			Code:    -500,
			Message: err.Error(),
		}
		c.ServeJSON(true)
		return
	}
	// 登录
	userInfo := models.User{
		Account: username,
	}
	user, db := userInfo.FindUser()
	if db.Error != nil {
		c.Data["json"] = global.ResponseData{
			Code:    -1,
			Message: "用户信息查询失败",
		}
		c.ServeJSON(true)
	}
	if user.Password != password {
		c.Data["json"] = global.ResponseData{
			Code:    -1,
			Message: "密码错误",
		}
		c.ServeJSON(true)
	}

	c.Data["json"] = global.ResponseData{
		Code:    0,
		Message: "登录成功",
	}
	c.ServeJSON(true)
	return
}

// 获取好友（联系人列表）
func (c *UserController) FriendList() {
	userId, _ := c.GetInt("user_id")

	if userId == 0 {
		c.Data["json"] = global.ResponseData{
			Code:    -1,
			Message: "登陆状态异常",
		}
		c.ServeJSON(true)
	}

	ur := models.UserRelationship{OwnerId: uint(userId)}
	// 获取用户的好友id
	list, res := ur.GetUserDistId()
	if res.Error != nil {
		c.Data["json"] = global.ResponseData{
			Code:    -1,
			Message: "用户数据获取失败",
		}
		c.ServeJSON(true)
	}

	// 使用并发编程来实现用户关系的查询
	distIdList := make(chan uint, res.RowsAffected)
	// 用户信息数据结构
	userInfo := make(map[uint]models.User)

	// 将当前登录用户的所有好友id放入通道
	global.WG.Add(1)
	go func() {
		defer global.WG.Done()
		for _, v := range list {
			distIdList <- v.DistId
		}
		close(distIdList)
	}()
	global.WG.Wait()

	// 根据每个好友id获取对应好友的用户信息
	for data := range distIdList {
		u := models.User{Model: gorm.Model{ID: data}}
		res = u.GetUserInfoById()
		if res.Error != nil {
			break
		}
		userInfo[data] = u
	}
	// 好友列表
	var friendList []models.FriendList

	// 拼凑
	for _, v := range list {
		friendList = append(friendList, models.FriendList{
			UserRelationship: v,
			UserInfo:         userInfo[v.DistId],
		})
	}

	c.Data["json"] = global.ResponseData{
		Code:    0,
		Message: "好友列表获取成功",
		Data: map[string]interface{}{
			"friend_list": friendList,
		},
	}
	c.ServeJSON(true)
}
