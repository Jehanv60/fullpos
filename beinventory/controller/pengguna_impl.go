package controller

import (
	"net/http"
	"strconv"

	"github.com/Jehanv60/helper"
	"github.com/Jehanv60/model/web"
	"github.com/Jehanv60/service"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
	Validate    *validator.Validate
}

func NewPenggunaController(UserService service.UserService, validate *validator.Validate) PenggunaController {
	return &UserControllerImpl{
		UserService: UserService,
		Validate:    validate,
	}
}

// Create implements PenggunaController.
func (controller *UserControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	penggunaCreateRequest := web.PenggunaCreateRequest{}
	helper.ReadFromBody(r, &penggunaCreateRequest)
	controller.UserService.Create(r.Context(), penggunaCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   "Data Berhasil Ditambahkan",
	}
	helper.WriteToResponse(w, webResponse)
}

// FindAll implements PenggunaController.
func (controller *UserControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	penggunaResponse := controller.UserService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   penggunaResponse,
	}
	helper.WriteToResponse(w, webResponse)
}

// FindById implements PenggunaController.
func (controller *UserControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("penggunaId"))
	helper.PanicError(err)
	penggunaResponse := controller.UserService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   penggunaResponse,
	}
	helper.WriteToResponse(w, webResponse)
}

// Update implements PenggunaController.
func (controller *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	penggunaUpdate := web.PenggunaUpdate{}
	helper.ReadFromBody(r, &penggunaUpdate)
	id, err := strconv.Atoi(params.ByName("penggunaId"))
	helper.PanicError(err)
	penggunaUpdate.Id = id
	controller.UserService.Update(r.Context(), penggunaUpdate)
	helper.PanicError(err)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   "Data Berhasil Diupdate",
	}
	helper.WriteToResponse(w, webResponse)
}
