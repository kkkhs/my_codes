def print_file_info(file_name):
    f = None
    try:
        f = open(file_name, 'r', encoding='UTF-8')
        content = f.read()
        print("文件的全部内容如下：")
        print(content)
    except Exception as e:
        print(f"文件出错，原因是：{e}")
    finally:
        if f:
            f.close()


def append_to_file(file_name, data):
    f = None

    f = open(file_name, 'a', encoding='UTF-8')
    f.write(data)
    f.write("\n")
    f.close()

    # 追加不用检查文件是否存在（不存在会创建）


