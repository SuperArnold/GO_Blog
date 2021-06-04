package routers

import (
	"time"

	"github.com/SuperArnold/GO_Blog/global"
	"github.com/SuperArnold/GO_Blog/internal/middleware"
	v1 "github.com/SuperArnold/GO_Blog/internal/routers/api/v1"
	"github.com/SuperArnold/GO_Blog/pkg/limiter"
	"github.com/gin-gonic/gin"

	_ "github.com/SuperArnold/GO_Blog/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {

	r := gin.New()

	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(60 * time.Second))
	r.Use(middleware.Translations())

	article := v1.NewArticle()
	tag := v1.NewTag()
	upload := NewUpload()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.GET("/testEmail", middleware.Recovery())
	url := ginSwagger.URL("http://127.0.0.1:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.POST("/upload/file", upload.UploadFile)
	r.POST("/auth", v1.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("tags/:id", tag.Delete)
		apiv1.PUT("tags/:id", tag.Update)
		apiv1.PATCH("tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("articles/:id", article.Delete)
		apiv1.PUT("articles/:id", article.Update)
		apiv1.PATCH("articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)

	}

	return r
}
