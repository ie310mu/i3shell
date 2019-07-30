i3shell是一个golang开发的命令行工具，用于开启一个http端口，通过此端口访问指定目录下的静态网页文件（html、css、js、img）
文档地址：http://ie310.cn/i3shell

下载源码：git clone https://github.com/ie310mu/i3blog
编译：go build -o i3shell.exe main.go   (windows)
          go build -o i3shell main.go          (linux)
您也可直接下载随源码附带的编译结果文件
          https://github.com/ie310mu/i3blog/i3shell.exe   (windows)
          https://github.com/ie310mu/i3blog/i3shell          (linux)

使用：
          i3shell   (默认使用8093端口，被占用时抛出异常)
          i3shell -p 8093   (使用指定端口，被占用时抛出异常)
          i3shell -p 8093 -d F:\TestFiles (默认使用当前工作目录，可指定目标目录)

如果需要实现动态接口的支持，请查看ie310go项目，源码地址：https://github.com/ie310mu/ie310go，文档地址：http://ie310.cn/ie310go
如果需要实现更高级、灵活的代理功能，请考虑使用nginx等软件
