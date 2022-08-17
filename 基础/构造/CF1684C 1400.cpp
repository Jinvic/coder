//https://codeforces.com/contest/1684/problem/C
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef pair<int, int> pii;
const int N = 300005;
const int inf = 0x7f7f7f7f;
const int mod = 1000000007;
const double eps = 1e-8;

int n, m;
bool judge(vector<vector<int>> &a, int x, int y)
{
    for (int i = 0; i < n; i++)
    {
        swap(a[i][x], a[i][y]);
        for (int j = 1; j < m; j++)
            if (a[i][j] < a[i][j - 1])
                return false;
    }
    return true;
}
signed main(void)
{
    ios::sync_with_stdio(false);
    cin.tie(0), cout.tie(0);
    int t;

    cin >> t;
    while (t--)
    {
        cin >> n >> m;
        vector<vector<int>> a(n, vector<int>(m));
        for (int i = 0; i < n; i++)
            for (int j = 0; j < m; j++)
                cin >> a[i][j];
        vector<int> wa;
        for (int i = 0; i < n; i++)
        {
            vector<int> b = a[i];
            sort(b.begin(), b.end());
            for (int j = 0; j < m; j++)
                if (a[i][j] != b[j])
                    wa.push_back(j);
            if (wa.size())
                break;
        }
        if (!wa.size())
            cout << 1 << " " << 1 << endl;
        else if (wa.size() > 2)
            cout << -1 << endl;
        else if (judge(a, wa[0], wa[1]))
            cout << wa[0] + 1 << " " << wa[1] + 1 << endl;
        else
            cout << -1 << endl;
    }
    return 0;
}
