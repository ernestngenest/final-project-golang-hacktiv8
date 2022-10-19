package controller

import (
	"final_project_hacktiv8/helpers"
	"final_project_hacktiv8/modules/socialMedias/dto"
	"final_project_hacktiv8/modules/socialMedias/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ControllerSocialMedia interface {
	Create(ctx *gin.Context)
	GetList(ctx *gin.Context)
	UpdateByID(ctx *gin.Context)
	DeleteByID(ctx *gin.Context)
}

type controller struct {
	srv service.ServiceSocialMedia
}

func New(srv service.ServiceSocialMedia) ControllerSocialMedia {
	return &controller{srv}
}

// Create new social media
// @Tags         socialmedias
// @Summary      Create new social media
// @Description  Create social media
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string                                                     true  "Bearer + user token"
// @Param        data           body      dto.Request                                                true  "data"
// @Success      201            {object}  helpers.BaseResponse{data=dto.Response}                    "CREATED"
// @Failure      400            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Bad Request"
// @Router       /socialmedias [POST]
func (c *controller) Create(ctx *gin.Context) {
	data := new(dto.Request)

	if err := ctx.ShouldBind(data); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helpers.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	data.UserID = ctx.MustGet("user_id").(uint)

	response, err := c.srv.Create(*data)
	if err != nil {
		ctx.JSON(helpers.GetStatusCode(err), helpers.NewResponse(helpers.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusCreated, helpers.NewResponse(http.StatusCreated, response, nil))
}

// Get all social media
// @Tags         socialmedias
// @Summary      Get all social media
// @Description  Get all social media
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string                                              true  "Bearer + user token"
// @Success      200            {object}  helpers.BaseResponse{data=dto.ResponseListWrapper}  "SUCCESS"
// @Router       /socialmedias [GET]
func (c *controller) GetList(ctx *gin.Context) {

	response, err := c.srv.GetList()
	if err != nil {
		ctx.JSON(helpers.GetStatusCode(err), helpers.NewResponse(helpers.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helpers.NewResponse(http.StatusOK, response, nil))
}

// Update by id social media
// @Tags         socialmedias
// @Summary      Update by id social media
// @Description  Update by id social media
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string                                                     true  "Bearer + user token"
// @Param        socialmediaid  path      int                                                        true  "ID of the social media"
// @Param        data           body      dto.Request                                                true  "data"
// @Success      200            {object}  helpers.BaseResponse{data=dto.Response}                    "SUCCESS"
// @Failure      400            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Bad Request"
// @Failure      404            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Record not found"
// @Router       /socialmedias/:socialmediaid [PUT]
func (c *controller) UpdateByID(ctx *gin.Context) {
	data := new(dto.Request)

	if err := ctx.ShouldBind(data); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helpers.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	idString := ctx.Param("socialmediaid")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, err))
		return
	}
	data.ID = uint(id)

	response, err := c.srv.UpdateByID(*data)
	if err != nil {
		ctx.JSON(helpers.GetStatusCode(err), helpers.NewResponse(helpers.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helpers.NewResponse(http.StatusOK, response, nil))
}

// Delete by id social media
// @Tags         socialmedias
// @Summary      Delete by id social media
// @Description  Delete by id social media
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string                                                     true  "Bearer + user token"
// @Param        socialmediaid  path      int                                                        true  "ID of the social media"
// @Success      200            {object}  helpers.BaseResponse{data=dto.Response}       "SUCCESS"
// @Failure      400            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Bad Request"
// @Failure      404            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Record not found"
// @Router       /socialmedias/:socialmediaid [DELETE]
func (c *controller) DeleteByID(ctx *gin.Context) {
	idString := ctx.Param("socialmediaid")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, err))
		return
	}

	err = c.srv.DeleteByID(uint(id))
	if err != nil {
		ctx.JSON(helpers.GetStatusCode(err), helpers.NewResponse(helpers.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helpers.NewResponse(http.StatusOK, map[string]interface{}{"message": "Your social media has been successfully deleted"}, nil))
}
