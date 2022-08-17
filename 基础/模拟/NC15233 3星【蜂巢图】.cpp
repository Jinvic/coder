#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
const int N = 1005;
int dp[2 * N];

signed main(void)
{
    ios::sync_with_stdio(false);
    cin.tie(0), cout.tie(0);
    int n, x;
    cin >> n;
    int h = 4 * n - 3;
    int d = n - 1;
    int l = 2 * n - 1;
    for (int i = 0; i < d; i++)
    {
        int k = n - i; // left node
        for (int j = 0; j <= i; j++, k += 2)
        {
            if (i == 0)
                cin >> dp[n];
            else
            {

                cin >> x;
                if (n - i == k) // left side
                    dp[k] = x + dp[k + 1];
                else if (n + i == k) // right side
                    dp[k] = x + dp[k - 1];
                else
                    dp[k] = x + max(dp[k], max(dp[k - 1], dp[k + 1]));
            }
        }
    }
    for (int i = d; i < h - d; i++)
    {
        int k;
        k = ((i - d) & 1) ? 2 : 1;
        while (k <= l)
        {

            cin >> x;
            if (k == 1) // lest side
            {
                if (i == d)
                    dp[k] = x + dp[k + 1];
                else
                    dp[k] = x + max(dp[k], dp[k + 1]);
            }
            else if (k == l) // right side
            {
                if (i == d)
                    dp[k] = x + dp[k - 1];
                else
                    dp[k] = x + max(dp[k], dp[k - 1]);
            }
            else
                dp[k] = x + max(dp[k], max(dp[k - 1], dp[k + 1]));
            k += 2;
        }
    }
    for (int i = d - 1; i >= 0; i--)
    {
        int k = n - i; // left node
        for (int j = i; j >= 0; j--, k += 2)
        {
            cin >> x;
            /* if (n - i == k) // left side
                dp[k] = x + max(dp[k], dp[k - 1]);
            else if (n + i == k) // right side
                dp[k] = x + max(dp[k], dp[k + 1]);
            else */
            dp[k] = x + max(dp[k], max(dp[k - 1], dp[k + 1]));
        }
    }
    cout << dp[n] << endl;
    return 0;
}
