mail-provider
=============

把smtp封装为一个简单http接口，配置到sender中用来发送报警邮件

## 安装方法

1.二进制安装(推荐)   
sha256sum
```bash
7f714c8e4542670833505c274c420be1c89c5247dd8c72121cf0baa1d33edd49  falcon-mail-provider.tar.gz
```

下载编译好的二进制
```bash
wget https://dl.cactifans.com/open-falcon/falcon-mail-provider.tar.gz
mkdir -p falcon-mail-provider
tar zxvf falcon-mail-provider.tar.gz -C falcon-mail-provider
cd falcon-mail-provider
```
修改cfg.json文件相关信息，使用
```bash
./control start
```
即可启动

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
启动之后使用以下命令测试：
```
curl http://127.0.0.1:4000/sender/mail -d "tos=a@a.com,b@b.com&subject=xx&content=yy"
```
是否能收到邮件，如收到邮件，表示配置成功，如未收到邮件，使用
```bash
./control tail
```
查看日志。   
在Alarm组件的配置文件里，配置对应地址即可
```json
...
"api": {
       ...
        "mail": "http://127.0.0.1:4000/sender/mail",
       ...
    },
 ...
```

## FAQ

1.如使用自建邮件系统请设置 skipVerify 为 true 以避免证书校验错误，即使未开启TLS。（因为默认会尝试StartTLS）

2.对于126.163等邮箱请控制发信频率以免被封
