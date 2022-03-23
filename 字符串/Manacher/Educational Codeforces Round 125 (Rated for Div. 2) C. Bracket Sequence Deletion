//https://codeforces.com/contest/1657/problem/C
#include <bits/stdc++.h>
using namespace std;

const int N = 1000006;
int n, hw[N];
// char a[N], s[N << 1];
string a, s;
void manacher()
{
    int maxright = 0, mid;
    for (int i = 1; i < n; i++)
    {
        if (i < maxright)
            hw[i] = min(hw[(mid << 1) - i], hw[mid] + mid - i);
        else
            hw[i] = 1;
        while (s[i + hw[i]] == s[i - hw[i]])
            ++hw[i];
        if (hw[i] + i > maxright)
        {
            maxright = hw[i] + i;
            mid = i;
        }
    }
}
void change()
{
    s[0] = s[1] = '#';
    for (int i = 0; i < n; i++)
    {
        s[i * 2 + 2] = a[i];
        s[i * 2 + 3] = '#';
    }
    n = n * 2 + 2;
    s[n] = 0;
}

int main(void)
{
    ios::sync_with_stdio(false);
    int t;
    // scanf("%d", &t);
    cin >> t;
    while (t--)
    {
        // scanf("%d", &n);
        // scanf("%s", a);
        cin >> n;
        cin >> a;
        int cnt = 0;
        bool flag = true;
        int left = 0;
        s.resize(n << 1);
        change();
        manacher();
        /*  for (int i = 0; i < n; i++)
             printf("%d ", hw[i]); */
        while (flag)
        {
            if (left >= a.size())
                break;
            flag = false;
            int cnt1 = 0, cnt2 = 0;
            int maxl = 0;
            bool flag2 = true;
            for (int i = left; i < a.size(); i++)
            {
                if (a[i] == '(')
                    cnt1++;
                else
                    cnt2++;
                if (cnt2 > cnt1)
                    flag2 = false;
                int mid = (i + left) / 2;
                int plen;
                if ((i + left) % 2)
                    plen = hw[mid * 2 + 3];
                else
                    plen = hw[mid * 2 + 2];
                if ((flag2 && cnt2 == cnt1) ||
                    ((i != left) && plen >= i - left + 1))
                {
                    maxl = i;
                    flag = true;
                    cnt++;
                    break;
                }
            }
            if (flag)
            {
                left = maxl + 1;
            }
        }
        cout << cnt << " " << a.size() - left << endl;
    }
}
