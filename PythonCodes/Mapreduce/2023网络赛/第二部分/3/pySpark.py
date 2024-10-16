from pyspark import SparkConf, SparkContext
import os

# 为Spark添加python解释器
os.environ['PYSPARK_PYTHON'] = "D:/Python/Python310/python.exe"

conf = SparkConf().setMaster("local[*]").setAppName("test_spark")
sc = SparkContext(conf=conf)

# 读取输入数据
data = sc.textFile("pri_user_adv.log")  # 你的输入文件路径

# 解析数据并创建一个元组 (省份, 广告, 点击数量)
parsed_data = data.map(lambda line: line.split(" ")).map(lambda tokens: ((tokens[1], tokens[4]), 1))

# 计算每个广告的总点击数量
ad_clicks = parsed_data.reduceByKey(lambda a, b: a + b)

# 切分省份和广告，然后按点击数量排序，选择前三个广告
top3_ads_by_province = ad_clicks \
    .map(lambda x: (x[0][0], [(x[0][1], x[1])])) \
    .reduceByKey(lambda a, b: a + b) \
    .mapValues(lambda ads: sorted(ads, key=lambda x: x[1], reverse=True)[:3])

# 打印结果
for province, top_ads in top3_ads_by_province.collect():
    for ad in top_ads:
        print("%s %s %s" % (province, ad[0], ad[1]))

# 停止 SparkContext
sc.stop()
