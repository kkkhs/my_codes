from pyspark import SparkConf, SparkContext
import os

# 为Spark添加python解释器
os.environ['PYSPARK_PYTHON'] = "D:/Python/Python310/python.exe"

conf = SparkConf().setMaster("local[*]").setAppName("test_spark")
sc = SparkContext(conf=conf)

rdd = sc.parallelize([('China', 2), ('is', 6), ('love', 123), ('my', 1233), ('motherland', 444), ('I', 566)])
# sortBy(func, ascending=False, numPartitions=1) -> （用哪个数据排序，升序/降序， 多少分区排序）
rdd = rdd.sortBy(lambda x: x[1], ascending=True, numPartitions=1)
print(rdd.collect())
