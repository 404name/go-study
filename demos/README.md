
一些框架demo
- chatgpt测试
- cobra是一个命令行启动框架，k8s etcd都是采用的这个
  - 构建一个携带服务的命令行工具 `cobra service` `cobra version` 
  - 产物exe启动自带gin后端+前端文件
  - 通过`go embed`映射静态文件到FS，go build打包时FS会携带静态文件嵌入
  - 启动时如果没有static文件夹则从嵌入式文件系统中提取静态文件