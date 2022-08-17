//https://ac.nowcoder.com/acm/problem/235685
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef pair<int, int> pii;
const int N = 200005;
const double eps = 1e-8;
int main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(0), cout.tie(0);
	int t;
	ll n, a, b;
	cin >> t;
	while (t--)
	{
		cin >> n >> a >> b;
		if (a > b)
			cout << "niuniu" << endl;
		else if (a < b)
			cout << "niumei" << endl;
		else
		{
			ll l = 1, r = 1e9, round = 0, mid;
			while (l <= r)
			{
				mid = (l + r) / 2;
				if ((mid + 1) * (mid) / 2 <= n)
				{
					round = mid;
					l = mid + 1;
				}
				else
					r = mid - 1;
			}
			if (round % 2)
				cout << "niumei" << endl;
			else
				cout << "niuniu" << endl;
		}
	}
	return 0;
}
