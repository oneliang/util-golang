module test

go 1.21.1

replace (
	"github.com/oneliang/util-golang/base" v0.0.0 => ../base
	"github.com/oneliang/util-golang/common" v0.0.0 => ./../common
	"github.com/oneliang/util-golang/atomic" v0.0.0 => ./../atomic
)
