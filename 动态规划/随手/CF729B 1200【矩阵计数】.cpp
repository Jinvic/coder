//https://codeforces.com/problemset/problem/729/B
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef pair<int, int> pii;
const int N = 1003;
const int inf = 0x7f7f7f7f;
const int mod = 32768;
const double eps = 1e-8;

int stage[N][N];
// bool vis[N][N][4];
signed main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(0), cout.tie(0);
	int n, m, x;
	cin >> n >> m;
	ll res = 0;
	for (int i = 1; i <= n; i++)
	{
		bool left = false; //左边是否有演员
		int cnt = 0;	   //假设右边有演员 空格计数
		for (int j = 1; j <= m; j++)
		{
			cin >> stage[i][j];
			if (!stage[i][j])
			{
				if (left) //左边有演员时，聚光灯可以朝左摆放
					res++;
				cnt++; //空格计数
			}
			else
			{
				if (!left)
					left = true;
				res += cnt; //该演员左边的空格 可以朝右摆放
				cnt = 0;
			}
		}
	}
	for (int j = 1; j <= m; j++) //同上
	{
		bool up = false;
		int cnt = 0;
		for (int i = 1; i <= n; i++)
		{
			// cin >> stage[i][j];
			if (!stage[i][j])
			{
				if (up)
					res++;
				cnt++;
			}
			else
			{
				if (!up)
					up = true;
				res += cnt;
				cnt = 0;
			}
		}
	}
	cout << res << endl;
	return 0;
}
