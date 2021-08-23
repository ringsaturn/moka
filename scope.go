package moka

import (
	"fmt"
	"os"
)

type Collector struct {
	LAINDOMAIN           string
	PodInstanceNo        string
	PodInstanceName      string
	PodInstanceNamespace string
}

var LainDomain string      // lain.example.com
var PodInstanceName string // "foo.web.web"
var PodInstanceNo string   // "1", "2"
var Prefix string          // set your custom prefix
var Scope string

func init() {
	LainDomain = os.Getenv("LAIN_DOMAIN")
	PodInstanceNo = os.Getenv("DEPLOYD_POD_INSTANCE_NO")
	PodInstanceName = os.Getenv("DEPLOYD_POD_NAME")
	Prefix = os.Getenv("MOKA_PREFIX")
	Scope = fmt.Sprintf("%v.%v.%v.%v", Prefix, LainDomain, PodInstanceName, PodInstanceNo)
}
