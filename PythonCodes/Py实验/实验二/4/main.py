import random
a_list = [x for x in range(10)]
print(f"所有元素的新列表: {a_list[::2]}")
print(f"具有偶数位置的元素列表: {a_list[::2]}")
a_list.reverse()
print(f"所有元素的逆序列表: {a_list[::2]}")
