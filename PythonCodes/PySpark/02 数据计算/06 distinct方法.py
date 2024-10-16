from pyspark import SparkConf, SparkContext
import os

# 为Spark添加python解释器
os.environ['PYSPARK_PYTHON'] = "D:/Python/Python310/python.exe"

conf = SparkConf().setMaster("local[*]").setAppName("test_spark")
sc = SparkContext(conf=conf)

rdd = sc.parallelize([1, 1, 3, 3, 5, 5, 7, 9, 9])

# distinct 去重
rdd = rdd.distinct()
print(rdd.collect())
