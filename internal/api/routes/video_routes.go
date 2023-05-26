package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/ahsmha/vod/internal/api/handlers"
)

// VideoRoutes represents the video-related routes and handlers.
type VideoRoutes struct {
	router       *gin.Engine
	videoHandler handlers.VideoHandler
}

// NewVideoRoutes creates a new instance of VideoRoutes.
func NewVideoRoutes(router *gin.Engine, videoHandler handlers.VideoHandler) *VideoRoutes {
	return &VideoRoutes{
		router:       router,
		videoHandler: videoHandler,
	}
}

// Setup sets up the video routes.
func (vr *VideoRoutes) Setup() {
	videosGroup := vr.router.Group("/videos")

	videosGroup.GET("", vr.videoHandler.GetAllVideos)
	videosGroup.GET("/:videoID", vr.videoHandler.GetVideoByID)
	videosGroup.POST("", vr.videoHandler.UploadVideo)
	videosGroup.PUT("/:videoID", vr.videoHandler.UpdateVideo)
	videosGroup.DELETE("/:videoID", vr.videoHandler.DeleteVideo)
}
