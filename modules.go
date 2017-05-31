package main

import (
	_ "github.com/callixbrasil/registrator/consul"
	_ "github.com/callixbrasil/registrator/consulkv"
	_ "github.com/callixbrasil/registrator/etcd"
	_ "github.com/callixbrasil/registrator/skydns2"
	_ "github.com/callixbrasil/registrator/zookeeper"
)
