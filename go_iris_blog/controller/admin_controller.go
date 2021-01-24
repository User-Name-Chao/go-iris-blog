package controller

import (
	"TestProject/service"
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

/**
 * 管理员控制器
 */
type AdminController struct {
	//iris框架自动为每个请求都绑定上下文对象：可作为接受参数
	Ctx iris.Context

	//admin功能实体：引入Service接口
	Service service.AdminService

	//session对象：存储session信息
	Session *sessions.Session
}

const (
	ADMIN = "admin"   //管理员登录成功后存储的session信息的key
)

//将发送请求的字段映射为指定字段
type AdminLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

/**
 * 管理员登录功能：json请求格式
 * 接口：/admin/login
 */
func (ac *AdminController) PostLogin(context iris.Context) mvc.Result {

	var adminLogin AdminLogin
	ac.Ctx.ReadJSON(&adminLogin)  //自动将请求的json字符串映射为AdminLogin结构体

	//数据参数检验
	if adminLogin.UserName == "" || adminLogin.Password == "" {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  "0",
				"success": "登录失败",
				"message": "用户名或密码为空,请重新填写后尝试登录",
			},
		}
	}

	//根据用户名、密码到数据库中查询对应的管理信息
	admin, exist := ac.Service.GetByAdminNameAndPassword(adminLogin.UserName, adminLogin.Password)

	//管理员不存在
	if !exist {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  "1",
				"success": "登录失败",
				"message": "用户名或者密码错误,请重新登录",
			},
		}
	}

	//管理员存在 设置session
	userByte, _ := json.Marshal(admin)
	ac.Session.Set(ADMIN, userByte)

	return mvc.Response{
		Object: map[string]interface{}{
			"status":  "1",
			"success": "登录成功",
			"message": "管理员登录成功",
		},
	}
}