//https://blog.csdn.net/chen_ze_hua/article/details/124380287
// #include <bits/stdc++.h>
#include <iostream>
#include <cstdio>
#include <cstdlib>
#include <algorithm>
using namespace std;
typedef long long ll;
const int N = 100005;
const int mod = 1000000007;
int h, w, a, b;
long long quickpow(long long a, long long b)
{
    if (b < 0)
        return 0;
    long long ret = 1;
    a %= mod;
    while (b)
    {
        if (b & 1)
            ret = (ret * a) % mod;
        b >>= 1;
        a = (a * a) % mod;
    }
    return ret;
}
long long inv(long long a)
{
    return quickpow(a, mod - 2);
}
ll fac[2 * N];
ll facinv[2 * N];
ll C(int n, int m)
{
    return fac[n] * facinv[n - m] % mod * facinv[m] % mod;
}
signed main(void)
{
    ios::sync_with_stdio(false);
    cin.tie(0), cout.tie(0);
    fac[0] = 1;
    for (ll i = 1; i < 2 * N; i++)
        fac[i] = fac[i - 1] * i % mod;
    facinv[2 * N - 1] = inv(fac[2 * N - 1]);
    for (ll i = 2 * N - 2; i >= 0; i--)
        facinv[i] = facinv[i + 1] * (i + 1) % mod;
    cin >> h >> w >> a >> b;
    ll r = 0;
    // for (ll i = b; i < w; i++)
    // {
    //     r += C(h - a - 1 + i, i) * C(w - i + a - 2, w - i - 1) % mod;
    //     r %= mod;
    // }
    for (ll i = a + 1; i <= h; i++)
    {
        r += C(i + w - b - 2, i - 1) * C(h - i + b - 1, h - i) % mod;
        r %= mod;
    }
    cout << (r % mod + mod) % mod;
    return 0;
}
