package moka

import (
	"fmt"
	"os"
)

var LainDomain string      // lain.example.com
var PodInstanceName string // "foo.web.web"
var PodInstanceNo string   // "1", "2"
var Prefix string          // set your custom prefix
var MokaPrefix string

func init() {
	LainDomain = os.Getenv("LAIN_DOMAIN")
	PodInstanceNo = os.Getenv("DEPLOYD_POD_INSTANCE_NO")
	PodInstanceName = os.Getenv("DEPLOYD_POD_NAME")
	Prefix = os.Getenv("MOKA_PREFIX")
	MokaPrefix = fmt.Sprintf("%v.%v.%v.%v", Prefix, LainDomain, PodInstanceName, PodInstanceNo)
}
