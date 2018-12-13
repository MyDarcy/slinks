# 短链接生成系统slinks |说明文档

---

## 主要功能
1. 输入长链接，输出短链接
2. 根据短链接，查询长链接
3. 根据短连接，跳转长链接

## 接口协议格式
POST http://s.xawei.me/v1/genlink  生成短链接
```
入参：
{
    "llink":"http://xxx.long link.xxx" //长链接
}
返回：
{
  "code": 1, //1成功，-1失败
  "data": {
    "llink": "http://xxx.long link.xxx", //长链接
    "slink": "http://s.xawei.me/s/xxx"      //对应长链接生成的短链接
  },
  "msg": "gen slink success"            //回包说明
}
```


----------

POST http://s.xawei.me/v1/genlink  查询短链接对应的长链接
```
入参：
{
    "slink":"http://s.xawei.me/s/xxx" //短链接
}
返回
{
  "code": 1,        //1成功，-1失败
  "data": {
    "Id": x,
    "Slink": "xxx",                     //短链接后缀
    "Llink": "http://xawei.me/about/"  //短链接s.xawei.me/s/xxx对应的长链接
  },
  "msg": "get link success"
}
```


----------

GET http://s.xawei.me/s/xxx  短链接跳转长链接
```
根据短链接直接跳转至对应的长链接。
```



## 示例
假设有长链接：
http://xawei.me/2017/07/30/%E5%8D%95%E5%AE%9E%E4%BE%8BRedis%E6%9E%84%E5%BB%BA%E5%88%86%E5%B8%83%E5%BC%8F%E9%94%81%E4%B8%AD%E7%9A%84ABA%E9%97%AE%E9%A2%98/
希望转换成短链接，则可调用接口http://s.xawei.me/v1/genlink 如下：
```
POST http://s.xawei.me/v1/genlink
参数：
{
    "llink":"http://xawei.me/2017/07/30/%E5%8D%95%E5%AE%9E%E4%BE%8BRedis%E6%9E%84%E5%BB%BA%E5%88%86%E5%B8%83%E5%BC%8F%E9%94%81%E4%B8%AD%E7%9A%84ABA%E9%97%AE%E9%A2%98/" //长链接
}
```

也可以直接使用以下curl命令：
```
curl --request POST \
  --url http://s.xawei.me/v1/genslink \
  --header 'content-type: application/json' \
  --data '{
	"llink": "http://xawei.me/2017/07/30/%E5%8D%95%E5%AE%9E%E4%BE%8BRedis%E6%9E%84%E5%BB%BA%E5%88%86%E5%B8%83%E5%BC%8F%E9%94%81%E4%B8%AD%E7%9A%84ABA%E9%97%AE%E9%A2%98/"
}'
```

得到返回结果：
```
{"code":1,"data":{"llink":"http://xawei.me/2017/07/30/%E5%8D%95%E5%AE%9E%E4%BE%8BRedis%E6%9E%84%E5%BB%BA%E5%88%86%E5%B8%83%E5%BC%8F%E9%94%81%E4%B8%AD%E7%9A%84ABA%E9%97%AE%E9%A2%98/","slink":"http://s.xawei.me/s/1"},"msg":"gen slink success"}
```
则我们可以直接在浏览器访问
http://s.xawei.me/s/1
即可跳转至长链接
http://xawei.me/2017/07/30/%E5%8D%95%E5%AE%9E%E4%BE%8BRedis%E6%9E%84%E5%BB%BA%E5%88%86%E5%B8%83%E5%BC%8F%E9%94%81%E4%B8%AD%E7%9A%84ABA%E9%97%AE%E9%A2%98/


## 参考资料
1. How do I create a URL shortener
https://stackoverflow.com/questions/742013/how-do-i-create-a-url-shortener