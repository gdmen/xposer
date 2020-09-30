package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"gorm.io/gorm"

	"garymenezes.com/xposer/common"
)

func handlePing(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		logPrefix := common.GetLogPrefix(c)
		glog.Infof("%s fcn start", logPrefix)

		p := Ping{
			SourceIP:   c.ClientIP(),
			DeviceName: c.GetString("DeviceName"),
			ReceivedAt: c.GetTime("TimeReceived"),
		}
		if dbc := db.Create(&p); dbc.Error != nil {
			msg := "Failed to save ping."
			glog.Errorf("%s %s: %v", logPrefix, msg, dbc.Error)
			c.JSON(http.StatusBadRequest, gin.H{"message": msg})
			return
		}
		c.Status(http.StatusOK)
	}
}
