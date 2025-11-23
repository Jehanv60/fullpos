package service

import (
	"context"

	"github.com/Jehanv60/model/web"
)

type BarangService interface {
	Create(ctx context.Context, request web.BarangCreateRequest, idUser int) web.BarangResponse
	Update(ctx context.Context, update web.BarangUpdate, idUser int) web.BarangResponse
	Delete(ctx context.Context, barangId int, idUser int)
	FindById(ctx context.Context, barangId int, idUser int) web.BarangResponse
	FindAll(ctx context.Context, idUser int) []web.BarangResponse
}

type UserService interface {
	Create(ctx context.Context, request web.PenggunaCreateRequest) web.UserResponse
	Update(ctx context.Context, update web.PenggunaUpdate) web.UserResponse
	FindById(ctx context.Context, penggunaId int) web.UserResponse
	FindByUserLogin(ctx context.Context, NamaPengguna string) web.UserResponse
	LoginAuth(ctx context.Context, request web.LoginRequest) web.LoginResponse
	FindAll(ctx context.Context) []web.UserResponse
}

type TransaksiService interface {
	Create(ctx context.Context, request web.TransactionCreateRequest, idUser int) web.TransaksiResponse
	ReportAll(ctx context.Context, idUser int) []web.TransaksiResponse
}
