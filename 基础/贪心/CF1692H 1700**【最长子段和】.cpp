//https://codeforces.com/contest/1692/problem/H
//没完全懂
#include <bits/stdc++.h>
using namespace std;
const int N = 200005;
const int inf = 1e10;

map<int, vector<int>> mp;//存不同数的下标
signed main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(0), cout.tie(0);
	int t, n, x;
	cin >> t;
	while (t--)
	{
		cin >> n;
		mp.clear();
		for (int i = 0; i < n; i++)
		{
			cin >> x;
			mp[x].push_back(i);
		}
		int a, l, r, ans = -inf;
		for (auto [k, b] : mp)
		{
			int sum = 0, p = 0, mn = inf;
			for (int i = 0; i < b.size(); i++)
			{
				if (i)
					sum -= b[i] - b[i - 1] - 1;//中间的数贡献都为-1
				if (sum < mn)//更小的需要舍去的前缀
				{
					mn = sum;
					p = b[i];//更新起点
				}
				sum++;//计算贡献
				if (sum - mn > ans)//舍去贡献小的前缀的子段和
				{
					a = k;
					l = p + 1;
					r = b[i] + 1;
					ans = sum - mn;
				}
			}
		}
		cout << a << " " << l << " " << r << endl;
	}
	return 0;
}
