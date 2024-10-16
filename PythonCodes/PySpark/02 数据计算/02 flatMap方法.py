from pyspark import SparkConf, SparkContext
import os

# 为Spark添加python解释器
os.environ['PYSPARK_PYTHON'] = "D:/Python/Python310/python.exe"

conf = SparkConf().setMaster("local[*]").setAppName("test_spark")
sc = SparkContext(conf=conf)

rdd = sc.parallelize(["itheima abc", "itcast 6666", "python 123"])

# 功能需求,将rdd里的每个单体提取出来
rdd1 = rdd.map(lambda x: x.split(" "))
print(rdd1.collect())  # [['itheima', 'abc'], ['itcast', '6666'], ['python', '123']]

# 使用flatMap方法相比map能解除一层嵌套
rdd1 = rdd.flatMap(lambda x: x.split(" "))
print(rdd1.collect())  # ['itheima', 'abc', 'itcast', '6666', 'python', '123']
