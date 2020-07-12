#include <iostream>

using namespace std;

int solve() {
    int n, m, r;
    cin >> n;
    r = 0;
    m = -1;
    int v[n];
    for (int i = 0; i < n; ++i) cin >> v[i];
    for (int i = 0; i < n; ++i) {
        if (v[i] > m) {
            m = v[i];
            if (i == n - 1) ++r;
            else if (v[i] > v[i + 1]) ++r;
        }
    }
    return r;
}

int main() {
    int t;
    cin >> t;
    for (int i = 1; i <= t; ++i) {
        cout << "Case #" << i << ": " << solve() << endl;
    }
    return 0;
}
