# FuckingExam

自动获取粘贴板内容到题库中搜索题目, 并将搜索结果发送到QQ

## 需要环境

* golang
* cqhttp

## 如何使用

修改main.go文件中第69行将链接改成自己的cqhttp的请求地址, 第72行改成自己的QQ号或者群号

使用`go build main.go`构建应用程序, 双击main.exe开始运行, 支持所有系统

## 特别注意

本人对任何使用本程序造成的行为和后果不负责任

如果考试页面内无法复制, 使用tampermonkey的解除复制限制脚本