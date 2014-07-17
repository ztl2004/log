## Email API Specification


### API Specification


#### Email API Specification


##### POST /v1/:app/log

记录 application 完整流程的全部日志,由系统的各个模块调用

###### Examole Request
```
POST /v1/233/log HTTP/1.1
Host: log.arkors.com
X-Arkors-application-Log: 5024442115e7bd738354c1fac662aed5
X-Arkors-application-Client: xxx.xx.xx.xxx,Email
Accept: application/json
{
  "level":"debug"
  "aciton":"insert into mysql finished"
  "parent":"5024442115e7bd738354c1fac662aed4"
}

```
###### Example Response
```
HTTP/1.1 201 OK
X-Arkors-application-Log: cb21df532c6647383af7efa0fd8405f2
Content-Type: application/json
{
  "app": 233,
  "id": 1
}
```
