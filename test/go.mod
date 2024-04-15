module test

go 1.21.1

replace (
	github.com/oneliang/util-golang/atomic v0.0.0 => ./../atomic
	github.com/oneliang/util-golang/base v0.0.0 => ../base
	github.com/oneliang/util-golang/common v0.0.0 => ./../common
	github.com/oneliang/util-golang/concurrent v0.0.0 => ./../concurrent
	github.com/oneliang/util-golang/file v0.0.0 => ./../file
	github.com/oneliang/util-golang/logging v0.0.0 => ./../logging
	github.com/oneliang/util-golang/resource v0.0.0 => ./../resource
	github.com/oneliang/util-golang/state v0.0.0 => ./../state
)

require (
	github.com/eclipse/paho.mqtt.golang v1.4.3 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
)
