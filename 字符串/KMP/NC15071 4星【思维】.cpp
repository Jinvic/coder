// https://ac.nowcoder.com/acm/problem/15071
// https://blog.nowcoder.net/n/08e6a06008f54a788eabd7c68c492741
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef unsigned long long ull;
typedef pair<int, int> pii;
typedef pair<ll, int> pli;
typedef pair<ll, ll> pll;
const int N = 1000006;
const int mod = 998244353;

int nex[N << 1];
vector<char> vs[N << 1];
int len[N << 1];
void getnext(int idx)
{
#define p vs[idx]
	int m = len[idx];
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

ll kmp(int idxs, int idxp)
{
	if (len[idxs] < len[idxp])
		return 0;
#define s vs[idxs]
#define p vs[idxp]
	int n = len[idxs], m = len[idxp];
	int i = 0, j = 0;
	ll cnt = 0;
	while (i < n)
	{
		if (j == -1 || s[i] == p[j])
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
char ss[N << 1];
int minl = N << 1;
vector<int> mins;
int cmp(int i, int j)
{
	for (int k = 0; k < minl; k++)
	{
		if (vs[i][k] != vs[j][k])
			return false;
	}
	return true;
}
int main(void)
{
	ios::sync_with_stdio(false);

	int n;
	scanf("%d", &n);
	for (int i = 0; i < n; i++)
	{
		scanf("%s", ss);
		len[i] = strlen(ss);
		if (len[i] < minl) //保存最短串
		{
			minl = len[i];
			mins.clear();
			mins.push_back(i);
		}
		// printf("%d\n",len[i]);
		vs[i].resize(len[i] + 5);
		for (int j = 0; j < len[i]; j++) //从char存到vector
			vs[i][j] = ss[j];
		// vs[i].push_back(ss[j]);
	}
	for (int i = 1; i < mins.size(); i++)
	{
		if (!cmp(0, i)) //如果有不一样的最小串，答案全为0
		{
			for (int i = 0; i < n; i++)
				printf("0/n");
			return 0;
		}
	}
	ll ans = 1;
	getnext(mins[0]);
	for (int j = 0; j < n; j++) //求最小串的匹配结果
	{
		ans *= kmp(j, mins[0]);
		ans %= mod;
		if (ans == 0)
			break;
	}
	for (int i = 0; i < n; i++)
	{
		if (len[i] != minl) //非最小串和最小串匹配结果必为0
		{
			printf("0\n");
		}
		else
			printf("%lld\n", ans);
	}
	return 0;
}
