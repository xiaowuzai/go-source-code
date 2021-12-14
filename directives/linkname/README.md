# linkname 简介
该目录是用来演示 `//go:linkname` 的一些操作

对于 `//go:linkname` 这个指令，在项目中是不建议大家使用的，因为会破坏项目结构
官方要求在使用`//go:linkname`的时候一定要引入 `unsafe`包，也是为了提示大家，这样操作是不安全的。

## 目录结构
example1 目录用来演示 `//go:linkname` 直接链接到导出函数的情况
example2 目录用来演示`//go:linkname` 非侵入式链接的情况

## example1解释
inner包里面包含了一个非导出函数`linka` 通过`//go:linkname`的方式链接到 outer包中的导出函数 `LinkA` 上
这样就可以通过执行 `outer.LinkA()` 来执行 inner 包非导出函数`linka`了。

在目录中的 `.s` 文件是用来规避编译器检查的。
