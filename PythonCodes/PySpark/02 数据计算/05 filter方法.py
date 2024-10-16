from pyspark import SparkConf, SparkContext
import os

# 为Spark添加python解释器
os.environ['PYSPARK_PYTHON'] = "D:/Python/Python310/python.exe"

conf = SparkConf().setMaster("local[*]").setAppName("test_spark")
sc = SparkContext(conf=conf)

rdd = sc.parallelize([1, 2, 3, 4, 5])

# filter 过滤 接受一个函数返回为true则保留
rdd = rdd.filter(lambda x: (x % 2 == 0))
print(rdd.collect())


