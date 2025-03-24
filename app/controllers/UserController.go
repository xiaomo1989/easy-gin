package controllers

import (
	"easy-gin/app/models"
	"easy-gin/configs"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"strconv"
	"time"
)

// get one
func UserGet(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	userModel := models.User{}
	data, err := userModel.UserGet(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// get list
func UserGetList(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	pageSize := ctx.DefaultQuery("page_size", "20")
	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)
	userModel := models.User{}
	users, err := userModel.UserGetList(pageInt, pageSizeInt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}

	err1 := configs.GetRedis().Set("name", "Gin+Redis", 10*time.Second)
	if err1 != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		//return
	}

	value, err2 := configs.GetRedis().Get("name")
	if err2 == redis.Nil {
		ctx.JSON(404, gin.H{"message": "数据不存在"})
		//return
	} else if err1 != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		//return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": users,
		"res":  value,
	})
}

// add one
func UserPost(ctx *gin.Context) {
	userModel := models.User{}
	if err := ctx.ShouldBind(&userModel); nil != err {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	id, err := userModel.UserAdd()

	if nil != err {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg": "success",
		"uid": id,
	})
}

// update
func UserPut(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)

	if nil != err || 0 == idInt {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "resource identifier not found",
		})
		return
	}

	userModel := models.User{}

	if err := ctx.ShouldBind(&userModel); nil != err {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	_, err = userModel.UserUpdate(idInt)

	if nil != err {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 更新成功返回 204
	ctx.JSON(http.StatusNoContent, gin.H{})
}

// delete
func UserDelete(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)

	if nil != err || 0 == idInt {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "resource identifier not found",
		})
		return
	}

	userModel := models.User{}

	_, err = userModel.UserDelete(idInt)

	if nil != err {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 删除成功返回 204
	ctx.JSON(http.StatusNoContent, gin.H{})
}
