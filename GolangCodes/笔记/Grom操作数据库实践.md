# Grom操作数据库实践

`首先需要下载mysql的驱动`

```go
go get gorm.io/driver/mysql
go get gorm.io/gorm
```

##  一、简单连接

```go
username := "root"  //账号
password := "123456"  //密码
host := "127.0.0.1" //数据库地址，可以是Ip或者域名
port := 3306        //数据库端口
Dbname := "test"   //数据库名
timeout := "10s"    //连接超时，10秒

// root:root@tcp(127.0.0.1:3306)/gorm?
dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
DB, err := gorm.Open(mysql.Open(dsn))
if err != nil {
    panic("连接数据库失败, error=" + err.Error())
}
// 连接成功
fmt.Println(db)
```

**准备使用的表结构体：**

```go
type Student struct {
  ID     uint   `gorm:"size:3"`
  Name   string `gorm:"size:8"`
  Age    int    `gorm:"size:3"`
  Gender bool
  Email  *string `gorm:"size:32"`
}
```

## 二、增加数据

首先使用 ` DB.AutoMigrate(&Student{}) `在数据库自动生成对应的表结构

初始表students为空:

![img](file:///C:\Users\KHS\AppData\Roaming\Tencent\Users\2964745405\QQ\WinTemp\RichOle\_TO{J9VYZ_CZG}1[ESP~O9C.png)

### 1、单条数据插入

```go
email := "xxx@qq.com"
// 创建记录
student := Student{
  Name:   "khs",
  Age:    19,
  Gender: true,
  Email:  &email,
}
db.Create(&student)
```

有两个地方需要**注意**

> 1. `指针类型`是为了更好的`存null类型`，但是传值的时候，也记得传指针
> 2. Create接收的是一个`指针`，而不是值

由于我们传递的是一个指针，调用完Create之后，student这个对象上面就有该记录的信息了，如创建的id:

```go
DB.Create(&student)
fmt.Printf("%#v\n", student)  
// main.Student{ID:0x2, Name:"khs", Age:19, Gender:false, Email:(*string)(0x11d40980)}
```

### 2、多条数据插入

```go
var studentList []Student
for i := 0; i < 100; i++ {
    email := fmt.Sprintf("%d@qq.com", i+1)
  	studentList = append(studentList, Student{
    	Name:   fmt.Sprintf("机器人%d号", i+1),
    	Age:    21,
   		Gender: true,
    	Email:  &email,
  	})
}
db.Create(&studentList)
```

## 三、查询数据

### 1、查询单条记录

```go
var student Student
DB.Take(&student)
fmt.Println(student)
```

获取单条记录的方法很多，我们对比sql就很直观了

```go
DB = DB.Session(&gorm.Session{Logger: Log})
var student Student
DB.Take(&student)  
// SELECT * FROM `students` LIMIT 1
DB.First(&student) 
// SELECT * FROM `students` ORDER BY `students`.`id` LIMIT 1
DB.Last(&student)  
// SELECT * FROM `students` ORDER BY `students`.`id` DESC LIMIT 1
```

#### 1.1 根据主键查询

```go
var student Student
DB.Take(&student, 2)
fmt.Println(student)

student = Student{} // 重新赋值
DB.Take(&student, "4")
fmt.Println(student)
```

Take的第二个参数，默认会根据主键查询，`可以是字符串，可以是数字`

#### 2.2 根据其他条件查询

```go
var student Student
DB.Take(&student, "name = ?", "机器人27号")
fmt.Println(student)
```

`使用？作为占位符，将查询的内容放入?`

```go
//同理于mysql语句:
SELECT * FROM `students` WHERE name = '机器人27号' LIMIT 1
```

这样可以有效的`防止sql注入`

他的原理就是将参数全部转义，如

```go
DB.Take(&student, "name = ?", "机器人27号' or 1=1;#")

SELECT * FROM `students` WHERE name = '机器人27号\' or 1=1;#' LIMIT 1
```

#### 1.3 根据struct查询

```go
var student Student
// 只能有一个主要值
student.ID = 2
//student.Name = "枫枫"
DB.Take(&student)
fmt.Println(student)
```

#### 1.4 获取查询结果

- **获取查询的记录数**

```go
count := DB.Find(&studentList).RowsAffected
```

- **是否查询失败**

```go
err := DB.Find(&studentList).Error
```

查询失败原因有：`查询为空，查询条件错误，sql语法错误`

- 可以使用判断

```go
var student Student
err := DB.Take(&student, "xx").Error
switch err {
	case gorm.ErrRecordNotFound:
  		fmt.Println("没有找到")
	default:
  		fmt.Println("sql错误")
}
```

### 2、查询多条记录

```go
var studentList []Student
DB.Find(&studentList)
for _, student := range studentList {
  	fmt.Println(student)
}

// 由于email是指针类型，所以看不到实际的内容
// 但是序列化之后，会转换为我们可以看得懂的方式
var studentList []Student
DB.Find(&studentList)
for _, student := range studentList {
  	data, _ := json.Marshal(student)
  	fmt.Println(string(data))
}
```

#### 2.1 根据主键列表查询

```go
var studentList []Student
DB.Find(&studentList, []int{1, 3, 5, 7})
DB.Find(&studentList, 1, 3, 5, 7)  // 一样的
fmt.Println(studentList)
```

#### 2.2 根据其他条件查询

```go
DB.Find(&studentList, "name in ?", []string{"khs", "zhangsan"})
```

## 四、修改数据

修改的前提的`先查询到记录`

### 1、Save保存所有字段

用于单个记录的`全字段更新`

它会保存所有字段，即使零值也会保存

```go
var student Student
DB.Take(&student)
student.Age = 23
// 全字段更新
DB.Save(&student)
// UPDATE `students` SET `name`='枫枫',`age`=23,`gender`=true,`email`='xxx@qq.com' WHERE `id` = 1
```

零值也会更新

```go
var student Student
DB.Take(&student)
student.Age = 0
// 全字段更新
DB.Save(&student)
// UPDATE `students` SET `name`='枫枫',`age`=0,`gender`=true,`email`='xxx@qq.com' WHERE `id` = 1
```

### 2 更新指定字段

可以使用select选择要更新的字段

```go
var student Student
DB.Take(&student)
student.Age = 21
// 全字段更新
DB.Select("age").Save(&student)
// UPDATE `students` SET `age`=21 WHERE `id` = 1
```

### 3 批量更新

例如我想给年龄21的学生，都更新一下邮箱

```go
var studentList []Student
DB.Find(&studentList, "age = ?", 21).Update("email", "is21@qq.com")
```

还有一种更简单的方式:

```go
DB.Model(&Student{}).Where("age = ?", 21).Update("email", "is21@qq.com")
// UPDATE `students` SET `email`='is22@qq.com' WHERE age = 21CopyErrorOK!
```

这样的更新方式也是可以更新零值的

### 4 更新多列

如果是结构体，它默认不会更新零值

```go
email := "xxx@qq.com"
DB.Model(&Student{}).Where("age = ?", 21).Updates(Student{
  Email:  &email,
  Gender: false,  // 这个不会更新
})

// UPDATE `students` SET `email`='xxx@qq.com' WHERE age = 21
```

如果想让他更新零值，用select就好

```go
email := "xxx1@qq.com"
DB.Model(&Student{}).Where("age = ?", 21).Select("gender", "email").Updates(Student{
  Email:  &email,
  Gender: false,
})
// UPDATE `students` SET `gender`=false,`email`='xxx1@qq.com' WHERE age = 21CopyErrorOK!
```

如果不想多写几行代码，则推荐`使用map`

```go
DB.Model(&Student{}).Where("age = ?", 21).Updates(map[string]any{
  "email":  &email,
  "gender": false,
})
```

### 5 更新选定字段

```go
Select选定字段
Omit忽略字段
```

## 五、删除数据

### 1、根据结构体删除

```go
// student 的 ID 是 `10`
db.Delete(&student)
// DELETE from students where id = 10;
```

### 2、删除多个数据

```go
db.Delete(&Student{}, []int{1,2,3})

// 查询到的切片列表
db.Find(&studentList)
db.Delete(&studentList) //全部删除
```

总结:

> 

