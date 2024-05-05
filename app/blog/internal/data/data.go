package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"my-kratos-blog/app/blog/internal/conf"
)

type Data struct {
}

func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}
