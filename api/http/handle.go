package httpapi

import (
	"fsm_client/pkg/ent"
	"fsm_client/pkg/httpclient"
	"fsm_client/pkg/sync"
	"fsm_client/pkg/types"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Handle struct {
	App    *gin.Engine
	Client *httpclient.Client
	Syncer *sync.Syncer
	DB     *gorm.DB
}

func New(app *gin.Engine, client *httpclient.Client, syncer *sync.Syncer, db *gorm.DB) *Handle {
	h := &Handle{App: app, Client: client, Syncer: syncer, DB: db}
	h.initRoute()
	return h
}

func (h *Handle) initRoute() {
	h.App.POST("/login", h.Login)
	h.App.GET("/logout", h.Logout)
	h.App.POST("/register", h.Register)
	h.App.GET("/status", h.GetLoginStatus)

	h.App.GET("/syncTask", h.GetSyncTask)
	h.App.POST("/syncTask", h.NewSyncTask)
	h.App.DELETE("/syncTask", h.RemoveSyncTask)
	h.App.GET("/syncTasks", h.GetSyncTasks)

	h.App.POST("/recover", h.RecoverSyncTask)
}

func (h *Handle) RecoverSyncTask(c *gin.Context) {
	var st types.RecSyncTask
	if err := c.BindJSON(&st); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	log.Println(st)
	return

	//todo
	if err := h.Syncer.RecoverTask(st); err != nil {
		return
	}

}

func (h *Handle) GetLoginStatus(c *gin.Context) {
	if h.Client.UserID == "" || h.Client.JWT == "" {
		c.JSON(http.StatusBadGateway, "客户端未登录初始化")
		return
	}
	c.JSON(http.StatusOK, "客户端登录初始化成功")
}

func (h *Handle) Register(c *gin.Context) {
	var user types.UserRegister
	if err := c.BindJSON(&user); err != nil {
		log.Println(err)
	}

	if err := h.Client.Register(user); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *Handle) Login(c *gin.Context) {
	var user types.UserLoginReq
	if err := c.BindJSON(&user); err != nil {
		log.Println(err)
	}

	if err := h.Client.Login(user); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *Handle) Logout(c *gin.Context) {
	h.Client.UserID = ""
	h.Client.JWT = ""
	h.Client.HttpClient = nil
	c.JSON(http.StatusOK, nil)
}

func (h *Handle) NewSyncTask(c *gin.Context) {
	var st types.NewSyncTask
	if err := c.BindJSON(&st); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	log.Println(st)
	return

	if err := h.Syncer.CreateSyncTask(st); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *Handle) RemoveSyncTask(c *gin.Context) {

}

func (h *Handle) GetSyncTask(c *gin.Context) {

}

func (h *Handle) GetSyncTasks(c *gin.Context) {
	var st []ent.SyncTask
	h.DB.Find(&st)
	c.JSON(http.StatusOK, st)
}
