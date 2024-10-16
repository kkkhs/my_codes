#include <iostream>
#include <vector>
#include <limits>

using namespace std;

vector<int> prim(const vector<vector<double>>& graph, int start_node) {
    int n = graph.size();
    vector<bool> visited(n, false);
    vector<double> dist(n, numeric_limits<double>::max());
    vector<int> parent(n, -1);
    dist[start_node] = 0;
    for (int i = 0; i < n; ++i) {
        double min_dist = numeric_limits<double>::max();
        int min_node = -1;
        for (int j = 0; j < n; ++j) {
            if (!visited[j] && dist[j] < min_dist) {
                min_dist = dist[j];
                min_node = j;
            }
        }
        if (min_node == -1)
            break;
        visited[min_node] = true;
        for (int j = 0; j < n; ++j) {
            if (graph[min_node][j] > 0 && !visited[j] && graph[min_node][j] < dist[j]) {
                dist[j] = graph[min_node][j];
                parent[j] = min_node;
            }
        }
    }
    return parent;
}

void print_mst(const vector<vector<double>>& graph, const vector<int>& parent) {
    int n = graph.size();
    vector<vector<double>> mst(n, vector<double>(n, 0));
    double total_dist = 0;
    for (int i = 1; i < n; ++i) {
        mst[i][parent[i]] = mst[parent[i]][i] = graph[i][parent[i]];
        total_dist += graph[i][parent[i]];
    }
    cout << "最小生成树邻接矩阵：" << endl;
    for (auto& row : mst) {
        for (auto& val : row)
            cout << val << " ";
        cout << endl;
    }
    cout << "油管最短距离：" << total_dist +5 << endl;
    
}


int main() {
    vector<vector<double>> graph = { {0, 1.3, 2.1, 0.9, 0.7, 1.8, 2.0, 1.8},
                                    {1.3, 0, 0.9, 1.8, 1.2, 2.8, 2.3, 1.1},
                                    {2.1, 0.9, 0, 2.6, 1.7, 2.5, 1.9, 1.0},
                                    {0.9, 1.8, 2.6, 0, 0.7, 1.6, 1.5, 0.9},
                                    {0.7, 1.2, 1.7, 0.7, 0, 0.9, 1.1, 0.8},
                                    {1.8, 2.8, 2.5, 1.6, 0.9, 0, 0.6, 1.0},
                                    {2.0, 2.3, 1.9, 1.5, 1.1, 0.6, 0, 0.5},
                                    {1.8, 1.1, 1.0, 0.9, 0.8, 1.0, 0.5, 0} };
    vector<int> parent = prim(graph, 0);
    print_mst(graph, parent);
    return 0;
}
