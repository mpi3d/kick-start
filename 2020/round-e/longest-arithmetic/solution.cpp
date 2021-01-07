#include <iostream>

using namespace std;

int solve() {
    int r, n, a, p, d, c;
    r = 0;
    c = 2;
    d = 1000000001;
    cin >> n;
    cin >> p;
    for (int i = 1; i < n; ++i) {
        cin >> a;
        if (d == p - a) ++c;
        else {
            c = 2;
            d = p - a;
        }
        if (r < c) r = c;
        p = a;
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
