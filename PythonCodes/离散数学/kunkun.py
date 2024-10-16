'''
# 邻接矩阵，用0表示没有边，用正整数表示边的权值（通行费用）
0 0 1892 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 216 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
1892 216 0 676 1145 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 676 0 0 0 511 842 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 1145 0 0 668 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 668 0 695 0 0 0 0 0 0 0 0 0 0 137 0 0 0 0 0 0 0 0
0 0 0 511 0 695 0 0 0 0 0 0 0 0 0 0 534 0 349 0 0 0 0 0 0 0
0 0 0 842 0 0 0 0 1100 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 1100 0 639 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 639 0 902 607 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 902 0 672 0 675 0 0 528 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 607 672 0 255 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 675 0 0 0 675 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 675 0 140 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 140 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 534 0 0 0 528 0 0 0 0 0 0 0 0 375 0 0 0 0 622 0
0 0 0 0 0 137 0 0 0 0 0 0 0 0 0 0 0 0 674 0 704 0 0 0 0 0
0 0 0 0 0 0 349 0 0 0 0 0 0 0 0 0 0 674 0 651 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 375 0 651 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 704 0 0 0 305 0 397 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 305 0 242 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 242 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 397 0 0 0 0 550
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 622 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 550 0 0
'''
city_map = {chr(65 + i): i for i in range(26)}

def dijkstra_core(adj_matrix, source, destination):
    n = len(adj_matrix)
    S = set()
    U = set(range(n))
    dist = [float("inf")] * n
    prev = [-1] * n
    dist[source] = 0
    while U and destination not in S:
        u = min(U, key=lambda x: dist[x])
        S.add(u)
        U.remove(u)
        for v in U:
            if adj_matrix[u][v] > 0 and dist[v] > dist[u] + adj_matrix[u][v]:
                dist[v] = dist[u] + adj_matrix[u][v]
                prev[v] = u
    if dist[destination] == float("inf"):
        return None
    path = []
    cost = dist[destination]
    curr = destination
    while curr != source:
        path.append(curr)
        curr = prev[curr]
    path.append(source)
    path.reverse()
    return path, cost

def dijkstra(city_map, adj_matrix):
    source = input("请输入源点城市：")
    destination = input("请输入目的城市：")
    result = dijkstra_core(adj_matrix, city_map[source], city_map[destination])
    if result is None:
        print("从{}到{}没有可达的路径。".format(source, destination))
    else:
        path, cost = result
        print("从{}到{}的最短路径是：".format(source, destination))
        print(' -> '.join([chr(65 + node) for node in path]))
        print("从{}到{}的最少花费是：{}".format(source, destination, cost))

def floyd_warshall_core(adj_matrix):
    n = len(adj_matrix)
    dist = [[adj_matrix[i][j] if i != j else 0 for j in range(n)] for i in range(n)]
    for k in range(n):
        for i in range(n):
            for j in range(n):
                if i != j and dist[i][k] > 0 and dist[k][j] > 0:
                    if dist[i][j] == 0 or dist[i][j] > dist[i][k] + dist[k][j]:
                        dist[i][j] = dist[i][k] + dist[k][j]
    return dist

def floyd_warshall(city_map, adj_matrix):
    dist = floyd_warshall_core(adj_matrix)
    print("可达矩阵（包含每两个城市间最少花费）：")
    for i in range(len(adj_matrix)):
        print(' '.join(str(dist[i][j]) for j in range(len(adj_matrix))))

def main():
    mode = input("请选择模式（单源 1或多源 2）：")
    print("请输入邻接矩阵，用空格分隔每个元素，用换行分隔每一行：")
    adj_matrix = [[int(x) for x in input().split()] for _ in range(len(city_map))]

    if mode == "1":
        dijkstra(city_map, adj_matrix)
    elif mode == "2":
        floyd_warshall(city_map, adj_matrix)

main()