module test

go 1.24.10

replace (
	github.com/oneliang/util-golang/atomic v0.0.0 => ./../atomic
	github.com/oneliang/util-golang/base v0.0.0 => ../base
	github.com/oneliang/util-golang/common v0.0.0 => ./../common
	github.com/oneliang/util-golang/concurrent v0.0.0 => ./../concurrent
	github.com/oneliang/util-golang/file v0.0.0 => ./../file
	github.com/oneliang/util-golang/goroutine v0.0.0 => ./../goroutine
	github.com/oneliang/util-golang/logging v0.0.0 => ./../logging
	github.com/oneliang/util-golang/resource v0.0.0 => ./../resource
	github.com/oneliang/util-golang/state v0.0.0 => ./../state
)

require (
	github.com/eclipse/paho.mqtt.golang v1.4.3
	github.com/oneliang/util-golang/atomic v0.0.0
	github.com/oneliang/util-golang/common v0.0.0
	github.com/oneliang/util-golang/concurrent v0.0.0
	github.com/oneliang/util-golang/file v0.0.0
	github.com/oneliang/util-golang/logging v0.0.0
	github.com/oneliang/util-golang/logging_ext v0.0.0-20240424080518-b855d0299252
	github.com/oneliang/util-golang/resource v0.0.0
	github.com/oneliang/util-golang/state v0.0.0
)

require (
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/mdns v1.0.5 // indirect
	github.com/miekg/dns v1.1.41 // indirect
	github.com/oneliang/util-golang/base v0.0.0 // indirect
	github.com/oneliang/util-golang/constants v0.0.0-20240424080518-b855d0299252 // indirect
	github.com/vishvananda/netlink v1.1.0 // indirect
	github.com/vishvananda/netns v0.0.0-20191106174202-0a2b9b5464df // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.23.0 // indirect
)
