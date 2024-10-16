#include <iostream>
#include <cstdio>
#include <cstring>
#include <algorithm>
using namespace std;
const int N = 105;
double dp[2][N];
int main() {
    int c, n, m, i, j, cur;
    while(~scanf("%d", &c), c) {
        scanf("%d%d", &n, &m);

        if(m > c || m > n || (n-m)&1) {//考虑奇偶，观察可得
            printf("0.000\n");
            continue;
        }

        if(n > 500) n = 500 - (n&1);//精度要求不高，n较大无影响

        memset(dp, 0, sizeof(dp));
        dp[0][0] = dp[1][1] = 1.0;//边界

        cur = 1;//运算优化，虽然感觉这里效果不大。。
        int x = n-1;
        for(i = 2; i <= n; ++i) {
            dp[1-cur][0] = dp[cur][1] / c;
            dp[1-cur][c] = dp[cur][c-1] / c;

            for(j = 1; j <= i && j < c; ++j) {
                dp[1-cur][j] = (j+1.0)/c * dp[cur][j+1] + (c-j+1.0)/c * dp[cur][j-1];
            }
            cur = 1-cur;
        }

        printf("%.3f\n", dp[n&1][m]);
    }
    return 0;
}