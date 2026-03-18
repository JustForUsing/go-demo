+ 此项目是一个用来折腾，从0到1，`简单的`Demo演示项目，主要目的是展示代码风格和折腾过Go相关的一些东西

  + MakeFile应用

     ```bash
    make build    # 构建Go应用和数据库镜像
    make run      # 启动所有docker相关容器
    make stop     # 停止所有docker容器
    make clean    # 清理所有容器和镜像
    make test     # 运行测试
    make logs     # 查看应用容器日志
    make web-air  # 热更新启动
    ```

  + air自动热加载

  + 动态配置文件的设置

    + 多数据库驱动支持，配置文件切换，configs/default.yaml（database.driver）

  + 一些比较有趣的扩展应用

    + go.uber.org/fx   依赖注入组件，代码写起来会更优雅一点
  + github.com/golang-migrate   比gorm自动根据模型建表更可靠一点
  
  + 仿`laravel`的一些封装

    ````
  	//仿env函数，快捷获取指定配置
    global.GetViperConfigString("database.dbname")
    
    //仿laravel orm查询对gorm进一步封装，经过性能评估，对性能开销影响较小
    internal/repos/query_builder.go
        //配合使用示例文件
        internal/repos/user_repo.go
    ````
  
  + 可以从入口文件`cmd/web/main.go`或`internal/pkg/engine/router.go`路由文件开始跟踪