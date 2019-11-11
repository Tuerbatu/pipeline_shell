# 管道壳子
这是一个基于netty的go语言版本，如果你想使用它，只需要引用bootstrap包，至于如何使用可以看这个项目自带的例子。你只需要注意mian.go中的配置以及要
加载的header。消息之间通过interface来传播，如果有你的header有需要的数据类型则通过断言将数据下载下来。如果你不想让一些数据向后传，你只需要将数据
清空就会避免其向后传递。
# pipeline_shell
This is a netty - based version of go. If you want to use it, you just need to refer to bootstrap. All you need to notice is the
configuration in mian. Go and the headers to load. Messages are propagated through an interface, and if you have a header
and you have the type of data that you want, you download the data through assertions. If you don't want some data to be
passed backwards, you just need to clear the data to avoid it.
