 #include <iostream>
 #include <stdio.h>
 #include <algorithm>
 #include <stdlib.h>
 #include <string>
 #include <string.h>
 #include <set>
 #include <queue>
 #include <math.h>
 #include <stdbool.h>
 
 #define ll long long
 #define inf 0x3f3f3f3f
 using namespace std;
 const int MAXN = 1005;
 const int N = 1005;
 char a[N],b[N];
 int dp[N][N];

 int main()
 {
     int lena,lenb,i,j;
     cin >> a >> b;
     memset(dp,0,sizeof(dp));
     lena = strlen(a);
     lenb = strlen(b);
     for(i=1;i<=lena;i++)
     {
         for(j=1;j<=lenb;j++)
         {
             if(a[i-1]==b[j-1])
             {
                 dp[i][j]=dp[i-1][j-1]+1;
             }
             else{
                 dp[i][j]=max(dp[i-1][j],dp[i][j-1]);
             }
         }
     }
     int k=0;
     char ans[1005];
     int p = lena-1;
     int q = lenb-1;     
     while (p>=0 && q>=0)
     {
         if (dp[p][q] + 1 == dp[p + 1][q + 1] && a[p] == b[q]) {
             ans[k++] = a[p];
             p--;
             q--;
         } else if (dp[p][q + 1] > dp[p + 1][q]) 
                //可以理解第二个字符串往后走一个比一个字符串往后走一个更大 ，
                //那么q+1位置是最大的，所以不能移动它要移动p
             p--;
         else // 同理
             q--;
     }
     for (i=k-1;i>=0;i--)
     {
        //倒序输出答案
         cout << ans[i];
     }
     cout << endl;

     return 0;
 }
 /*使用动态规划的思想，将问题分解为小的子问题。
假设两个字符串序列分别为：X{x0, x1, x2,......, xm}, Y{y0, y1, y2,......, yn}。从后往前比较字符。
如果xm == yn， 则这个字符就是子序列中的一个字符， LCS就是序列{x0, x1, x2,......, xm-1}和{y0, y1, y2,......, yn-1}的LCS加上xm（或yn）。
如果xm != yn, 则求序列{x0, x1, x2,......, xm-1}与{y0, y1, y2,......, yn}的LCS1，{x0, x1, x2,......xm}与{y0, y1, y2......, yn-1}的LCS2，所求的LCS = max{LCS1, LCS2}.
*/


 // https://blog.csdn.net/qq_45800517/article/details/109210477