//https://ac.nowcoder.com/acm/problem/213329
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef unsigned long long ull;
typedef pair<int, int> pii;
typedef pair<ll, int> pli;
typedef pair<ll, ll> pll;
const int N = 10000007;
const int mod = 998244353;

int nex[N];
char s[N], p[N];
int n, m;
void getnext()
{
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
}

ll kmp()
{
	int i = 0, j = 0;
	ll cnt = 0;
	int pre = -1;//上一个匹配串的起始位置
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
			//左右长度相乘
			//+1是取0个
			//-pre是去重
			cnt += (n - i + 1LL) * (i - m - pre);
			j = nex[j];
			pre = i - m;
		}
	}
	return cnt;
}

int main(void)
{
	ios::sync_with_stdio(false);
	scanf("%d%d", &n, &m);
	scanf("%s", s);
	scanf("%s", p);
	getnext();
	printf("%lld", kmp());

	return 0;
}
