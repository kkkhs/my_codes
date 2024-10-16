#include <iostream>
using namespace std;
const int N = 20;//对角线元素 2n-1 取20防止越界 
int n;
char g[N][N]; //存储图 
bool row[N],col[N], dg[N], udg[N]; //udg 副对角线 /

void dfs(int x,int y,int s){    //xy为坐标 (x,y) s为 n皇后放置个数
    if(y == n){   // 从左至右从上到下搜索
        y = 0;
        x ++;
    }

    if(x == n){
        if(s == n){      //放满
            for(int i = 0; i < n;i ++){  //输出答案
                puts(g[i]);
            }
            puts("");
        }
        return;  //继续执行
    }

    //1、该位置不放：尝试下一个位置
    dfs(x, y + 1, s);

    //2、该位置放
    if(! row[x] && !col[y] && !dg[x + y] && !udg[x - y + n])
    {
        g[x][y] = 'Q';//放皇后
        row[x] = col[y] = dg[x + y] = udg[x - y + n] = true;

        dfs(x, y + 1, s + 1);  // 找下一个

        //回溯时恢复现场
        row[x] = col[y] = dg[x + y] = udg[x - y + n] = false;
        g[x][y] = '.';
    }
}


int main () {
    cout <<"请输入棋盘大小："<<endl;
    cin >> n;
    for (int i = 0; i < n;i ++) {
        for (int j = 0; j < n; j ++) {
            g[i][j] = '.'; //初始化全部空格子
        }
    }
    dfs(0,0,0); //从(0,0)开始找

    return 0;
}


