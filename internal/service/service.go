package service

import (
	"github.com/binbinly/pkg/cache"
	"github.com/binbinly/pkg/transport/ws"
	"github.com/redis/go-redis/v9"
	"project-layout/internal/app"
	"project-layout/internal/repository"
)

// 用于触发编译期的接口的合理性检查机制
var _ IService = (*Service)(nil)

// IService 服务接口定义
type IService interface {
	Close() error
}

var Svc IService

// Service struct
type Service struct {
	opts options
	repo repository.IRepo
	rdb  *redis.Client
}

// New init service
func New(ws ws.Server, opts ...Option) (s *Service) {
	rdb := app.InitRedis()
	s = &Service{
		opts: newOptions(opts...),
		repo: repository.New(app.InitDB(), cache.NewRedisCache(rdb)),
		rdb:  rdb,
	}
	Svc = s
	return s
}

// Close service
func (s *Service) Close() error {
	return s.rdb.Close()
}
