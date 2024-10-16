from pyspark.sql import SparkSession
from pyspark.ml.recommendation import ALS
from pyspark.ml.evaluation import RegressionEvaluator
from pyspark.sql import Row
from pyspark.sql.functions import col

# 创建SparkSession
spark = SparkSession.builder.appName("MovieRecommendation").getOrCreate()

# 读取数据文件
data = spark.read.csv("capter5_2ml.csv", header=True, inferSchema=True)

# 拆分数据为训练集和验证集（8:2比例）
train_data, test_data = data.randomSplit([0.8, 0.2])

# 创建ALS模型
als = ALS(maxIter=5, regParam=0.01, userCol="userId", itemCol="movieId", ratingCol="rating")

# 拟合模型
model = als.fit(train_data)

# 为一组用户生成十大电影推荐
users = [1, 2, 3]  # 这里是示例的用户ID
user_recommendations = model.recommendForUserSubset(spark.createDataFrame([(user,) for user in users], ["userId"]), 10)

# 生成前十名用户推荐的一组指定的电影
movies = [10, 20, 30]  # 这里是示例的电影ID
movie_recommendations = model.recommendForItemSubset(spark.createDataFrame([(movie,) for movie in movies], ["movieId"]), 10)

# 输出结果
print("Top 10 movie recommendations for specified users:")
user_recommendations.show(truncate=False)

print("Top 10 user recommendations for specified movies:")
movie_recommendations.show(truncate=False)

# 停止SparkSession
spark.stop()
