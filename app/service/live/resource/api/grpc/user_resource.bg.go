package v1

import (
	"fmt"
	"net/http"
	"context"
	"github.com/gin-gonic/gin"
)

var _ context.Context

type UserResource interface {
	// Query请求单个资源
	Query(ctx context.Context, req *QueryReq) (resp *QueryResp, err error)
}

func userResourceQuery(c *gin.Context) {
	p := new(QueryReq)

	if err := c.ShouldBind(p); err != nil {
		return
	}
	
	resp, _ := v1UserResourceSvc.Query(c, p)
	fmt.Println("你好，获取单个用户")
	c.JSON(http.StatusOK, resp)
}

var v1UserResourceSvc UserResource

func RegisterV1UserResourceService(e *gin.Engine, svc UserResource) {
	v1UserResourceSvc = svc
	e.GET("xlive/resource/v1/userResource/Query", userResourceQuery)
}