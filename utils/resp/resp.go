package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseHelper struct{}

var RespHelper = &ResponseHelper{}

type JsonStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

type JsonErrorStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

func (r *ResponseHelper) OK(c *gin.Context, data interface{}) {
	json := &JsonStruct{
		Code: Success.Code,
		Msg:  Success.Message,
		Data: data,
	}
	c.JSON(http.StatusOK, json)
}

func (r *ResponseHelper) OkWithMessage(c *gin.Context, data interface{}, code int, msg interface{}) {
	json := &JsonStruct{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	c.JSON(http.StatusOK, json)
}

func (r *ResponseHelper) Fail(c *gin.Context, code int, msg interface{}) {
	json := &JsonErrorStruct{Code: code, Msg: msg}
	c.JSON(http.StatusOK, json)
}
