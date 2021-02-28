package model

import (
	"container/list"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net"
	"net/http"
	"time"
)

// HostAgentLabel 定义了agent的属性，例如master/slave
type HostAgentLabel int

// HostAgentStatus 定义了Agent的状态，online/offline
type HostAgentStatus int

// HostAgentModel 定义了主机上Agent的结构，同时进行数据的后端校验
type HostAgentModel struct {
	Label      HostAgentLabel
	Address    string `form:"ipaddress" binding:"required,validIP"`
	Name       string
	Status     HostAgentStatus
	Date       time.Time
	DockerList list.List
	ContainerList list.List


}
 
// CheckIPValidation 用作验证IP有效性，自定义验证器，和结构体中的Address结构体绑定
var CheckIPValidation validator.Func = func(fl validator.FieldLevel) bool {
	// @IMPROVE 改成ping通ping不通？
	ipaddr, ok := fl.Field().Interface().(string)
	if ok {
		return net.ParseIP(ipaddr) != nil
	}
	return false
}

func main() {
	route := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("validIP", CheckIPValidation)
	}

	route.GET("/agent", getAgentMsg)
	route.Run(":8085")
}

func getAgentMsg(c *gin.Context) {
	var b HostAgentModel
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Parsing IP successfully"})
		// 返回结果的逻辑
		// @TODO
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
