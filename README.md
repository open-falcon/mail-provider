mail-provider
=============

把smtp封装为一个简单http接口，配置到sender中用来发送报警邮件

## 使用方法

```
curl http://$ip:4000/sender/mail -d "tos=a@a.com,b@b.com&subject=xx&content=yy"
```
