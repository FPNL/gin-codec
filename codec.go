package codec

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var DefaultCodec Codec = &codec{}

type Codec interface {
	Bind(*gin.Context, any) error
	BindVars(*gin.Context, any) error
	BindQuery(*gin.Context, any) error
	BindForm(*gin.Context, any) error
	Result(*gin.Context, any) error
}

type codec struct{}

func (c *codec) Bind(ctx *gin.Context, in any) error {
	return ctx.ShouldBind(in)
}

func (c *codec) BindVars(ctx *gin.Context, in any) error {
	return ctx.ShouldBindUri(in)
}

func (c *codec) BindQuery(ctx *gin.Context, in any) error {
	return ctx.BindQuery(in)
}

func (c *codec) BindForm(ctx *gin.Context, in any) error {
	return ctx.ShouldBindWith(in, binding.Form)
}

// Result only supports JSON, XML, and text/plain, default is JSON.
func (c *codec) Result(ctx *gin.Context, out any) error {
	switch ctx.ContentType() {
	case "application/xml":
		ctx.XML(200, out)
	case "application/json":
		ctx.JSON(200, out)
	case "text/plain":
		s, ok := out.(string)
		if !ok {
			return fmt.Errorf("invalid out for text/plain")
		}
		ctx.String(200, s)
	default:
		ctx.JSON(200, out)
	}

	return nil
}
