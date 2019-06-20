package controllers

import (
	"testing"
	"github.com/gin-gonic/gin"
	"webservice/db"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"net/http/httptest"
)

func TestUserListApiFail(t *testing.T) {
	uc := UserController{}
	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(resp)
	originConn := db.GetConn()
    conn, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer conn.Close()
	defer db.SetConn(originConn)
	sqlxDB := sqlx.NewDb(conn, "sqlmock")
	db.SetConn(sqlxDB)
	uc.HandleUserListRequest(ctx)
}

func TestUserListApiSuccess(t *testing.T) {
	uc := UserController{}
	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(resp)
	uc.HandleUserListRequest(ctx)
}
