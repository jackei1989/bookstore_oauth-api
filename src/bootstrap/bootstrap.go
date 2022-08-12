package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/jackei1989/bookstore_oauth-api/src/domain/access_token"
	"github.com/jackei1989/bookstore_oauth-api/src/http"
	"github.com/jackei1989/bookstore_oauth-api/src/repository/db"
)

var (
	router = gin.Default()
)

func BootApplication() {
	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	_ = router.Run(":8080")
}
