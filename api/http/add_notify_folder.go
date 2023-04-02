package httpapi

//
//import (
//	"os"
//
//	fsn "fsm_client/pkg/fsnotify"
//
//	"github.com/gin-gonic/gin"
//)
//
//func (h *Handle) AddSync() {
//
//}
//
//func (h *Handle) AddWatchPath(c *gin.Context) {
//	value := c.Query("key")
//	if value == "" {
//		c.AbortWithStatusJSON(400, gin.H{
//			"msg": "请登陆后重试",
//		})
//		return
//	}
//
//	_, err := os.Stat(value)
//
//	//err := h.watch.Add(value)
//	if err != nil {
//		c.AbortWithStatusJSON(400, gin.H{
//			"msg": "请登陆后重试",
//		})
//		return
//	}
//
//	c.JSON(200, gin.H{
//		"msg": "添加成功",
//	})
//
//}
//
//func (h *Handle) RemoveNotifyPath(c *gin.Context) {
//	value := c.Query("key")
//	if value == "" {
//		c.AbortWithStatusJSON(400, gin.H{
//			"msg": "请登陆后重试",
//		})
//		return
//	}
//
//	err := h.watch.Remove(value)
//	if err != nil {
//		c.AbortWithStatusJSON(400, gin.H{
//			"msg": "请登陆后重试",
//		})
//		return
//	}
//}
//
//func AddNotifyFolder(c *gin.Context) {
//	value := c.Query("key")
//	if value == "" {
//		return
//	}
//
//	fsn.TasksChan <- fsn.Task{
//		Name: value,
//		Op:   "add",
//	}
//
//	c.JSON(200, gin.H{
//		"res": "add notify folder success",
//	})
//}
//
//func DelNotifyFolder(c *gin.Context) {
//
//	value := c.Query("key")
//	if value == "" {
//		return
//	}
//
//	fsn.TasksChan <- fsn.Task{
//		Name: value,
//		Op:   "del",
//	}
//
//	c.JSON(200, gin.H{
//		"res": "del notify folder success",
//	})
//}
