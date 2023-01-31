package api

import (
	"fmt"
	"go.uber.org/zap"
	"strconv"
	"wobbs-server/vo"

	"github.com/gin-gonic/gin"

	"wobbs-server/common"
	"wobbs-server/config"
	"wobbs-server/dto"
	"wobbs-server/logic"
)

func GetPostList(ctx *gin.Context) {
	query1 := ctx.DefaultQuery("page", strconv.Itoa(1))
	query2 := ctx.DefaultQuery("page_size", strconv.Itoa(10))
	page, _ := strconv.ParseInt(query1, 10, 32)
	pageSize, _ := strconv.ParseInt(query2, 10, 32)
	postList := logic.GetPostList(int(page), int(pageSize))
	common.Success(ctx, postList)
}

func GetPostDetail(ctx *gin.Context) {
	pidStr := ctx.Param("post_id")
	if pidStr == "" {
		zap.L().Error("post_id")
		common.FailByMsg(ctx, "post_id为空")
		return
	}
	pid, err := strconv.ParseInt(pidStr, 10, 32)
	if err != nil {

	}
	detail := logic.GetPostDetail(int32(pid))
	category := logic.GetCategoryById(detail.CategoryID)
	author := logic.GetUserById(detail.AuthorID)
	common.Success(ctx,
		vo.PostDetail{AuthorName: author.Username,
			CategoryName: category.Name,
			Post:         detail})
}

func CreatePost(ctx *gin.Context) {
	var postDTO dto.PostDTO
	if err := ctx.ShouldBind(&postDTO); err != nil {
		config.ValidateError(ctx, err)
		return
	}
	userId, exists := ctx.Get("userId")
	if !exists {
		fmt.Println("未登录")
	}
	logic.CreatePost(userId.(int64), postDTO)
	common.Success(ctx, nil)
}

func PostVoting(ctx *gin.Context) {
	var voteDTO dto.VoteDTO
	if err := ctx.ShouldBind(&voteDTO); err != nil {
		config.ValidateError(ctx, err)
		return
	}
	userId, _ := ctx.Get("userId")
	logic.PostVoting(userId.(int64), voteDTO)
	common.Success(ctx, nil)
}
