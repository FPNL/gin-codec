<div align="center">
	<h1>gin-codec</h1>
	<p>
		<b>Instead of Bind and Render, call Decode and encode for custom case</b>
	</p>
	<br>
	<br>
	<br>
</div>

# install

```bash
go get github.com/FPNL/codec
```

# usage example

```go
package main

import (
	"github.com/FPNL/codec"
	"github.com/gin-gonic/gin"
)

func handler(coder codec.Codec) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var in any
		var out any

		// decode
		if err := coder.Bind(ctx, &in); err != nil {
			_ = ctx.Error(err)
			return
		}

		// decode
		if err := coder.BindQuery(ctx, &in); err != nil {
			_ = ctx.Error(err)
			return
		}

		// decode
		if err := coder.BindVars(ctx, &in); err != nil {
			_ = ctx.Error(err)
			return
		}

		// Do what you want to in

		// encode
		err := coder.Result(ctx, out)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
	}
}
```
