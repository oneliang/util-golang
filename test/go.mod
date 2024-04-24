module test

go 1.22.0

toolchain go1.22.2

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
	github.com/ByteArena/poly2tri-go v0.0.0-20170716161910-d102ad91854f // indirect
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/benoitkugler/textlayout v0.3.0 // indirect
	github.com/benoitkugler/textprocessing v0.0.3 // indirect
	github.com/eclipse/paho.mqtt.golang v1.4.3 // indirect
	github.com/go-fonts/latin-modern v0.3.1 // indirect
	github.com/go-text/typesetting v0.0.0-20231013144250-6cc35dbfae7d // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/llgcode/draw2d v0.0.0-20240322162412-ee6987bd01dc // indirect
	github.com/tdewolff/canvas v0.0.0-20240420213651-d5a04e36ef50 // indirect
	github.com/tdewolff/font v0.0.0-20240417221047-e5855237f87b // indirect
	github.com/tdewolff/minify/v2 v2.20.5 // indirect
	github.com/tdewolff/parse/v2 v2.7.3 // indirect
	golang.org/x/image v0.15.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gopkg.in/gcfg.v1 v1.2.3 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	star-tex.org/x/tex v0.4.0 // indirect
)
