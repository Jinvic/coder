//https://www.luogu.com.cn/problem/P5788
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef pair<int, int> pii;
const int N = 3000006;
const int inf = 0x7f7f7f7f;
const int mod = 1000000007;
const double eps = 1e-8;
int a[N], f[N];
signed main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(0), cout.tie(0);
	int n;
	cin >> n;
	stack<int> s;
	for (int i = 1; i <= n; i++)
	{
		cin >> a[i];
		while (s.size() && a[i] > a[s.top()])
		{
			f[s.top()] = i;
			s.pop();
		}
		s.push(i);
	}
	for (int i = 1; i <= n; i++)
		cout << f[i] << " ";
	return 0;
}
