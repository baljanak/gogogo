package etcdctl_test

import (
	"fmt"
	"time"

	"github.com/coreos/etcd/client"
	"github.com/ghodss/yaml"
	"golang.org/x/net/context"

	log "github.com/Sirupsen/logrus"
)

var e string = "http://127.0.0.1:2379"
var cfile string = "/test/config/person.yaml"

type Person struct {
	Name string `json:"name"` // Affects YAML field names too.
	Age  int    `json:"age"`
}

func ExampleEtcdClt() {
	cfg := client.Config{
		Endpoints: []string{e},
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	kapi := client.NewKeysAPI(c)

	// Marshal a Person struct to YAML.
	p := Person{"John", 30}
	y, err := yaml.Marshal(p)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	log.Print(string(y))

	resp, err := kapi.Set(context.Background(), cfile, string(y), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Set is done. Metadata is %q\n", resp)
		fmt.Println("Set is done.")
	}
	resp, err = kapi.Get(context.Background(), cfile, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		// print common key info
		log.Printf("Get is done. Metadata is %q\n", resp)
		fmt.Println("Get is done.")
		// print value
		log.Printf("%q key has %q value\n", resp.Node.Key, resp.Node.Value)
	}

	// Output:
	// Set is done.
	// Get is done.
}
