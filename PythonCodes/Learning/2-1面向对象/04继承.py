class phone:
    IMEI = None
    producer = "khs"

    def call_by_4G(self):
        print("正在4G通话")


class Phone2023(phone):
    face_id = "10001"

    def call_by_5G(self):
        print("正在5G通话")


phone1 = phone()
phone1.call_by_4G()
p2 = Phone2023()
p2.call_by_4G()
p2.call_by_5G()


class NFC:
    nfc_type = "第五代"
    producer = "khs"

    def read_card(self):
        print("NFC读卡")

    def write_card(self):
        print("NFC写卡")


class RemoteControl:
    rc_type = "红外遥控"

    def control(self):
        print("红外遥控开启！")


# 多继承
class MyPhone(phone, NFC, RemoteControl):
    # pass 表示类是空的  补全语法
    pass


mp = MyPhone()
mp.call_by_4G()
mp.control()
# 输出同名属性时，左边的优先级最高


# super 使用复写后的父类的方法和属性
