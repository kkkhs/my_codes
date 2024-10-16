#include<iostream>
using namespace std;
const int N = 200000 + 10;
double pi[N], li[N], dp[N] = {0};

int main(){
    int n, t, q;
    cin >> n >> t >> q;
    for(int i = 0; i < n; i ++) cin >> pi[i];
    for(int i = 0; i < n; i ++) cin >> li[i];
    while(q--){
        int tj, rj = 0;
        cin >> tj >> rj;
        for (int i = 0; i <= t; i++) dp[i] = 0;

        if(tj == 1){
            
            dp[rj] = (dp[rj] * li[rj] + pi[rj]) / (li[rj] + 1);
            li[rj] ++;
        }
        else{
            
            dp[rj] = (dp[rj] * li[rj] - pi[rj]) / (li[rj] - 1);
            li[rj] --;
        }
        cout << dp[n];
    }
}
