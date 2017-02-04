# httpvf
http REST API verify

## get

```
go get github.com/toukii/httpvf/vf
```

## example

__test.yaml:__

```yaml
url: http://upload.daoapp.io/uploadform/
method: POST
body: '{"name":"toukii"}'
filename : vf.yaml
resp:
  code: 200
  cost: 10
  body: success
```


verify:

```
vf -v test.yaml
```

__[more verify] vf.yaml:__


```yaml

-
  url: http://upload.daoapp.io/loadfile/install.sh
  method: GET
  body: hello
  resp:
    code: 200
    cost: 80
    body: world

-
  url: http://upload.daoapp.io/loadfile/a
  method: POST
  body: hello
  resp:
    code: 403
    cost: 90
    body: world
-
  url: http://upload.daoapp.io/uploadform/
  method: POST
  body: '{"name":"toukii"}'
  filename : vf.yaml
  resp:
    code: 200
    cost: 10
    body: success
```
