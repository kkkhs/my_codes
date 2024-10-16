#include<bits/stdc++.h>

using namespace std;
const int M = 1e5+10, N = 1e6+10;//M为模板串长度，N为模式串长度
int m, n;
char p[M], s[N];//p为模板串（短），s为模式串（长）
int ne[M]; //next[]数组，这样取名避免和头文件next冲突

void get_next()//求next[]数组
{
        for(int i=2, j=0; i<=m; ++i)
        {
                while(j && p[i]!=p[j+1]) j = ne[j];
                if(p[i]==p[j+1]) ++j;
                ne[i] = j;
        }
}

int main()
{
        cin>>m>>p+1>>n>>s+1;//下标从1开始
        get_next();//求next[]数组
        
        //匹配操作
        for(int i=1, j=0; i<=n; ++i)
        {
                while(j && s[i]!=p[j+1]) j = ne[j];
                if(s[i]==p[j+1]) ++j;
                if(j==m)//满足匹配条件
                {
                        //匹配完成后的具体操作
                        //如：输出以0开始的匹配子串的首字母下标
                        printf("%d ", i-m);//若从1开始，加1
                        j = ne[j];//继续匹配
                }
        }
        
        return 0;
}
/*
输入样例：
3
aba
5
ababa

核心思想：

在每次失配时，不是把p串往后移一位，而是把p串往后移动至下一次可以和前面部分匹配的位置，
这样就可以 跳过大多数的失配步骤。
每次p串移动的步数通过查找next[ ]数组确定的。


假设 模板串 p = “abcab”，
则对于next[1]：前缀集合为空，后缀集合为空，next[1] = 0;
next[2]：前缀集合 { “a” }，后缀集合 { “b” }，两集合中无匹配字符串，next[2] = 0;
next[3]：前缀集合 { “a”, “ab” }，后缀集合 { “c”, “bc” }，两集合中无匹配字符串，next[3] = 0;
next[4]：前缀集合 { “a”, “ab”, “abc” }，后缀集合 { “a”, “ca”, “bca” }，两集合中最长匹配字符串为 “a”，next[4] = 1;
next[5]：前缀集合 { “a”, “ab”, “abc”, “abca” }，后缀集合 { “b”, “ab”, “cab”, “bcab” }，两集合中最长匹配字符串为 “ab”，next[5] = 2;
（注意以上说的前后缀都指的是“非平凡”）
*/


//https://blog.csdn.net/Jacob0824/article/details/123416168
