from pyspark import SparkConf, SparkContext
import os

# 为Spark添加python解释器
os.environ['PYSPARK_PYTHON'] = "D:/Python/Python310/python.exe"

conf = SparkConf().setMaster("local[*]").setAppName("test_spark")
sc = SparkContext(conf=conf)
rdd = sc.parallelize([1, 2, 3, 4, 5])

# 通过map方法将全部数据都乘以10 再+5
rdd1 = rdd.map(lambda x: x * 10).map(lambda x: x+5)
print(rdd1.collect())
