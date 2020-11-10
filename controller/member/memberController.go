package member

import (
	"apiDemoProject/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 获取会员详情
// @Tags 会员
// @Accept  json
// @Produce  json
// @Param member body model.Member true "会员"
// @Param token header string false "Token"
// @Param x-username header string false "用户名"
// @Success 200 {object} model.Response 成功后返回值
// @Router /api/member/info/get [post]
func GetInfo(c *gin.Context) {
	data := make(map[string]interface{})
	data["user_info"] = model.Member{UserName: "chen", NickName: "chen", Age: 1, Sex: 1}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

// @Summary 获取所有会员
// @Tags 会员
// @Accept  json
// @Produce  json
// @Param token header string false "Token"
// @Param x-username header string false "用户名"
// @Success 200 {object} model.Response 成功后返回值
// @Router /api/member/list/get [post]
func GetAll(c *gin.Context) {
	data := make(map[string]interface{})
	data["result"] = "OK"
	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

// @Summary 上传会员
// @Tags 会员
// @Accept  json
// @Produce  json
// @Param member body model.Member true "会员"
// @Param token header string false "Token"
// @Param x-username header string false "用户名"
// @Success 200 {object} model.Response 成功后返回值
// @Router /api/member/upload/post [post]
func Upload(c *gin.Context) {
	data := make(map[string]interface{})
	data["result"] = "OK"
	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

// @Summary 更新会员
// @Tags 会员
// @Accept  json
// @Produce  json
// @Param member body model.Member true "会员"
// @Param token header string false "Token"
// @Param x-username header string false "用户名"
// @Success 200 {object} model.Response 成功后返回值
// @Router /api/member/update/post [post]
func Update(c *gin.Context) {
	data := make(map[string]interface{})
	data["result"] = "OK"
	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

// @Summary 删除会员
// @Tags 会员
// @Accept  json
// @Produce  json
// @Param token header string false "Token"
// @Param x-username header string false "用户名"
// @Success 200  {object} model.Response 成功后返回值
// @Router /api/member/delete/post [post]
func Delete(c *gin.Context) {
	data := make(map[string]interface{})
	data["result"] = "OK"
	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}
