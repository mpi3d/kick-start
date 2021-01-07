#include <iostream>

using namespace std;

int solve() {
    int k, p, a, r;
    cin >> k;
    int m[k];
    r = 0;
    cin >> p;
    m[0] = p;
    for (int i = 1; i < k; ++i) {
        cin >> a;
        if (a > p) m[i] = m[i - 1] + 1;
        else if (a < p) m[i] = m[i - 1] - 1;
        else m[i] = m[i - 1];
        p = a;
    }
    p = m[0];
    for (int i = 1; i < k; ++i) {
        a = m[i] - p;
        if (a > 3 || -3 > a) {
            p = m[i];
            ++r;
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
