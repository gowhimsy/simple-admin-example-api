package svc

import (
	exampleClient "github.com/suyuan32/simple-admin-example-rpc/client/example"
	schoolClient "github.com/suyuan32/simple-admin-example-rpc/client/school"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/suyuan32/simple-admin-example-api/internal/config"
	i18n2 "github.com/suyuan32/simple-admin-example-api/internal/i18n"
	"github.com/suyuan32/simple-admin-example-api/internal/middleware"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config     config.Config
	ExampleRpc exampleClient.Example
	SchoolRpc  schoolClient.School
	Casbin     *casbin.Enforcer
	Authority  rest.Middleware
	Trans      *i18n.Translator
}

func NewServiceContext(c config.Config) *ServiceContext {

	rds := c.RedisConf.NewRedis()
	if !rds.Ping() {
		logx.Error("initialize redis failed")
		return nil
	}

	cbn, err := c.CasbinConf.NewCasbin(c.DatabaseConf.Type, c.DatabaseConf.GetDSN())
	if err != nil {
		logx.Errorw("initialize casbin failed", logx.Field("detail", err.Error()))
		return nil
	}

	trans := &i18n.Translator{}
	trans.NewBundle(i18n2.LocaleFS)
	trans.NewTranslator()

	return &ServiceContext{
		Config:     c,
		Authority:  middleware.NewAuthorityMiddleware(cbn, rds).Handle,
		Trans:      trans,
		ExampleRpc: exampleClient.NewExample(zrpc.MustNewClient(c.ExampleRpc)),
		SchoolRpc:  schoolClient.NewSchool(zrpc.MustNewClient(c.SchoolRpc)),
	}
}
