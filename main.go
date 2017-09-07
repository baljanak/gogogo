package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/coreos/etcd/client"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func WatchEtcd(e, cfile string) {

	cfg := client.Config{
		Endpoints:               []string{e},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	kapi := client.NewKeysAPI(c)

	watcher := kapi.Watcher(cfile, nil)
	for {
		resp, err := watcher.Next(context.TODO())
		if err != nil {
			if _, ok := err.(*client.ClusterError); ok {
				log.Error(err)
				time.Sleep(time.Second * 5)
				continue
			}
			log.Fatal(err)
		}
		log.Info("Updated value:\n%s", resp.Node.Value)
		log.Info("go DoSomething()")
	}
}

func main() {
	l := os.Getenv("HTTP_ADDR")
	e := os.Getenv("ETCD")
	cfile := "/test/config/deployment.yaml"

	go WatchEtcd(e, cfile)

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Info("Starting http server on ", l)
	http.ListenAndServe(l, router)
}
