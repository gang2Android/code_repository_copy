# code_repository_copy

通过命令行的方式自动将源git的代码同步到目标git仓库中

根目录的config.json为需要同步的git仓库地址

```json
[
  {
    "source_repository": "源git地址,如：https://github.com/gang2Android/go_cron.git",
    "source_branch": "main",
    "target_repository": "目标git地址,如：https://github.com/gang2Android/go_cron1.git"
    "target_branch": "main2"
  }
]
```

部署

```shell
set GOOS=linux
set GOARCH=amd64
go build -o cron_app
# 上传到服务器目录下
mkdir sites
vi config.json

chmod 777 cron_app
nohup ./cron_app &
```
