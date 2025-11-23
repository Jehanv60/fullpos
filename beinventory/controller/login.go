package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/Jehanv60/helper"
	"github.com/Jehanv60/model/web"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

// function untuk login
func (controller *UserControllerImpl) LoginAuth(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userCreateRequest := web.LoginRequest{}
	helper.ReadFromBody(r, &userCreateRequest)
	userData := controller.UserService.LoginAuth(r.Context(), userCreateRequest)
	helper.GoDoEnv()
	writeToken := &http.Cookie{
		Name:     os.Getenv("Token"),
		Value:    userData.Token,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 1),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, writeToken)
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Credentials", "true")
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   "Username Atau Email Dan Password Sesuai",
	}
	helper.WriteToResponse(w, webResponse)
}

/*
contoh UserOrEmailn tanpa interface
var DB = NewDb1()

	func NewDb1() *sql.DB {
		var err error
		var db *sql.DB
		db, err = sql.Open("postgres", "host=localhost user=han port=5432 password=solo dbname=pos1 sslmode=disable")
		if err != nil {
			panic(err)
		}
		db.SetMaxIdleConns(5)
		db.SetMaxOpenConns(20)
		db.SetConnMaxIdleTime(10 * time.Minute)
		db.SetConnMaxLifetime(60 * time.Minute)
		return db
	}

UserOrEmailId, err := UserOrEmailSelect(userCreateRequest.UserOrEmail)
helper.PanicError(err)

	func UserOrEmailSelect(NamaUserOrEmail string) (domain.UserOrEmail, error) {
		var err error
		tx, err := DB.Begin()
		helper.PanicError(err)
		SQL := "select id, UserOrEmail, email, password from UserOrEmail where UserOrEmail = $1"
		rows, err := tx.Query(SQL, NamaUserOrEmail)
		helper.PanicError(err)
		UserOrEmail := domain.UserOrEmail{}
		fmt.Println(UserOrEmail)
		defer rows.Close()
		if rows.Next() {
			rows.Scan(&UserOrEmail.Id, &UserOrEmail.UserOrEmail, &UserOrEmail.Email, &UserOrEmail.Sandi)
			return UserOrEmail, nil
		} else {
			return UserOrEmail, nil
		}

}
*/
