package api

import (
	"github.com/binbinly/pkg/errno"
	"github.com/binbinly/pkg/logger"
	"github.com/binbinly/pkg/util"
	"github.com/gin-gonic/gin"
)

const (
	// _defaultLimit 默认分页大小
	_defaultLimit = 20
)

// BindJSON 绑定请求参数
func BindJSON(c *gin.Context, form any) error {
	if err := c.ShouldBindJSON(form); err != nil {
		logger.Debugf("[api.bind.json] param err: %v", err)
		return err
	}

	return nil
}

// GetUserID 返回用户id
func GetUserID(c *gin.Context) int {
	if c == nil {
		return 0
	}
	// uid 必须和 middleware/auth 中的 uid 命名一致
	return c.GetInt("uid")
}

// GetPage 获取分页起始偏移量
func GetPage(c *gin.Context) (int, int) {
	offset := 0
	page := util.MustInt(c.Query("p"))
	if page > 0 {
		offset = (page - 1) * _defaultLimit
	}
	return offset, _defaultLimit
}

// Error response err
func Error(err error) *errno.Error {
	switch {
	case err != nil:
		logger.Warnf("[api] err:%v", err)
		return errno.ErrDatabase
	default:
		return nil
	}
}
