# Mysql基础

[TOC]

**注释：**

> - 单行注释：--内容--     #内容
> - 多行注释：/ * 内容 */

## 通用语法及分类

- DDL: 数据定义语言，用来定义数据库对象（数据库、表、字段）
- DML: 数据操作语言，用来对数据库表中的数据进行增删改
- DQL: 数据查询语言，用来查询数据库中表的记录
- DCL: 数据控制语言，用来创建数据库用户、控制数据库的控制权限

## 一、DDL

- **数据定义语言**

### 1、数据库操作

1. 查询

   - 查询<font color='red'>所有</font>数据库

     > show databases;

   - 查询<font color='red'>当前</font>数据库

     > select database();

2. 创建

   > <font color='red'>create</font> database [if not exists] 数据库名 [default charset 字符集] [collate 排序规则]

3. 删除

   > <font color='red'>drop</font> database [if exists] 数据库名;

4. 使用

   > <font color='red'>use</font> 数据库名;
   
   ##### 注意事项
   
   - UTF8字符集长度为3字节，有些符号占4字节，所以推荐用utf8mb4字符集

### 2、表操作——查询

1. 查询当前数据库<font color='red'>所有表</font>

   > <font color='red'>show</font> tables;

2. 查询<font color='red'>表结构</font>

   > <font color='red'>desc</font> 表名

3. 查询指定表的<font color='red'>建表语句</font>

   > show create table 表名

4. 创建表

   > create table 表名(
   >
   > ​		字段1	字段1类型  [comment   '字段1注释']，
   >
   > ​		字段2	字段2类型  [comment  '字段2注释']，
   >
   > ​		....
   >
   > ​		字段n	字段n类型  [comment  '字段n注释']，
   >
   > )[comment '表注释']

   **最后一个字段后面没有逗号**
   
   ```mysql
   CREATE TABLE user(
   	id int comment '编号',
   	name varchar(50) comment '姓名',
   	age int comment '年龄',
   	gender varchar(1) comment '性别'
   )comment '用户表';
   ```

### 3、表操作——修改

1. 添加字段：

   > <font color='red'>ALTER</font>  <font color='red'>TABLE</font> 表名 <font color='red'>ADD</font> 字段名 类型(长度) [COMMENT 注释] [约束];
   > 				例：ALTER TABLE khsbase ADD nickname varchar(20) COMMENT '昵称';

2. 修改数据类型：

   > <font color='red'>ALTER TABLE</font> 表名 <font color='red'>MODIFY</font> 字段名 新数据类型(长度);

3. 修改字段名和字段类型：

   > <font color='red'>ALTER TABLE</font> 表名<font color='red'> CHANGE </font>旧字段名 新字段名 类型(长度) [COMMENT 注释] [约束];
   > 		例： ALTER TABLE emp CHANGE nickname username varchar(30) COMMENT '昵称';

4. 删除字段：

   > <font color='red'>ALTER TABLE</font> 表名<font color='red'> DROP</font> 字段名;

5. 修改表名：

   > <font color='red'>ALTER TABLE </font>表名 <font color='red'>RENAME TO</font> 新表名

6. 删除表：

   > <font color='red'>DROP TABLE</font> [IF EXISTS] 表名;

7. 删除表，并重新创建该表：

   > <font color='red'>TRUNCATE TABLE</font> 表名;

## 二、DML

- **数据增删改语句**

### 1. 添加数据

- **指定字段：**

> <font color='red'>INSERT INTO</font> 表名 <font color='red'>(字段名1, 字段名2, ...)</font> <font color='red'>VALUES</font> (值1, 值2, ...);

- **全部字段：**

> <font color='red'>INSERT INTO</font> 表名 <font color='red'>VALUES</font> (值1, 值2, ...);

- **批量添加数据：**

> <font color='red'>INSERT INTO</font> 表名 <font color='red'>(字段名1, 字段名2, ...) VALUES</font> (值1, 值2, ...), (值1, 值2, ...), (值1, 值2, ...);
>
> <font color='red'>INSERT INTO</font> 表名 <font color='red'>VALUES</font> (值1, 值2, ...), (值1, 值2, ...), (值1, 值2, ...);

#### 注意事项

- <font color='red'>字符串</font>和<font color='red'>日期类型</font>数据应该包含在<font color='red'>引号</font>中
- 插入的数据大小应该在字段的规定范围内

### 2. 更新和删除数据

- **修改数据：**

> <font color='red'>UPDATE</font> 表名 <font color='red'>SET</font> 字段名1 = 值1, 字段名2 = 值2, ... [ <font color='red'>WHERE</font> 条件 ];
> 	例：
> 	UPDATE emp SET name = 'Jack' WHERE id = 1;

- **删除数据：**

> <font color='red'>DELETE FROM</font> 表名 [ <font color='red'>WHERE</font> 条件 ];

## 三、DQL：

- **数据查询语句**

语法：

```mysql
SELECT
      字段列表
FROM
      表名字段
WHERE
      条件列表
GROUP BY
      分组字段列表
HAVING
      分组后的条件列表
ORDER BY
      排序字段列表
LIMIT
      分页参数
```

### 1. 基础查询

- **查询多个字段：**

> <font color='red'>SELECT</font> 字段1, 字段2, 字段3, ... <font color='red'>FROM</font> 表名;
>
> <font color='red'>SELECT</font>  *  <font color='red'>FROM</font> 表名;

- **设置别名：**

> <font color='red'>SELECT</font> 字段1 [ <font color='red'>AS</font> 别名1 ], 字段2 [ AS 别名2 ], 字段3 [ AS 别名3 ], ...<font color='red'> FROM</font> 表名;
> <font color='red'>SELECT</font> 字段1 [ 别名1 ], 字段2 [ 别名2 ], 字段3 [ 别名3 ], ... <font color='red'>FROM</font> 表名;

- **去除重复记录：**

> <font color='red'>SELECT DISTINCT</font> 字段列表 <font color='red'>FROM</font> 表名;

**转义：**

> `SELECT * FROM 表名 WHERE name LIKE '/_张三' ESCAPE '/'`
> / 之后的_不作为通配符

### 2. 条件查询

- **语法：**

> `SELECT 字段列表 FROM 表名 WHERE 条件列表;`

- **条件：**

| 比较运算符      | 功能                                                         |
| :-------------- | :----------------------------------------------------------- |
| >               | 大于                                                         |
| >=              | 大于等于                                                     |
| <               | 小于                                                         |
| <=              | 小于等于                                                     |
| =               | 等于                                                         |
| <> 或 !=        | 不等于                                                       |
| BETWEEN … AND … | 在某个范围内（含最小、最大值）                               |
| IN(…)           | 在in之后的列表中的值，多选一                                 |
| LIKE '占位符'   | <font color='red'>模糊匹配</font>（ _匹配单个字符，%匹配任意个字符） |
| IS NULL         | 是NULL                                                       |

| 逻辑运算符 | 功能                         |
| :--------- | :--------------------------- |
| AND 或 &&  | 并且（多个条件同时成立）     |
| OR 或 \|\| | 或者（多个条件任意一个成立） |
| NOT 或 !   | 非，不是                     |

**例子：**

```mysql
-- 年龄等于30
	select * from employee where age = 30;
-- 年龄小于30
	select * from employee where age < 30;
-- 小于等于
	select * from employee where age <= 30;
-- 没有身份证
	select * from employee where idcard is null or idcard = '';
-- 有身份证
	select * from employee where idcard;
	select * from employee where idcard is not null;
-- 不等于
	select * from employee where age != 30;
-- 年龄在20到30之间
	select * from employee where age between 20 and 30;
	select * from employee where age >= 20 and age <= 30;
-- 下面语句不报错，但查不到任何信息
	select * from employee where age between 30 and 20;
-- 性别为女且年龄小于30
	select * from employee where age < 30 and gender = '女';
-- 年龄等于25或30或35
	select * from employee where age = 25 or age = 30 or age = 35;
	select * from employee where age in (25, 30, 35);
-- 姓名为两个字
	select * from employee where name like '__';
-- 身份证最后为X
	select * from employee where idcard like '%X';
```

### 3. 聚合函数

- **常见聚合函数：**

| 函数  | 功能     |
| :---- | :------- |
| count | 统计数量 |
| max   | 最大值   |
| min   | 最小值   |
| avg   | 平均值   |
| sum   | 求和     |

- **语法：**

> <font color='red'>SELECT  聚合函数</font>(字段列表)  <font color='red'>FROM</font>  表名;
> 例：
> 	SELECT count(id) from employee where workaddress = "广东省";



### 4. 分组查询

- **语法：**

> <font color='red'>SELECT</font>  字段列表  <font color='red'>FROM</font>  表名  [ <font color='red'>WHERE</font> 条件 ]  <font color='red'>GROUP BY</font>  分组字段名  [ <font color='red'>HAVING </font>分组后的过滤条件 ];

**where 和 having 的<font color='red'>区别</font>：**

> - 执行时机不同：where是分组之前进行过滤，不满足where条件不参与分组；having是分组后<font color='red'>对结果进行过滤</font>。
> - 判断条件不同：where不能对聚合函数进行判断，而having可以。

**例子：**

```mysql
--  根据性别分组，统计男性和女性数量（只显示分组数量，不显示哪个是男哪个是女）
	select count(*) from employee group by gender;
-- 根据性别分组，统计男性和女性数量
	select gender, count(*) from employee group by gender;
-- 根据性别分组，统计男性和女性的平均年龄
	select gender, avg(age) from employee group by gender;
-- 年龄小于45，并根据工作地址分组
	select workaddress, count(*) from employee where age < 45 group by workaddress;
-- 年龄小于45，并根据工作地址分组，获取员工数量大于等于3的工作地址(别名为address_count)
	select workaddress, count(*) address_count from employee where age < 45 group by workaddress having 		 	 address_count >= 3;
```

#### 注意事项

> - 执行顺序：where > 聚合函数 > having
> - 分组之后，查询的字段一般为<font color='red'>聚合函数和分组字段</font>，查询其他字段无任何意义

### 5. 排序查询

- **语法：**

> <font color='red'>SELECT</font> 字段列表 <font color='red'>FROM</font> 表名 <font color='red'>ORDER BY</font> 字段1 排序方式1, 字段2 排序方式2;

- **排序方式：**

> - <font color='red'>ASC</font>:  升序（默认）
> - <font color='red'>DESC</font>:  降序

**例子：**

```mysql
-- 根据年龄升序排序
	SELECT * FROM employee ORDER BY age ASC;
	SELECT * FROM employee ORDER BY age;
-- 两字段排序，根据年龄升序排序，入职时间降序排序
	SELECT * FROM employee ORDER BY age ASC, entrydate DESC;
```

#### 注意事项

> 如果是<font color='red'>多字段排序</font>，当第一个字段值相同时，才会根据第二个字段进行排序

### 6. 分页查询

- **语法：**

> <font color='red'>SELECT</font> 字段列表 <font color='red'>FROM</font> 表名 <font color='red'>LIMIT</font> 起始索引, 查询记录数;

- **例子：**

```mysql
-- 查询第一页数据，展示10条
	SELECT * FROM employee LIMIT 0, 10;
-- 查询第二页
	SELECT * FROM employee LIMIT 10, 10;
```

#### 注意事项

> - 起始索引从0开始，<font color='red'>起始索引 = （查询页码 - 1） * 每页显示记录数</font>
> - 分页查询是数据库的方言，不同数据库有不同实现，MySQL是<font color='red'>LIMIT</font>
> - 如果查询的是第一页数据，起始索引可以省略，直接简写 LIMIT 10

#### DQL执行顺序

> FROM -> WHERE -> GROUP BY -> SELECT -> ORDER BY -> LIMIT

## 四、DCL:

### 1. 管理用户

#### **1.1 查询用户：**

> USE mysql;
>
> <font color='red'>SELECT</font>  *  FROM user;

#### **1.2 创建用户:**

> <font color='red'>CREATE  USER  '用户名'@'主机名'  IDENTIFIED  BY  '密码';</font>

#### **1.3 修改用户密码：**

> <font color='red'>ALTER USER '用户名'@'主机名' IDENTIFIED WITH mysql_native_password BY '新密码';</font>

#### **1.4删除用户：**

> `DROP USER '用户名'@'主机名';`

**例子：**

```mysql
-- 创建用户test，只能在当前主机localhost访问
	create user 'test'@'localhost' identified by '123456';
-- 创建用户test，能在任意主机访问
	create user 'test'@'%' identified by '123456';
	或create user 'test' identified by '123456';
-- 修改密码
	alter user 'test'@'localhost' identified with mysql_native_password by '1234';
-- 删除用户
	drop user 'test'@'localhost';
```

##### 注意事项

> - 主机名可以使用<font color='red'> % 通配</font>

### 2. 权限控制

- 常用权限：

| 权限                | 说明               |
| :------------------ | :----------------- |
| ALL, ALL PRIVILEGES | 所有权限           |
| SELECT              | 查询数据           |
| INSERT              | 插入数据           |
| UPDATE              | 修改数据           |
| DELETE              | 删除数据           |
| ALTER               | 修改表             |
| DROP                | 删除数据库/表/视图 |
| CREATE              | 创建数据库/表      |

- 更多权限请看[权限一览表](https://dhc.pythonanywhere.com/article/public/1/#权限一览表)

	

#### 2.1 查询权限：

> `SHOW GRANTS FOR '用户名'@'主机名';`

#### 2.2 授予权限：

> `GRANT 权限列表 ON 数据库名.表名 TO '用户名'@'主机名';`

#### 2.3 撤销权限：

> `REVOKE 权限列表 ON 数据库名.表名 FROM '用户名'@'主机名';`

##### 注意事项

- 多个权限用<font color='red'>逗号</font>分隔
- 授权时，<font color='red'>数据库名和表名可以用 * 进行通配</font>，代表所有

## 五、事务

事务是<font color='red'>一组操作</font>的集合，事务会把所有操作作为一个整体一起向系统提交或撤销操作请求，即这些操作<font color='red'>要么同时成功，要么同时失败</font>。

- **基本操作：**

```mysql
-- 1. 查询张三账户余额
	select * from account where name = '张三';
-- 2. 将张三账户余额-1000
	update account set money = money - 1000 where name = '张三';
-- 此语句出错后张三钱减少但是李四钱没有增加
	模拟sql语句错误
-- 3. 将李四账户余额+1000
	update account set money = money + 1000 where name = '李四';
-- 查看事务提交方式
	SELECT @@AUTOCOMMIT;
-- 设置事务提交方式，1为自动提交，0为手动提交，该设置只对当前会话有效
	SET @@AUTOCOMMIT = 0;
-- 提交事务
	COMMIT;
-- 回滚事务
	ROLLBACK;
-- 设置手动提交后上面代码改为：
	select * from account where name = '张三';
	update account set money = money - 1000 where name = '张三';
	update account set money = money + 1000 where name = '李四';
	commit;
```

- **操作方式二：**

**开启事务：**

> <font color='red'>START</font> TRANSACTION 或 <font color='red'>BEGIN</font> TRANSACTION;

**提交事务：**

> <font color='red'>COMMIT</font>;

**回滚事务：**

> <font color='red'>ROLLBACK</font>;

**操作实例：**

```mysql
start transaction;
select * from account where name = '张三';
update account set money = money - 1000 where name = '张三';
update account set money = money + 1000 where name = '李四';
commit;
```

### 四大特性ACID

- 原子性(Atomicity)：事务是不可分割的最小操作但愿，要么全部成功，要么全部失败
- 一致性(Consistency)：事务完成时，必须使所有数据都保持一致状态
- 隔离性(Isolation)：数据库系统提供的隔离机制，保证事务在不受外部并发操作影响的独立环境下运行
- 持久性(Durability)：事务一旦提交或回滚，它对数据库中的数据的改变就是永久的

### 并发事务

| 问题       | 描述                                                         |
| :--------- | :----------------------------------------------------------- |
| 脏读       | 一个事务读到另一个事务还没提交的数据                         |
| 不可重复读 | 一个事务先后读取同一条记录，但两次读取的数据不同             |
| 幻读       | 一个事务按照条件查询数据时，没有对应的数据行，但是再插入数据时，又发现这行数据已经存在 |

> 这三个问题的详细演示：https://www.bilibili.com/video/BV1Kr4y1i7ru?p=55cd

- **并发事务隔离级别：**

| 隔离级别              | 脏读 | 不可重复读 | 幻读 |
| :-------------------- | :--- | :--------- | :--- |
| Read uncommitted      | √    | √          | √    |
| Read committed        | ×    | √          | √    |
| Repeatable Read(默认) | ×    | ×          | √    |
| Serializable          | ×    | ×          | ×    |

- √ 表示在当前隔离级别下该问题<font color='red'>会出现</font>
- Serializable 性能最低；Read uncommitted 性能最高，数据安全性最差



- **查看事务隔离级别：**

> <font color='red'>SELECT @@TRANSACTION_ISOLATION</font>;

- **设置事务隔离级别：**

> <font color='red'>SET</font> [ SESSION | GLOBAL ] <font color='red'>TRANSACTION</font> <font color='red'>ISOLATION LEVEL</font> {READ UNCOMMITTED | READ COMMITTED | REPEATABLE READ | SERIALIZABLE };

SESSION 是会话级别，表示只针对当前会话有效，GLOBAL 表示对所有会话有效
