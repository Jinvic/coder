//https://codeforces.com/problemset/problem/797/B
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef pair<int, int> pii;
const int N = 100005;
const int inf = 0x7f7f7f7f;
const int mod = 1000000007;
const double eps = 1e-8;

signed main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(0), cout.tie(0);
	int n, x;
	cin >> n;
	int sum = 0;
	int odd_neg_max = 0, odd_pos_min = 0;
	for (int i = 0; i < n; i++)
	{
		cin >> x;
		if (x > 0)
			sum += x;
		if (x & 1)
		{
			if (x > 0)
				odd_pos_min = (odd_pos_min) ? min(odd_pos_min, x) : x;
			else
				odd_neg_max = (odd_neg_max) ? max(odd_neg_max, x) : x;
		}
	}
	if (!(sum & 1))
	{
		bool f1 = false, f2 = false;
		int sum1, sum2;
		if (odd_pos_min)
		{
			sum1 = sum - odd_pos_min;
			f1 = true;
		}
		if (odd_neg_max)
		{
			sum2 = sum + odd_neg_max;
			f2 = true;
		}
		if (f1 && f2)
			sum = max(sum1, sum2);
		else if (f1)
			sum = sum1;
		else
			sum = sum2;
	}
	cout << sum << endl;
	return 0;
}
