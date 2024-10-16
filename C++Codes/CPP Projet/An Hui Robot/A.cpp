#include<iostream>
using namespace std;

int main(){
    double N, M, K;
    int count = 0;
    cin >> N >> M >> K;
    double t = (100 - K) / 100; // t表示下一年是这一年的倍数

    while(1){
        if(N <= M){
            cout << count;
            break;
        }
        else{
            N *= t;
            count ++;
        }
    }
    return 0;
}