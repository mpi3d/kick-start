#include <iostream>

using namespace std;

int solve() {
    int n, k, a, r, s;
    cin >> n >> k;
    r = 0;
    s = k;
    for (int i = 0; i < n; ++i) {
        cin >> a;
        if (a != s) s = k;
        if (a == s) {
            if (s == 1) {
                s = k;
                ++r;
            }
            else --s;
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
