from pyspark import SparkConf, SparkContext
import os
import json

# 为Spark添加python解释器
os.environ['PYSPARK_PYTHON'] = "D:/Python/Python310/python.exe"

conf = SparkConf().setMaster("local[*]").setAppName("test_spark")
sc = SparkContext(conf=conf)

# TODO 需求1：城市销售额排名

# 1、读取文件得到rdd
rdd = sc.textFile("orders.txt")

# 2、取出一个json字符串
json_rdd = rdd.flatMap(lambda x: x.split("|"))

# 3、将一个json字符串转换为字典
dict_rdd = json_rdd.map(lambda x: json.loads(x))

# 4、取出城市和销售额数据
tuple_rdd = dict_rdd.map(lambda x: (x['areaName'], int(x['money'])))

# 5、按城市分组按销售额聚合
city_money = tuple_rdd.reduceByKey(lambda a, b: a + b)

# 6、按销售额聚合结果排序
city_result = city_money.sortBy(lambda x: x[1], ascending=False, numPartitions=1)
print(city_result.collect())

# TODO 需求2：全部城市有哪些商品类型在售卖

# 取出商品并去重
city_category = dict_rdd.map(lambda x: x['category']).distinct()
print(city_category.collect())

# TODO 需求3：北京市有哪些商品在售卖

# 1、过滤北京市的数据
BJ = dict_rdd.filter(lambda x: x['areaName'] == '北京')

# 2、取出全部商品类别并去重
BJ_category = BJ.map(lambda x: x['category']).distinct()
print(BJ_category.collect())
