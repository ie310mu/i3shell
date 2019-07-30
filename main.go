package main

import (
	"flag"
	"fmt"
	"i3work/utils/dirutil"
	"time"

	"github.com/ie310mu/ie310go"
	"github.com/ie310mu/ie310go/common/logsagent"
	"github.com/ie310mu/ie310go/common/pc"
	"github.com/ie310mu/ie310go/route"
)

func main() {
	//异常处理
	defer func() {
		if err := recover(); err != nil {
			logsagent.Error("Exception has been caught. ", err)
			logsagent.Warn("Press return to exit.")
			var empty string
			fmt.Scanln(&empty)
		}
	}()

	if !appinit() {
		return
	}

	go startInWebBrowser()
	ie310go.Run(nil)
}

func startInWebBrowser() {
	<-time.After(500 * time.Millisecond)
	pc.OpenUrl("http://127.0.0.1:" + *p + "/")
}

var (
	h = flag.Bool("h", false, "")
	p = flag.String("p", "8093", "the port of service")
	d = flag.String("d", "", "the dir, default is current work dir")
)

func appinit() bool {
	flag.Parse()
	if *h {
		flag.PrintDefaults()
		return false
	}

	ie310go.AppLogsInitFunc = nil
	ie310go.InitLogs(`{"filename":"i3shell.log"}`)

	httpConfig := route.ServerHTTPConfig{Port: *p, ServeStaticFunc: route.DefaultServeStatic}
	if *d != "" {
		httpConfig.DefaultStaticDir = *d
	} else {
		httpConfig.DefaultStaticDir = dirutil.GetCurrentPath()
	}
	logsagent.Info("the path is " + httpConfig.DefaultStaticDir)
	srv := route.NewServerHTTP(httpConfig, "httpServer")

	route.RegisterServer(srv)
	return true
}
