//https://codeforces.com/problemset/problem/628/B
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef pair<int, int> pii;
const int N = 300005;
const int inf = 0x7f7f7f7f;
const int mod = 32768;
const double eps = 1e-8;
char s[N];
// 0	4	8	12	16
// 20	24	28	32	36
// 40	44	48	52	56
// 60	64	68	72	76
// 80	84	88	92	96
signed main(void)
{
	// ios::sync_with_stdio(false);
	// cin.tie(0), cout.tie(0);
	scanf("%s", s);
	int n = strlen(s);
	ll r = 0;
	int cnt = 0; //两位数数量
	if (s[n - 1] == '0' ||
		s[n - 1] == '4' ||
		s[n - 1] == '8')
		r++; //一位数计数
	for (int i = n - 2; i >= 0; i--)
	{
		// r += cnt;
		if (s[i] == '0' ||
			s[i] == '4' ||
			s[i] == '8')
			r++; //一位数计数
		int a = (s[i] - '0'), b = (s[i + 1] - '0');
		a %= 4;
		switch (a) //两位数计数
		{
		case 0:
			if (b == 0 || b == 4 || b == 8)
				cnt++;
			break;
		case 1:
			if (b == 2 || b == 6)
				cnt++;
			break;
		case 2:
			if (b == 0 || b == 4 || b == 8)
				cnt++;
			break;
		case 3:
			if (b == 2 || b == 6)
				cnt++;
			break;
		default:
			break;
		}
		r += cnt; //对2+位的数，若若后两位为4倍数，则该数必为4倍数
	}
	printf("%lld\n", r);
	return 0;
}
