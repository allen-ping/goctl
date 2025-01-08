package response

import (
	"net/http"

	errorx "github.com/go-kratos/kratos/v2/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		ex := new(errorx.Error)
		if errorx.As(err, &ex) {
			e := errorx.FromError(err)
			body.Code = int(e.Code)
			body.Message = e.Message
		} else {
			body.Code = errorx.UnknownCode
			body.Message = err.Error()
		}
	} else {
		body.Code = 0
		body.Message = "OK"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
