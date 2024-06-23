package common

import (
	"admin/code"
	"admin/proto/code_proto"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestLogicErrorHandle(t *testing.T) {
	ctx := &gin.Context{}
	err := code.NewCode(code_proto.ErrorCode_Error)
	result := code.NewCode(code_proto.ErrorCode_Success)
	HandleLogicError(ctx, err, "水电费是的", result)
}
