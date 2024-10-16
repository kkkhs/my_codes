#include <iostream>
#include <string>
using namespace std;


int minDistance(string A, string B)
{
    int aLength = A.length();
    int bLength = B.length();
    int D[bLength + 1][aLength + 1];
    for(int j = 0 ; j <= aLength; j++){
            //初始化第一行
            D[0][j] = j;
    }
    for(int i = 0; i <= bLength; i++){
        //初始化第一列
        D[i][0] = i;
    }

    for(int i = 1; i <= bLength; i++){
        for(int j = 1; j <= aLength; j++){
            //初始化表格
            //D[i][j]=min(min(D[i-1][j]+1,D[i][j-1]+1),(A[j-1]==B[i-1]?D[i-1][j-1]: D[i-1][j-1]+1));
            //D[i][j]是指图中数字区域的每个单元格的值。
            D[i][j] = min(min( D[i - 1][j] + 1, D[i][j - 1] + 1), (A[j - 1] == B[i - 1] ? D[i - 1][j - 1] : D[i - 1][j - 1] + 1));
        }
    }

    return D[bLength][aLength];
}

int main()
{
    string a, b;
    cin >> a;
    cin >> b;

    cout << minDistance(a, b);

    return 0;
}

/*
开一个二维数组d[i][j]来记录a0-ai与b0-bj之间的编辑距离，
要递推时，需要考虑对其中一个字符串的删除操作、插入操作和替换操作分别花费的开销，
从中找出一个最小的开销即为所求。
*/