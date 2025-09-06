//https://codeforces.com/problemset/problem/489/C
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef pair<int, int> pii;
const int N = 100005;
const int inf = 0x7f7f7f7f;
const int mod = 32768;
const double eps = 1e-8;

int a[105];
signed main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(0), cout.tie(0);
	int m, s;
	cin >> m >> s;
	if (s > 9 * m || (m > 1 && s == 0))
	{
		cout << "-1 -1" << endl;
		return 0;
	}
	int ss = s;
	s -= 1;
	for (int i = 0; i < m; i++)
	{
		if (s >= 9)
		{
			a[i] = 9;
			s -= 9;
		}
		else if (s)
		{
			a[i] = s;
			s = 0;
		}
		else
			a[i] = 0;
	}
	a[m - 1] += 1;
	for (int i = m - 1; i >= 0; i--)
		cout << a[i];
	cout << " ";
	s = ss;
	for (int i = 0; i < m; i++)
	{
		if (s >= 9)
		{
			a[i] = 9;
			s -= 9;
		}
		else if (s)
		{
			a[i] = s;
			s = 0;
		}
		else
			a[i] = 0;
	}
	for (int i = 0; i < m; i++)
		cout << a[i];
	cout << endl;
	return 0;
}
