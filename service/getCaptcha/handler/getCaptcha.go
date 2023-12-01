package handler

import (
	"context"
	"encoding/json"
	"image/color"
	pb "my-micro/service/getCaptcha/proto/getCaptcha"
	"my-micro/web/model"

	"github.com/afocus/captcha"
)

type GetCaptcha struct{}

// Return a new handler
func New() *GetCaptcha {
	return &GetCaptcha{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetCaptcha) Call(ctx context.Context, req *pb.Request, rsp *pb.Response) error {

	cp := captcha.New()
	cp.SetFont("service/getCaptcha/conf/comic.ttf")
	cp.SetSize(128, 64)
	cp.SetDisturbance(captcha.MEDIUM)
	cp.SetFrontColor(color.RGBA{132, 63, 0, 53})
	cp.SetBkgColor(color.RGBA{42, 163, 110, 99})
	img, code := cp.Create(4, captcha.NUM)
	imgBuf, _ := json.Marshal(img)
	// 保存验证码
	model.SaveImgCode(req.Uuid, code)
	rsp.B = imgBuf
	return nil
}
