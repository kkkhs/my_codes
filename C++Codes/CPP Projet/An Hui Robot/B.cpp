#include<iostream>
using namespace std;

int main() {
	int T;
	int x;
	cin >> T;
	while (T --) {
		cin >> x;
		int a = 7;
		if (x == 1) cout << 3 << endl;
		else if (x == 2) cout << 5 << endl;
		else {
			int j = (x - 2) / 3;	// j是第几组
			int k = (x - 2) % 3;	// k是第j组的第几个
			if (k != 0) j++;		//往上进位

			a += (j - 1) * 4;		//j组的首个数字

			if (k == 1) cout << a << endl;		//输出答案
			else if (k == 2) cout << a + 1<< endl;
			else cout << a + 2 << endl;
		}
	}
	return 0;
}