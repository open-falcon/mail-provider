mail-provider
=============

把smtp封装为一个简单http接口，配置到sender中用来发送报警邮件

## 安装方法

1.二进制安装(推荐)
下载编译好的二进制
```bash
wget https://dl.cactifans.com/open-falcon/mail-provider.tar.gz
mkdir -p mail-provider
tar zxvf mail-provider.tar.gz  -C mail-provider
cd mail-provider
```
修改cfg.json文件相关信息，使用
```bash
./control start
```
即可启动客户端

2.源码编译（如无科学上网方法，请勿尝试）
下载之后为源码，安装golang环境，环境配置参考[golang环境配置](http://book.open-falcon.org/zh/quick_install/prepare.html)
编译方法
```bash
cd $GOPATH/src
mkdir github.com/open-falcon/ -p
cd github.com/open-falcon/
git clone https://github.com/open-falcon/mail-provider.git
cd mail-provider
go get ./...
./control build
```
编译成功之后，修改cfg.json文件相关信息，使用
```bash
./control start
```
即可启动


## 使用方法

测试发送
```
curl http://$ip:4000/sender/mail -d "tos=a@a.com,b@b.com&subject=xx&content=yy"
```
在Alarm组件里修改配置mail地址为
```
 "mail": "http://127.0.0.1:4000/sender/mail",
```
## FAQ

1.此插件目前不支持smtp SSL协议

2.对于126.163等邮箱请控制发信频率以免被封
