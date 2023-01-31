package formatter

import "github.com/labstack/echo/v4"

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseOKWithData struct {
	Response
	Data interface{} `json:"data"`
}

type ResponseErrorWithData struct {
	Response
	Data interface{} `json:"data"`
}

func SuccessWithDataJSON(ctx echo.Context, statusCode int, msg string, data interface{}) error {
	r := &ResponseOKWithData{
		Response: Response{
			Status:  "ok",
			Message: msg,
		},
		Data: data,
	}
	return ctx.JSON(statusCode, r)
}

func SuccessJSON(ctx echo.Context, statusCode int, msg string) error {
	r := &Response{
		Status:  "ok",
		Message: msg,
	}
	return ctx.JSON(statusCode, r)
}

func ErrorJSON(ctx echo.Context, statusCode int, msg string) error {
	r := &Response{
		Status:  "error",
		Message: msg,
	}
	return ctx.JSON(statusCode, r)
}

func ErrorWithDataJSON(ctx echo.Context, statusCode int, msg string, data interface{}) error {
	r := &ResponseErrorWithData{
		Response: Response{
			Status:  "error",
			Message: msg,
		},
		Data: data,
	}
	return ctx.JSON(statusCode, r)
}
