google/wire 是 Go 语言的编译时依赖注入框架，与 Spring IoC 一样，wire 的目的也是让开发者从对项目中大量依赖的创建和管理中解脱出来，但两者在实现方式上有着很大的不同。

安装命令：
`go get github.com/google/wire/cmd/wire`

在开始之前，需要先了解 wire 中的两个核心概念：Provider, Injector。

## Provider

用于创建组件的普通函数。这些函数将所需的依赖作为参数传入，创建组件并返回。

这里的组件可以是任何类型，但不像 Spring 中可以通过给同一类型组件不同的名字完成依赖注入，wire 只支持一个类型对应一个 Provider。因此 Provider 不应该被用来返回 int、string 等简单类型。

示例：

```go
func NewService(client db.Client, config *conf.Config) *Service {
	return nil
}

```

## Injector

Injector 是项目所有组件(Provider)及其依赖关系的声明， wire 根据 Injector 生成依赖注入代码。

##

参考：

1. [使用 google/wire 对 Go 项目进行依赖注入](https://www.jianshu.com/p/7d05a1c71d08)
2. [Mastering Wire](https://itnext.io/mastering-wire-f1226717bbac)
3. [官方指南](https://github.com/google/wire/blob/main/docs/guide.md)
4. [CodeReviewComments#interfaces](https://github.com/golang/go/wiki/CodeReviewComments#interfaces)