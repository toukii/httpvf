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
