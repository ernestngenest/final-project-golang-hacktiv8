package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"final_project_hacktiv8/helpers"
	"final_project_hacktiv8/modules/photos/dto"
	"final_project_hacktiv8/modules/photos/service"
)

type ControllerPhoto interface {
	Create(ctx *gin.Context)
	GetPhotos(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type controller struct {
	srv service.ServicePhoto
}

func New(srv service.ServicePhoto) ControllerPhoto {
	return &controller{srv: srv}
}

// Create new photo
// @Tags         photos
// @Summary      Create new photo
// @Description  Create new photo
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string                                                     true  "Bearer + user token"
// @Param        data           body      dto.Request                                                true  "data"
// @Success      201            {object}  helpers.BaseResponse{data=dto.Response}                    "CREATED"
// @Failure      400            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Bad Request"
// @Failure      401            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Unauthorization"
// @Router       /photos [POST]
func (c *controller) Create(ctx *gin.Context) {
	data := new(dto.Request)

	err := ctx.ShouldBind(data) // bind request body to data
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	// must have user id from bearer
	userID := ctx.MustGet("user_id")
	data.UserID = userID.(uint)

	response, err := c.srv.Create(*data)
	if err != nil {
		ctx.JSON(helpers.GetStatusCode(err), helpers.NewResponse(helpers.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusCreated, helpers.NewResponse(http.StatusCreated, response, nil))
}

// GetPhotos a photo
// @Tags         photos
// @Summary      GetPhotos a photo
// @Description  GetPhotos a photo
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string                                                     true  "Bearer + user token"
// @Success      200            {object}  helpers.BaseResponse{data=[]dto.ResponseGet}               "SUCCESS"
// @Failure      400            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Bad Request"
// @Failure      401            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Unauthorization"
// @Failure      404            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Not Found"
// @Router       /photos [GET]
func (c *controller) GetPhotos(ctx *gin.Context) {
	response, err := c.srv.GetPhotos()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helpers.NewResponse(http.StatusOK, response, nil))
}

// Update a photo by ID
// @Tags         photos
// @Summary      Update a photo by ID
// @Description  Update a photo by ID
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string                                                     true  "Bearer + user token"
// @Param        id             path      string                                                     true  "photo ID"
// @Param        data           body      dto.Request                                                true  "data"
// @Success      200            {object}  helpers.BaseResponse{data=dto.ResponseUpdate}              "SUCCESS"
// @Failure      400            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Bad Request"
// @Failure      401            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Unauthorization"
// @Failure      404            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Not Found"
// @Router       /photos/{id} [PUT]
func (c *controller) Update(ctx *gin.Context) {
	data := new(dto.Request)

	err := ctx.ShouldBind(data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}
	photoParamID := ctx.Param("photoID")
	photoID, err := strconv.Atoi(photoParamID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(http.StatusBadRequest, nil, err))
		return
	}

	// must have user id from bearer
	userID := ctx.MustGet("user_id")
	data.UserID = userID.(uint)

	update, err := c.srv.Update(*data, photoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.NewResponse(helpers.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helpers.NewResponse(http.StatusOK, update, nil))
}

// Delete a photo
// @Tags         photos
// @Summary      Delete a photo
// @Description  Delete a photo
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string                                                     true  "Bearer + user token"
// @Param        photoID        path      int                                                        true  "ID of the photo"
// @Success      200            {object}  helpers.BaseResponse                                       "SUCCESS"
// @Failure      400            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Bad Request"
// @Failure      401            {object}  helpers.BaseResponse{errors=helpers.ExampleErrorResponse}  "Unauthorization"
// @Router       /photos/:photoID [DELETE]
func (c *controller) Delete(ctx *gin.Context) {
	paramKeyID := ctx.Param("photoID")
	photoID, _ := strconv.Atoi(paramKeyID)
	err := c.srv.Delete(photoID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, helpers.NewResponse(helpers.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helpers.NewResponse(http.StatusOK, "Your Photo has been successfully deleted", nil))
}
