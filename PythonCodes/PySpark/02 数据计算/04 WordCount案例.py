from pyspark import SparkConf, SparkContext
import os

# 为Spark添加python解释器
os.environ['PYSPARK_PYTHON'] = "D:/Python/Python310/python.exe"

conf = SparkConf().setMaster("local[*]").setAppName("test_spark")
sc = SparkContext(conf=conf)

# 1 读取文件数据
rdd = sc.textFile("words.txt")
print(rdd.collect())

# 2 取出单词
rdd = rdd.flatMap(lambda x: x.split(" "))

# 3 转变成元组：(单词, 1)
rdd = rdd.map(lambda x: (x, 1))
print(rdd.collect())

# 4 分组相加
rdd = rdd.reduceByKey(lambda a, b: a + b)
print(rdd.collect())

# 5 结果排序
rdd1 = rdd.sortBy(lambda x: x[1], ascending=True, numPartitions=1)
print(rdd1.collect())
