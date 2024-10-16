import numpy as np
import pandas as pd
import matplotlib.pyplot as plt

# 数据
x = [22460, 11226, 34547, 4851, 5444, 2662, 4549]
y = [7326, 4490, 11546, 2396, 2208, 1608, 2035]

# 设置中文显示
plt.rcParams['font.sans-serif'] = ['SimHei']
plt.rcParams['font.size'] = 13

# 绘制散点图
plt.scatter(x, y)
plt.title('Scatter Plot')  # 设置标题
plt.xlabel('人均国内生产总值(GDP)')  # 设置x轴标签
plt.ylabel('人均消费水平')  # 设置y轴标签

# 计算相关系数和预测值
slope, intercept = np.polyfit(x, y, 1)
corr = np.corrcoef(x, y)[0, 1]
print('相关系数为:', corr)
sp = slope * 5000 + intercept
print('预测的人均消费水平为:', sp)

# 显示图形
plt.show()
