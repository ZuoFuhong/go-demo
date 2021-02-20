## RSA非对称加密算法：公钥用于加密，私钥用于解密

1. 生成私钥

```
openssl genrsa -out private.key 2048
```

2. 生成公钥

```
openssl rsa -in private.key -pubout -out public.key
```