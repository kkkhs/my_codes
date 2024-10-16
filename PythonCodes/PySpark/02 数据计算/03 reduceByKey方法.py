from pyspark import SparkConf, SparkContext
import os

# 为Spark添加python解释器
os.environ['PYSPARK_PYTHON'] = "D:/Python/Python310/python.exe"

conf = SparkConf().setMaster("local[*]").setAppName("test_spark")
sc = SparkContext(conf=conf)

# reduceByKey方法:将元组根据key计算同一个Key的Value
rdd = sc.parallelize([('a', 1), ('b', 2), ('b', 1), ('a', 5)])

# 用法：reduceByKey(func)  func(V, V) -> V
result = rdd.reduceByKey(lambda x, y: x + y)
print(result.collect())
