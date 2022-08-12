//https://codeforces.com/problemset/problem/1673/B
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef pair<int, int> pii;
const int N = 200005;
const int inf = 0x7f7f7f7f;
const int mod = 1000000007;
const double eps = 1e-8;
// int dp[N];
int nex[N];
bool vis[26];
void getnext(string s)
{
	int j = 0, k = -1;
	nex[0] = -1;
	while (j < s.size())
	{
		if (k == -1 || s[j] == s[k])
		{
			j++, k++;
			nex[j] = k;
		}
		else
			k = nex[k];
	}
}
signed main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(0), cout.tie(0);
	int t;
	cin >> t;
	string s;
	while (t--)
	{
		cin >> s;
		getnext(s);
		memset(vis, false, sizeof(vis));
		int len = s.size() - nex[s.size()];
		bool flag = true;
		for (int i = 0; i < len; i++)
		{
			int c = s[i] - 'a';
			if (vis[c])
			{
				flag = false;
				break;
			}
			else
				vis[c] = true;
		}
		if (flag)
			cout << "YES" << endl;
		else
			cout << "NO" << endl;
	}
	return 0;
}
