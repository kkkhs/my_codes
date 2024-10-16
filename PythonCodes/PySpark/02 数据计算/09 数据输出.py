from pyspark import SparkConf, SparkContext
import os

os.environ["PYSPARK_PYTHON"] = 'D:\\Python\\Python310\\python.exe'
os.environ['HADOOP_HOME'] = 'D:\\hadoop-3.1.0'
conf = SparkConf().setMaster('local[*]').setAppName("spark_test")
sc = SparkContext(conf=conf)

rdd = sc.parallelize([1, 2, 3, 4, 5], 1)

# collect算子，将rdd转换为list对象
rdd_list: list = rdd.collect()
print(rdd_list)
print(type(rdd_list))

# reduce算子，对RDD两两聚合
num = rdd.reduce(lambda a, b: a + b)
print(num)

# take算子, 取出RDD前N个元素， 组成LIST
take_list = rdd.take(3)
print(take_list)

# count算子， 计算rdd有多少数据
n = rdd.count()
print(n)

#  saveAsTextFile算子，输出到文件中
rdd.saveAsTextFile("1.txt")

sc.stop()
