#include <iostream>
#include <vector>
#include <iomanip>

using namespace std;

int main() {
    int n, t, q;
    cin >> n >> t >> q;

    vector<int> p(n), l(n);
    for (int i = 0; i < n; i++) {
        cin >> p[i];
    }
    for (int i = 0; i < n; i++) {
        cin >> l[i];
    }

    vector<double> dp(n, 0);
    dp[0] = t;

    for (int i = 0; i < q; i++) {
        int tj, rj;
        cin >> tj >> rj;
        rj--;

        if (tj == 1) {
            dp[rj] = (dp[rj] * l[rj] + 1.0 * p[rj]) / (l[rj] + 1);
            l[rj]++;
        } else if (tj == 2) {
            dp[rj] = (dp[rj] * l[rj] - 1.0 * p[rj]) / (l[rj] - 1);
            l[rj]--;
        }

        cout << fixed << setprecision(9) << dp[n - 1] << endl;
    }

    return 0;
}
