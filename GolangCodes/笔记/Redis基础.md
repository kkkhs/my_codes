# Redis基础

[TOC]

## 一、Redis基本使用

> <font color='red'>set</font> key value
>
> get key
>

## 二、对String的操作

键值对 key-value

> 添加 set key value
>
> 获取 get key
>
> 删除 <font color='red'>del</font> key

> 设置存活时间 <font color='red'>setex </font>key time value

> 一次性设置多个 <font color='red'>mset</font> key value [key value]...
>
> 一次性获取多个 <font color='red'>mget</font> key [key]...

## 三、对Hash的基本使用

适用于结构体的存储<font color='red'> key不能重复</font>

> 添加 <font color='red'>hset</font> key field value   如：hset user1 name "smith"
>
> 获取 <font color='red'>hget</font> key fild   如：hget user1 name
>
> 一次性获取全部 <font color='red'>hgetall</font> key
>
> 删除 <font color='red'>hdel</font> key file   如：hdel user1 job

**使用细节：**

> 一次性设置多个 <font color='red'>hmset</font> key field [key field]...
>
> 一次性获取多个 <font color='red'>hmget</font> key [key]...
>
> 查看字段大小 <font color='red'>hlen</font> key
>
> 判断key中某一field是否存在 <font color='red'>hexists</font> key feild 

## 三、对List的基本使用

按照插入的<font color='red'>顺序</font>排序 可以添加一个元素到列表的头部或者尾部，本质是<font color='red'>链表</font>

> 左边插入数据 <font color='red'>lpush</font> key value [value]...   如：lpush city beijing shanghai tianjing
>
> 右边插入数据 <font color='red'>rpush</font> key value [value]...
>
> 区间获取数据 <font color='red'>lrange</font> key start stop   如：lange city 0 -1
>
> 删除列表 <font color='red'>del</font> key

> 左边弹出数据 <font color='red'>lpop</font> key
>
> 右边弹出数据 <font color='red'>rpop </font>key

**使用细节：**

> - <font color='red'>index</font>,按照所有下标获得元素(从左到右，从0开始)
>
> - 获取列表长度 <font color='red'>llen</font> key
> - 如果列表的值全移除，对应的键也就消失了

## 四、对Set的基本使用

<font color='red'>无序</font>集合，底层时HashTable, 且<font color='red'>元素的值不能重复</font>

> 添加值 <font color='red'>sadd</font> key value [value]...
>
> 取出所有值 <font color='red'>smembers</font> key
>
> 判断值是否是成员 <font color='red'>sismember</font> key
>
> 删除指定值 <font color='red'>srem</font> key value