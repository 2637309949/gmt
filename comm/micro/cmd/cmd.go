package cmd

import (
	"comm/conf"
	"comm/jwt"
	"net/http"
	"regexp"
	"strings"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/coreos/pkg/httputil"
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

func setPlugin(name string, h plugin.Handler) {
	plugin.Register(plugin.NewPlugin(plugin.WithName(name), plugin.WithHandler(h)))
}

func authp(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		whiteList := conf.Load("comm", "jwt_white_list").StringList([]string{})
		for i := range whiteList {
			match, _ := regexp.MatchString(whiteList[i], r.URL.Path)
			if match {
				h.ServeHTTP(w, r)
				return
			}
		}

		jwtSecret := conf.Load("comm", "jwt_secret").String()
		raw := r.Header.Get("Authorization")
		raw = strings.ReplaceAll(raw, "Bearer ", "")
		dtk, e := jwt.Decode(jwtSecret, raw)
		if e != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		r.Header.Set("x-user", dtk.User)
		h.ServeHTTP(w, r)
	})
}

func hystrixp(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.Method + "-" + r.RequestURI
		config := hystrix.CommandConfig{
			Timeout: 300,
		}
		hystrix.ConfigureCommand(name, config)
		if err := hystrix.Do(name,
			func() error {
				h.ServeHTTP(w, r)
				return nil
			},
			func(err error) error {
				return httputil.WriteJSONResponse(w, http.StatusBadGateway, err)
			},
		); err != nil {
			httputil.WriteJSONResponse(w, http.StatusInternalServerError, err)
			return
		}
	})
}

func Init() {
	setPlugin("auth", authp)
	setPlugin("hystrix", hystrixp)
	setRegistry("consul", func(opts ...registry.Option) registry.Registry {
		return cache.New(consul.NewRegistry(opts...))
	})
	cmd.Init()
}
