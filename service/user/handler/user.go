package handler

import (
	"context"
	"my-micro/infra/imitate/vsm"
	pb "my-micro/service/user/proto"
	"my-micro/web/model"
	"my-micro/web/utils"
)

type User struct{}

// Return a new handler
func New() *User {
	return &User{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *User) SendSms(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	verify := model.CheckImgCode(req.Uuid, req.ImgCode)
	if !verify {
		rsp.Errno = utils.CHECK_FAILD
		rsp.Errmsg = utils.RecodeText(utils.CHECK_FAILD)
		return nil
	}
	phone := req.Phone
	code, _ := vsm.GenVerifyCode(phone)
	model.SaveSmsCode(phone, code)
	rsp.Errno = utils.SUCCESS
	rsp.Errmsg = utils.RecodeText(utils.SUCCESS)
	return nil
}

func (e *User) Register(ctx context.Context, req *pb.RegisterRequest, rsp *pb.RegisterResponse) error {
	// 校验验证码是否正确
	verify := model.CheckSmsCode(req.GetMobile(), req.GetSmsCode())
	if !verify {
		rsp.Errno = utils.CHECK_FAILD
		rsp.Errmsg = utils.RecodeText(utils.CHECK_FAILD)
		return nil
	}
	err := model.RegisterUser(req.Mobile, req.Password)
	if err != nil {
		rsp.Errno = utils.SYSTEM_ERROR
		rsp.Errmsg = utils.RecodeText(utils.SYSTEM_ERROR)
		return err
	}
	rsp.Errno = utils.SUCCESS
	rsp.Errmsg = utils.RecodeText(utils.SUCCESS)
	return nil
}
