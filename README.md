# httpvf
http REST API verify

## get

```
go get github.com/toukii/httpvf/vf
```

## example


__vf.yml:__


```yml
-
  url: http://upload.daoapp.io/topic
  method: POST
  header:
    Content-Type: "application/x-www-form-urlencoded"
  body: "title=topic001&content=content001"
  resp:
    code: 200
    cost: 900
-
  url: http://upload.daoapp.io/loadfile/topic001
  method: GET
  resp:
    code: 200
    cost: 80
    body: content001
-
  url: http://upload.daoapp.io/uploadform/
  method: POST
  filename : vf.yml
  resp:
    code: 200
    cost: 10
```


verify:

```
vf -v vf.yml
```

## 请求

 - url: [请求地址]
 
 - method: [请求方法]
 
 - body: [请求body]
 
 - n: [请求个数]
 
 - interval: [请求间隔]
 
 - runtine: [请求并发数]
 
 - upload: [web前段传入的文件名（input name）]@[上传文件名]

 - header: 请求的header参数，map结构
 
```
header:
   Content-Type: "application/x-www-form-urlencoded"
```
 
 - param: GET 请求的参数，map结构

```
param:
  name: toukii
  position: dev-ops
```

## 验证返回body

 - code: [响应码]
 
 - cost: [响应时间，单位ms]
 
 - body: [直接验证内容]
 
 - regex: [正则表达式]


### json

 - 路径以","分割

 - 路径若有纯数字，为数字加上""
 
 - 数组下标从0开始，直接写数字
 
 例如，返回的json内容如下：
 
```json
[
    {
        "Map": {
            "1": "hello"
        },
        "Message": "This is toukii,r1",
        "Cost": 0.315
    }
]
```

验证hello的写法为：

```yml
json: 
  '0,Map,"1"': hello
```