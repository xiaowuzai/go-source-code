# linkname 简介
该目录是用来演示 `//go:linkname` 的一些操作

对于 `//go:linkname` 这个指令，在项目中是不建议大家使用的，因为会破坏项目结构
官方要求在使用`//go:linkname`的时候一定要引入 `unsafe`包，也是为了提示大家，这样操作是不安全的。

## 目录结构
example1 目录用来演示 `//go:linkname` 直接链接到导出函数的情况
example2 目录用来演示`//go:linkname` 非侵入式链接的情况

## example1 解释
inner包里面包含了一个非导出函数`linka` 通过`//go:linkname`的方式链接到 outer包中的导出函数 `LinkA` 上
这样就可以通过执行 `outer.LinkA()` 来执行 inner 包非导出函数`linka`了。

在目录中的 `.s` 文件是用来规避编译器检查的。
因为在执行 `go run` 或者 `go build` 的时候，默认会对 go 语法进行完整性校验

这样虽然可以访问 `inner` 包中的私有函数，但是这种方式很不灵活

## example2 解释
inner包里面包含了一个非导出函数`linkb` 通过`//go:linkname`的方式给自己起了一个符号名。
在 outer 包同样通过 `//go:linkname` 的方式来进行链接。
上这样就可以通过执行 `outer.LinkB()` 来间接执行 inner 包非导出函数`linkb`了。

## 注意要点
1. 如果包内要使用`//go:linkname` 那么一定要导入 `unsafe` 包，如果不需要使用，可以通过 `import ( _ "unsafe" )` 来忽略引入未使用的包
2. 在调用方（即 outer 包）一定要 `import` 实现函数的包（即：import inner 包）
3. 不建议在项目中使用这种方式，因为会影响程序结构。


