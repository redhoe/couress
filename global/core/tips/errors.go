package tips

import (
	"fmt"
)

type AppError struct {
	errCode int
	errMsg  string
	data    map[string]interface{}
	isClone bool
}

func NewAppError(code int, msg string) *AppError {
	return &AppError{
		errCode: code,
		errMsg:  msg,
		data:    nil,
		isClone: false,
	}
}

func (ae *AppError) GetCode() int {
	return ae.errCode
}

func (ae *AppError) String() string {
	return fmt.Sprintf("errCode:%d,errMsg:%s,data:%v", ae.errCode, ae.errMsg, ae.data)
}

func (ae *AppError) SetMessage(msg string) *AppError {
	if ae.isClone == false {
		ae = ae.clone()
	}
	ae.errMsg = msg
	return ae
}

func (ae *AppError) GetMessage(lang string, args ...any) string {
	return getLang(ae.errMsg, lang, args)
}

func (ae *AppError) SetMapData(data map[string]interface{}) *AppError {
	if ae.isClone == false {
		ae = ae.clone()
	}
	ae.data = data
	return ae
}

func (ae *AppError) SetAny(data any) *AppError {
	if ae.isClone == false {
		ae = ae.clone()
	}
	ae.data = map[string]interface{}{
		"data": data,
	}
	return ae
}

func (ae *AppError) GetData() map[string]interface{} {
	if ae.data == nil {
		ae.data = map[string]interface{}{}
	}
	return ae.data
}

func (ae *AppError) MapData() map[string]any {
	return map[string]any{
		"errCode": ae.GetCode(),
		"errMsg":  ae.GetMessage(DefaultLang.String()),
		"data":    nil,
	}
}

func (ae *AppError) clone() *AppError {
	return newApiErrorClone(ae.errCode, ae.errMsg)
}

func newApiErrorClone(code int, msg string) *AppError {
	return &AppError{
		errCode: code,
		errMsg:  msg,
		data:    nil,
		isClone: true,
	}
}
