package controller

import (
	"net/http"
	"restful_api/helper"
	"restful_api/service"
	"restful_api/web"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpln struct {
	Service service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpln{
		Service: categoryService,
	}
}
func (controller *CategoryControllerImpln) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	CategoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &CategoryCreateRequest)

	CategoryResponse := controller.Service.Create(request.Context(), CategoryCreateRequest)

	WebResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   CategoryResponse,
	}

	helper.WriteToResponseBody(writer, WebResponse)

}
func (controller *CategoryControllerImpln) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	CategoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &CategoryUpdateRequest)

	CategoryId := params.ByName("categoryId")

	id, errid := strconv.Atoi(CategoryId)
	helper.ErrorT(errid)

	CategoryUpdateRequest.Id = id
	CategoryResponse := controller.Service.Update(request.Context(), CategoryUpdateRequest)

	WebResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   CategoryResponse,
	}
	helper.WriteToResponseBody(writer, WebResponse)

}
func (controller *CategoryControllerImpln) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	CategoryId := params.ByName("categoryId")

	id, err := strconv.Atoi(CategoryId)
	helper.ErrorT(err)

	controller.Service.Delete(request.Context(), id)

	WebResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, WebResponse)

}
func (controller *CategoryControllerImpln) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	CategoryId := params.ByName("categoryId")

	id, err := strconv.Atoi(CategoryId)
	helper.ErrorT(err)

	ResponseWeb := controller.Service.FindById(request.Context(), id)
	WebResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ResponseWeb,
	}
	helper.WriteToResponseBody(writer, WebResponse)

}
func (controller *CategoryControllerImpln) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	CategoryResponses := controller.Service.FindAll(request.Context())

	WebResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   CategoryResponses,
	}
	helper.WriteToResponseBody(writer, WebResponse)

}
