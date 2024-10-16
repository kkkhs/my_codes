from pyspark import SparkConf, SparkContext
import os

os.environ["PYSPARK_PYTHON"] = 'D:\\Python\\Python310\\python.exe'
os.environ['HADOOP_HOME'] = 'D:\\hadoop-3.1.0'
conf = SparkConf().setMaster('local[*]').setAppName("spark_test")
sc = SparkContext(conf=conf)

f = sc.textFile(".//search_log.txt", 1)
# 1、
text_rdd = f.map(lambda x: x.split('\t'))
time_rdd = text_rdd.map(lambda x: x[0].split(":"))
time_tuple = time_rdd.map(lambda x: (x[0], 1))
result1 = time_tuple.reduceByKey(lambda a, b: a + b).sortBy(lambda x: x[1], ascending=False).take(3)
print(result1)
# 2、
word_rdd = text_rdd.map(lambda x: x[2])
word_tuple = word_rdd.map(lambda x: (x, 1))
word_add = word_tuple.reduceByKey(lambda a, b: a + b)
word_result = word_add.sortBy(lambda x: x[1], ascending=False, numPartitions=1)
result12 = word_result.take(3)
print(result12)
# 3、
text_rdd.filter(lambda x: x[2] == "黑马程序员")
time_rdd = text_rdd.map(lambda x: x[0].split(":"))
time_tuple = time_rdd.map(lambda x: (x[0], 1))
result1 = time_tuple.reduceByKey(lambda a, b: a + b)
result3 = result1.take(1)
print(result3)

text_rdd.map(lambda x: {"time": x[0], "keyword": x[2]}).saveAsTextFile('2')
