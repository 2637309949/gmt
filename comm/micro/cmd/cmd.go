package cmd

import (
	"comm/conf"
	"comm/jwt"
	"net/http"
	"regexp"

	cfg "github.com/micro/go-micro/v2/config/cmd"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/cache"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/micro/v2/cmd"
	"github.com/micro/micro/v2/plugin"
)

func setRegistry(name string, rr func(opts ...registry.Option) registry.Registry) {
	cfg.DefaultRegistries[name] = rr
}

func consulRegistry(opts ...registry.Option) registry.Registry {
	return cache.New(consul.NewRegistry(opts...))
}

func authHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		whiteList := conf.Load("comm", "jwt_white_list").StringList([]string{})
		for i := range whiteList {
			match, _ := regexp.MatchString(whiteList[i], r.URL.Path)
			if match {
				h.ServeHTTP(w, r)
				return
			}
		}
		dtk, e := jwt.Decode(r.Header.Get("Authorization"))
		if e != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		r.Header.Set("x-user", dtk.User)
		h.ServeHTTP(w, r)
	})
}

func init() {
	plugin.Register(plugin.NewPlugin(plugin.WithName("jwt"), plugin.WithHandler(authHandler)))
	setRegistry("consul", consulRegistry)
}

func Init() {
	cmd.Init()
}
