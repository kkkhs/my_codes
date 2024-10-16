#include <iostream>
#include <queue>
using namespace std;

using namespace std;

typedef pair<int, int> PII;

const int N = 1e2 + 7;
int g[N][N], d[N][N];
int n, m;


int bfs() {
    int dx[] = {-1, 0, 1, 0}, dy[] = {0, 1, 0, -1};
    queue <PII> q;

    for (auto &v : d) 
        for (auto &x : v) {
            x = - 1;
        }

    d[0][0] = 0;
    q.push({0, 0});

    while (!q.empty()) {
        auto t = q.front();
        q.pop();

        for (int i = 0; i < 4; i++) {
            int x = t.first + dx[i], y = t.second + dy[i];

            if (x >= 0 && x < n && y >= 0 && y < m && g[x][y] == 0 && d[x][y] == -1) {
                d[x][y] = d[t.first][t.second]  + 1;
                q.push({x, y});
            }
        }

    }


        return d[n - 1][m - 1];
}


int main() {
    cin >> n >> m;
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < m; j++) {
            cin >> g[i][j];
        }
    }

    cout << bfs() << endl;

    return 0;
}
/*
输入样例：
5 5
0 1 0 0 0
0 1 0 1 0
0 0 0 0 0
0 1 1 1 0
0 0 0 1 0
*/