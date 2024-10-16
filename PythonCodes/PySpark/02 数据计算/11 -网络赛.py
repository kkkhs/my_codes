from pyspark import SparkConf, SparkContext
import os

os.environ["PYSPARK_PYTHON"] = 'D:\\Python\\Python310\\python.exe'
os.environ['HADOOP_HOME'] = 'D:\\hadoop-3.1.0'
conf = SparkConf().setMaster('local[*]').setAppName("spark_test")
sc = SparkContext(conf=conf)

rdd = sc.textFile('pri_user_adv.log', 1)

# 1、得到第一个元组(省份、广告、点击数量), 并将每个广告的点击数量相加
r1 = rdd.map(lambda x: x.split(" ")).\
    map(lambda x: ((x[1], x[4]), 1)).\
    reduceByKey(lambda a, b: a + b)

# 2、创建新的格式(省份, [(广告、点击数量)])， 将每个省份的广告合并，排序取前三
r2 = r1.map(lambda x: (x[0][0], [(x[0][1], x[1])])).\
    reduceByKey(lambda a, b: a + b).\
    map(lambda x: (x[0], sorted(x[1], key=lambda i: i[1], reverse=True)[:3]))

# 3、输出答案
for province, top in r2.collect():
    for i in top:
        print(province, i[0], i[1])

# 4、关闭sc
sc.stop()

