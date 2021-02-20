## 使用Go实现TLS 服务器和客户端

传输层安全协议（Transport Layer Security，缩写：TLS），及其前身安全套接层（Secure Sockets Layer，缩写：SSL）是一种安全协议，
目的是为互联网通信提供安全及数据完整性保障。

SSL包含记录层（Record Layer）和传输层，记录层协议确定传输层数据的封装格式。传输层安全协议使用X.509认证，之后利用非对称加密演算来对
通信方做身份认证，之后交换对称密钥作为会谈密钥（Session key）。这个会谈密钥是用来将通信两方交换的资料做加密，保证两个应用间通信的保
密性和可靠性，使客户与服务器应用之间的通信不被攻击者窃听。

### 1.TLS历史

```
1994年早期，NetScape公司设计了SSL协议（Secure Sockets Layer）的1.0版，但是未发布。
1994年11月，NetScape公司发布SSL 2.0版，很快发现有严重漏洞。
1996年11月，SSL 3.0版问世，得到大规模应用。
1999年1月，互联网标准化组织ISOC接替NetScape公司，发布了SSL的升级版TLS 1.0版。
2006年4月和2008年8月，TLS进行了两次升级，分别为TLS 1.1版和TLS 1.2版。最新的变动是2011年TLS 1.2的修订版。
现在正在制定 tls 1.3。
```

### 2.生成服务端证书

1. 生成服务端的私钥

```
openssl genrsa -out server.key 2048
```

2. 生成服务端证书

```
openssl req -new -x509 -key server.key -out server.pem -days 3650
```

### 3.生成客户端证书

1. 生成客户端的私钥

```
openssl genrsa -out client.key 2048
```

2. 生成客户端的证书

```
openssl req -new -x509 -key client.key -out client.pem -days 3650
```
