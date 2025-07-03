# Makefile命令

## proto 文件生成go文件
- make proto dir=admin_proto

# 项目

## 鉴权

### JWT
##### 公私钥生成
```shell
# 生成私钥
openssl genpkey -algorithm RSA -out private.pem -pkeyopt rsa_keygen_bits:2048

# 提取公钥
openssl rsa -pubout -in private.pem -out public.pem

```
### 配置秘钥字符串生成
1. 读取private.pem文件内容，并base64编码
2. 读取public.pem文件内容，并base64编码，配置到jwt配置项的public
