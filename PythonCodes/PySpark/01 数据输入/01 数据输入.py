from pyspark import SparkConf, SparkContext

# 创建SparkConf类对象
conf = SparkConf().setMaster("local[*]").setAppName("test_spark_app")

# 基于类对象创建SparkContext类对象
sc = SparkContext(conf=conf)

# 打印版本
print(sc.version)

#
# 1 通过parallelize方法将python对象加载到Spark内,成为RDD对象
rdd1 = sc.parallelize([1, 2, 3, 4, 5])
rdd2 = sc.parallelize((1, 2, 3, 4, 5))
rdd3 = sc.parallelize("abcdrfg")
rdd4 = sc.parallelize({1, 2, 3, 4, 5})
rdd5 = sc.parallelize({"key1": "value", "key2": "value2"})

# 使用collect()方法查看rdd内容
print(rdd1.collect())
print(rdd2.collect())
print(rdd3.collect())
print(rdd4.collect())
print(rdd5.collect())

# 2 通过textFile方法, 读取文件数据加载到Spark内, 成为RDD对象
rdd = sc.textFile("words.txt")
print(rdd.collect())

