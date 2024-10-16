from pymysql import Connection

con = Connection(
    host="localhost",   # 主机名
    port=3306,          # 端口
    user="root",        # 账户
    password="123456",  # 密码
    autocommit=True     # 修改数据自动确认
)

print(con.get_server_info())

# 执行非查询sql语句
# 获取游标对象
cursor = con.cursor()
# 选择数据库
con.select_db("demo")
# 执行sql
# cursor.execute("create table khs(id int)")

# 执行查询sql语句
cursor.execute("select * from student")

results = cursor.fetchall()
for r in results:
    print(r)

# 数据插入
cursor.execute("insert into student values(10002, '李连杰', 33)")
# 手动确认 commit
con.commit()

# 关闭连接
con.close()
