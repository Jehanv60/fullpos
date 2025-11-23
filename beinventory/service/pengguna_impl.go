package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Jehanv60/exception"
	"github.com/Jehanv60/helper"
	"github.com/Jehanv60/model/domain"
	"github.com/Jehanv60/model/web"
	"github.com/Jehanv60/repository"
	"github.com/Jehanv60/util"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type UserServiceImpl struct {
	PenggunaRepository repository.PenggunaRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewUserService(penggunaRepository repository.PenggunaRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		PenggunaRepository: penggunaRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

// Create implements UserService.
func (service *UserServiceImpl) Create(ctx context.Context, request web.PenggunaCreateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	util.ErrValidateSelf(err)
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	penggunass := service.PenggunaRepository.FindByPenggunaRegister(ctx, tx, request.Pengguna, request.Email)
	if penggunass.Pengguna == request.Pengguna {
		panic(exception.NewSameFound(fmt.Sprintf("%s Sudah Digunakan", request.Pengguna)))
	}
	if penggunass.Email == request.Email {
		panic(exception.NewSameFound(fmt.Sprintf("%s Sudah Digunakan", request.Email)))
	}
	penggunas := domain.Pengguna{
		Pengguna: request.Pengguna,
		Email:    request.Email,
		Sandi:    request.Sandi,
	}
	hashedPass, err := util.Hashpassword(penggunas.Sandi)
	helper.PanicError(err)
	penggunas.Sandi = hashedPass
	penggunas = service.PenggunaRepository.Save(ctx, tx, penggunas)
	return helper.ToPenggunaResponse(penggunas)
}

// FindAll implements UserService.
func (service *UserServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	penggunas := service.PenggunaRepository.FindAll(ctx, tx)
	return helper.ToPenggunaResponses(penggunas)
}

// FindById implements UserService.
func (service *UserServiceImpl) FindById(ctx context.Context, penggunaId int) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	penggunas := service.PenggunaRepository.FindById(ctx, tx, penggunaId)
	if penggunas.Id != penggunaId {
		panic(exception.NewNotFound("Data Tidak Ditemukan"))
	}
	return helper.ToPenggunaResponse(penggunas)
}

// FindByUserLogin implements UserService.
func (service *UserServiceImpl) FindByUserLogin(ctx context.Context, namaPengguna string) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	penggunas := service.PenggunaRepository.FindByPenggunaLogin(ctx, tx, namaPengguna)
	return helper.ToPenggunaResponse(penggunas)
}

// Update implements UserService.
func (service *UserServiceImpl) Update(ctx context.Context, update web.PenggunaUpdate) web.UserResponse {
	service.Validate.RegisterValidation("alphanumdash", util.ValidateAlphanumdash)
	err := service.Validate.Struct(update)
	util.ErrValidateSelf(err)
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	penggunas := service.PenggunaRepository.FindById(ctx, tx, update.Id)
	if penggunas.Id != update.Id {
		panic(exception.NewNotFound("Data Tidak Ditemukan"))
	}
	penggunas.Pengguna = update.Pengguna
	penggunas.Email = update.Email
	penggunas.Sandi = update.Sandi
	hashedPass, err := util.Hashpassword(penggunas.Sandi)
	helper.PanicError(err)
	penggunas.Sandi = hashedPass
	penggunass := service.PenggunaRepository.FindByPenggunaRegister(ctx, tx, update.Pengguna, update.Email)
	if penggunas.Id == penggunass.Id {
		penggunas = service.PenggunaRepository.Update(ctx, tx, penggunas)
		return helper.ToPenggunaResponse(penggunas)
	}
	if penggunas.Id != penggunass.Id {
		if penggunas.Pengguna == penggunass.Pengguna {
			panic(exception.NewSameFound(fmt.Sprintf("%s Sudah Digunakan", update.Pengguna)))
		}
		if penggunas.Email == penggunass.Email {
			panic(exception.NewSameFound(fmt.Sprintf("%s Sudah Digunakan", update.Email)))
		}
	}
	penggunas = service.PenggunaRepository.Update(ctx, tx, penggunas)
	return helper.ToPenggunaResponse(penggunas)
}

func (service *UserServiceImpl) LoginAuth(ctx context.Context, request web.LoginRequest) web.LoginResponse {
	// service.Validate.RegisterValidation("alphanumdash", util.ValidateAlphanumdash)
	err := service.Validate.Struct(request)
	util.ErrValidateSelf(err)
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	userData := service.PenggunaRepository.LoginAuth(ctx, tx, request.UserOrEmail)
	isvalid := util.Unhashpassword(request.Password, userData.Password)
	if !isvalid {
		panic(exception.NewNotFound("username Atau Email Dan Password Tidak Sesuai"))
	}
	claims := jwt.MapClaims{}
	claims["Username"] = userData.Username
	claims["id"] = userData.Id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token, err := util.GenerateToken(&claims)
	helper.PanicError(err)
	return helper.ToLoginResponse(token)
}
