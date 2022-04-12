//kmp变形
//https://ac.nowcoder.com/acm/problem/14694
//https://blog.nowcoder.net/n/f0cd95b0d66446eb8bcab629d0097df0
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef unsigned long long ull;
typedef pair<int, int> pii;
typedef pair<ll, int> pli;
typedef pair<ll, ll> pll;
const int N = 200005;
const int mod = 998244353;

int nex[N];
int a[N], b[N];
int n, m, k;
void getnext()
{
#define p b
	int j = 0, k = -1;
	nex[0] = -1;
	while (j < m)
	{
		if (k == -1 || p[j] == p[k])
		{
			j++, k++;
			nex[j] = k;
		}
		else
			k = nex[k];
	}
#undef p
}

ll kmp()
{
#define s a
#define p b
	int i = 0, j = 0;
	ll cnt = 0;
	while (i < n)
	{
		if (j == -1 || (s[i] + p[j]) % k == 0)
		{
			i++, j++;
		}
		else
			j = nex[j];
		if (j == m)
		{
			cnt++;
			j = nex[j];
		}
	}
#undef s
#undef p
	return cnt;
}

int main(void)
{
	ios::sync_with_stdio(false);
	int t;
	scanf("%d", &t);
	while (t--)
	{
		scanf("%d%d%d", &n, &m, &k);
		for (int i = 0; i < n; i++)
		{
			scanf("%d", &a[i]);
			a[i] %= k;
			if (i)
				a[i - 1] = ((a[i - 1] - a[i]) % k + k) % k;
		}
		for (int i = 0; i < m; i++)
		{
			scanf("%d", &b[i]);
			b[i] %= k;
			if (i)
				b[i - 1] = ((b[i - 1] - b[i]) % k + k) % k;
		}
		n--, m--;
		getnext();
		printf("%lld\n", kmp());
	}
	return 0;
}
