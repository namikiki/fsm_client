package httpapi

import (
	"fsm_client/pkg/ent"

	"github.com/gin-gonic/gin"
)

func (h *Handle) GetTasks(c *gin.Context) {
	var tasks []ent.SyncTask
	h.DB.Find(&tasks)
	c.JSON(200, tasks)
}

func (h *Handle) Create(c *gin.Context) {
	if err := h.Sync.CreateSyncTask(c.Param("name"), c.Param("root")); err != nil {
		c.JSON(301, err)
		return
	}
	c.JSON(200, "创建成功")
}

func (h *Handle) Delete() {

}
