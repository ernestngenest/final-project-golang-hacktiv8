package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"final_project_hacktiv8/helpers"
	"final_project_hacktiv8/modules/comments/dto"
	"final_project_hacktiv8/modules/comments/service"
)

type ControllerComment interface {
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type controller struct {
	srv service.ServiceComment
}

func New(srv service.ServiceComment) ControllerComment {
	return &controller{srv: srv}
}

// Create a comment
// @Tags         comments
// @Summary      Create a comment
// @Description  Create a comment
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string                                                     true  "Bearer + user token"
// @Param        data           body      dto.Request                                                true  "data"
// @Success      201            {object}  helpers.BaseResponse{data=dto.ResponseInsert}              "CREATED"
// @Failure      400            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Bad Request"
// @Failure      401            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Unauthorization"
// @Router       /comments [POST]
func (c *controller) Create(ctx *gin.Context) {
	request := new(dto.Request)
	err := ctx.ShouldBind(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	userID := ctx.MustGet("user_id")
	request.UserID = userID.(uint)

	// service
	create, err := c.srv.Create(*request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(helpers.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusCreated, helpers.NewResponse(http.StatusCreated, create, nil))
}

// Get comments
// @Tags         comments
// @Summary      Get comments
// @Description  Get comments
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string                                                     true  "Bearer + user token"
// @Success      200            {object}  helpers.BaseResponse{data=[]dto.Response}                  "OK"
// @Failure      400            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Bad Request"
// @Failure      401            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Unauthorization"
// @Router       /comments [GET]
func (c *controller) Get(ctx *gin.Context) {
	responses, err := c.srv.Get()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(helpers.GetStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helpers.NewResponse(http.StatusOK, responses, nil))
}

// Update a comment
// @Tags         comments
// @Summary      Update a comment
// @Description  Update a comment
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string                                                     true  "Bearer + user token"
// @Param        commentID      path      int                                                        true  "ID of the comment"
// @Param        data           body      dto.RequestUpdate                                          true  "data"
// @Success      200            {object}  helpers.BaseResponse{data=dto.ResponseUpdate}              "OK"
// @Failure      400            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Bad Request"
// @Failure      401            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Unauthorization"
// @Router       /comments/:commentID [PUT]
func (c *controller) Update(ctx *gin.Context) {
	paramKeyID := ctx.Param("commentID")
	commentID, err := strconv.Atoi(paramKeyID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, err))
		return
	}

	request := new(dto.RequestUpdate)

	userID := ctx.MustGet("user_id")
	request.UserID = userID.(uint)

	err = ctx.ShouldBind(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, err))
		return
	}
	update, err := c.srv.Update(*request, uint(commentID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, err))
		return
	}

	ctx.JSON(http.StatusAccepted, helpers.NewResponse(http.StatusAccepted, update, nil))
}

// Delete a comment
// @Tags         comments
// @Summary      Delete a comment
// @Description  Delete a comment
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string                                                     true  "Bearer + user token"
// @Param        commentID      path      int                                                        true  "ID of the comment"
// @Success      200            {object}  helpers.BaseResponse                                       "OK"
// @Failure      400            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Bad Request"
// @Failure      401            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Unauthorization"
// @Router       /comments/:commentID [DELETE]
func (c *controller) Delete(ctx *gin.Context) {
	paramKeyID := ctx.Param("commentID")
	commentID, err := strconv.Atoi(paramKeyID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, err))
		return
	}

	err = c.srv.Delete(uint(commentID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helpers.NewResponse(http.StatusOK, "Your comment successfully deleted", nil))
}
