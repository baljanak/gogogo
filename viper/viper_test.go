package viper_test

import (
	"fmt"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

var e string = "http://127.0.0.1:2379"
var cfile string = "/test/config/person.yaml"

type Person struct {
	Name string `json:"name"` // Affects YAML field names too.
	Age  int    `json:"age"`
}

func ExampleViperRemote() {
	p := Person{}
	v := viper.New()
	v.AddRemoteProvider("etcd", e, cfile)
	v.SetConfigType("yaml")
	v.ReadRemoteConfig()
	v.Unmarshal(&p)

	fmt.Println(p)

	//Output:
	// {John 30}
}
