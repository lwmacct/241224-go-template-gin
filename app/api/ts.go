package api

import (
	"github.com/gin-gonic/gin"
	// _ "github.com/lwmacct/241224-go-template-gin/app/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Ts struct{}

func New() *Ts {
	return &Ts{}
}

func (ts *Ts) Init() {
}

func (ts *Ts) Run() {
	r := gin.Default()

	r.POST("/demo", ts.DemoHandler)

	// Swagger UI 路由（可自定义路径）
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")

}

// ExampleReq 请求示例的结构体
type ExampleReq struct {
	// 注意 example 里的值需要用双引号括起来
	Name string `json:"name" example:"Tom"`
	Age  int    `json:"age" example:"18"`
}

// ExampleResp 返回示例的结构体
type ExampleResp struct {
	Code int    `json:"code" example:"0"`
	Msg  string `json:"msg"  example:"success"`
}

// @Summary      测试接口（JSON示例）
// @Description  这是一个接收 ExampleReq 并返回 ExampleResp 的示例
// @Tags         demo
// @Accept       json
// @Produce      json
// @Param        data  body      ExampleReq  true  "请求示例"
// @Success      200   {object}  ExampleResp "返回示例"
// @Router       /demo [post]
func (ts *Ts) DemoHandler(c *gin.Context) {
	var req ExampleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// 这里省略具体处理...
	c.JSON(200, ExampleResp{
		Code: 0,
		Msg:  "Hello " + req.Name,
	})
}
