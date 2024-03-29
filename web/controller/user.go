package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"image/png"
	"my-micro/infra/imitate/vdb"
	user "my-micro/service/user/proto"
	"my-micro/web/model"
	"my-micro/web/proto/getCaptcha"
	"my-micro/web/utils"
	"net/http"

	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
)

func GetSession(c *gin.Context) {
	resp := make(map[string]string)
	resp["errno"] = utils.SYSTEM_ERROR
	resp["errmsg"] = utils.RecodeText(utils.SYSTEM_ERROR)
	c.JSON(http.StatusOK, resp)
}

func GetImageCd(c *gin.Context) {
	// 指定服务发现
	consulClient := utils.InitMicro()
	// 初始化客户端
	microClient := getCaptcha.NewGetCaptchaService("go.micro.srv.getCaptcha", consulClient.Client())
	// 调用服务端接口
	uuid := c.Param("uuid")
	response, err := microClient.Call(context.TODO(), &getCaptcha.Request{Uuid: uuid})
	if err != nil {
		fmt.Printf("%v", err)
	}
	// 反序列化字节流，变为 img
	var img captcha.Image
	err = json.Unmarshal(response.B, &img)
	// 写浏览器数据
	png.Encode(c.Writer, img)
}

// https://localhost:8080//api/v1.0/smscode/13218001299?imageCode=St442C&uuid=fk36osfdiijoty34454435
func GetSmsCd(c *gin.Context) {
	consulClient := utils.InitMicro()
	microClient := user.NewUserService("go.micro.srv.user", consulClient.Client())

	resp, err := microClient.SendSms(context.TODO(), &user.Request{
		Uuid:    c.Query("uuid"),
		ImgCode: c.Query("imageCode"),
		Phone:   c.Param("phone"),
	})
	if err != nil {
		fmt.Printf("%v", err)
	}
	c.JSON(http.StatusOK, resp)
}

// 注册
func PostRet(c *gin.Context) {

	// 定义匿名结构体，接受数据
	// c.PostForm("mobile") 这样接收不到，这样只能接收 Form 表单的数据
	var req struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
		SmsCode  string `json:"sms_code"`
	}
	c.Bind(&req)

	consulClient := utils.InitMicro()
	microClient := user.NewUserService("go.micro.srv.user", consulClient.Client())
	registerResponse, err := microClient.Register(context.TODO(), &user.RegisterRequest{
		Mobile:   req.Mobile,
		Password: req.Password,
		SmsCode:  req.SmsCode,
	})
	if err != nil {
		fmt.Printf("%v", err)
	}
	c.JSON(http.StatusOK, registerResponse)
}

// 获取地域信息
func GetArea(c *gin.Context) {
	var areas []model.Area

	data := vdb.GetString("areas")

	if data != "" {
		json.Unmarshal([]byte(data), &areas)
	} else {
		model.GlobalConn.Select("*").Find(&areas)
		marshal, _ := json.Marshal(areas)
		vdb.SetString("areas", string(marshal))
	}
	c.JSON(http.StatusOK, areas)
}
