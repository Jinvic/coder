//https://codeforces.com/problemset/problem/1276/A
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef pair<int, int> pii;
const int N = 100005;
const int inf = 0x7f7f7f7f;
const int mod = 1000000007;
const double eps = 1e-8;
// int dp[N];
signed main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(0), cout.tie(0);
	int t;
	cin >> t;
	const string s1 = "one", s2 = "two", s3 = "twone";
	string s;
	vector<int> v;
	while (t--)
	{
		v.clear();
		cin >> s;
		for (int i = 0; i < s.size(); i++)
			if (i + 2 < s.size() && s.substr(i, 3) == s1)
			{
				v.push_back(i + 2);
				i += 2;
			}
			else if (i + 2 < s.size() && s.substr(i, 3) == s2)
			{
				if (i + 4 < s.size() && s.substr(i, 5) == s3)
				{
					v.push_back(i + 3);
					i += 4;
				}
				else
				{
					v.push_back(i + 2);
					i += 2;
				}
			}
		cout << v.size() << endl;
		for (auto x : v)
			cout << x << " ";
		cout << endl;
	}
	return 0;
}
