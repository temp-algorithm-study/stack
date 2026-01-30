
#include <iostream>
#include <string>
using namespace std;

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int N;
    cin >> N;

    int b[10000];
    int bcount = 0;

    string cmd;
    for (int i = 0; i < N; i++) {
        cin >> cmd;

        if (cmd == "push") {
            int x;
            cin >> x;
            b[bcount] = x;
            bcount++;
        }
        else if (cmd == "pop") {
            if (bcount == 0) {
                cout << -1 << '\n';
            } else {
                cout << b[bcount - 1] << '\n';
                bcount--;
            }
        }
        else if (cmd == "size") {
            cout << bcount << '\n';
        }
        else if (cmd == "empty") {
            cout << (bcount == 0 ? 1 : 0) << '\n';
        }
        else if (cmd == "top") {
            if (bcount == 0) cout << -1 << '\n';
            else cout << b[bcount - 1] << '\n';
        }
    }

    return 0;
}
