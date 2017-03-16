package controller

import (
	"net/http"
	"github.com/mageddo/dns-proxy-server/events/local"
	"encoding/json"
	"golang.org/x/net/context"
	"github.com/mageddo/log"
)

func init(){
	Get("/hello/", func(ctx context.Context, res http.ResponseWriter, req *http.Request, url string){
		res.Write([]byte("It works from controller!!!"))
	})

	Get("/hostname/", func(ctx context.Context, res http.ResponseWriter, req *http.Request, url string){
		res.Header().Add("Content-Type", "application/json")
		json.NewEncoder(res).Encode(local.GetConfiguration(ctx))
	})

	Post("/hostname/", func(ctx context.Context, res http.ResponseWriter, req *http.Request, url string){
		logger := log.GetLogger(ctx)
		res.Header().Add("Content-Type", "application/json")
		logger.Infof("m=/hostname/, status=begin, action=create-hostname")
		var hostname local.HostnameVo
		json.NewDecoder(req.Body).Decode(&hostname)
		logger.Infof("m=/hostname/, status=parsed-host, host=%+v", hostname)
		err := local.AddHostname(ctx, hostname.Env, hostname)
		if err != nil {
			logger.Infof("m=/hostname/, status=error, action=create-hostname, err=%+v", err)
			BadRequest(res, "Env not found")
		}
		logger.Infof("m=/hostname/, status=success, action=create-hostname")
	})

	Delete("/hostname/", func(ctx context.Context, res http.ResponseWriter, req *http.Request, url string){
		logger := log.GetLogger(ctx)
		res.Header().Add("Content-Type", "application/json")
		logger.Infof("m=/hostname/, status=begin, action=delete-hostname")
		var hostname local.HostnameVo
		json.NewDecoder(req.Body).Decode(&hostname)
		logger.Infof("m=/hostname/, status=parsed-host, action=delete-hostname, host=%+v", hostname)
		err := local.RemoveHostname(ctx, hostname.Env, hostname)
		if err != nil {
			logger.Infof("m=/hostname/, status=error, action=delete-hostname, err=%+v", err)
			BadRequest(res, "Env not found")
		}
		logger.Infof("m=/hostname/, status=success, action=delete-hostname")
	})
}