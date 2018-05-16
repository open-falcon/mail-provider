mail-provider
=============

把smtp封装为一个简单http接口，配置到sender中用来发送报警邮件

## 安装方法

1.二进制安装(推荐)
下载编译好的二进制
```bash
wget http://cactifans.hi-www.com/open-falcon/mail-provider.tar.gz
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
git clone https://github.com/GitHamburg/mail-provider.git
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
下载之后为源码，需要编译

```
curl http://$ip:4000/sender/mail -d "tos=a@a.com,b@b.com&subject=xx&content=yy"
```

## FAQ

1.此插件目前支持smtp SSL协议
```
"type": "smtp_ssl"
```

2.对于126.163等邮箱请控制发信频率以免被封
