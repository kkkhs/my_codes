from pymysql import Connection

con = Connection(
    host="localhost",   # 主机名
    port=3306,          # 端口
    user="root",        # 账户
    password="123456",  # 密码
    autocommit=True     # 修改数据自动确认
)

all_data = [[]]

cursor = con.cursor()
con.select_db("py_sql")
for record in all_data:
    sql = f"insert into orders(order_date, order_id, money, province) values('{record.date}', '{record.order_id}', '{record.money}', '{record.province}')"
    cursor.execute(sql)

con.close()