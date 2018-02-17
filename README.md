# mail-provider

把smtp封装为一个简单http接口，配置到sender中用来发送报警邮件

## 配置

拷贝 cfg.example.json 为 cfg.json，修改邮件配置，如果使用ssl，那么需要将配置中的 `ssl` 参数设置为true，并且配置 smtp 的服务地址为ssl地址和端口。

`from` 参数可以随意填写，比如 `监控`，也可以设置格式为 `显示名字<email>`，比如 `监控<monitor@example.com>`，这样收件方就会显示配置的显示名字。
## 使用方法

```
curl http://$ip:4000/sender/mail -d "tos=a@a.com,b@b.com&subject=xx&content=yy"
```
