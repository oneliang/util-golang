module test

go 1.22.2

replace (
	github.com/oneliang/util-golang/atomic v0.0.0 => ./../atomic
	github.com/oneliang/util-golang/base v0.0.0 => ../base
	github.com/oneliang/util-golang/common v0.0.0 => ./../common
	github.com/oneliang/util-golang/concurrent v0.0.0 => ./../concurrent
	github.com/oneliang/util-golang/file v0.0.0 => ./../file
	github.com/oneliang/util-golang/logging v0.0.0 => ./../logging
	github.com/oneliang/util-golang/resource v0.0.0 => ./../resource
	github.com/oneliang/util-golang/state v0.0.0 => ./../state
	github.com/oneliang/util-golang/goroutine v0.0.0 => ./../goroutine
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
	github.com/oneliang/util-golang/base v0.0.0 // indirect
	github.com/oneliang/util-golang/constants v0.0.0-20240424080518-b855d0299252 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
)
