//https://ac.nowcoder.com/acm/problem/232808
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef unsigned long long ull;
typedef pair<int, int> pii;
typedef pair<ll, int> pli;
typedef pair<ll, ll> pll;
const int N = 2000006;
const int mod = 998244353;

struct Point
{
	ll x, y, k;
	Point() { x = 0, y = 0; }
	inline void read(void)
	{
		scanf("%lld%lld", &x, &y);
		if (y < 0||(y==0&&x<0))
			k = 0;
		else if ((y == 0 && x == 0))//注意特判（0,0）
			k = 1;
		else
			k = 2;
	}
	inline void output(void) { printf("%lld %lld\n", x, y); }
};
inline ll Cross(Point a, Point b) { return a.x * b.y - a.y * b.x; }
bool cmp(Point a, Point b)
{
	if (a.k != b.k)
		return a.k < b.k;
	else
		return Cross(a, b) > 0;
}
Point p[200005];
int main(void)
{
	int n;
	scanf("%d", &n);
	for (int i = 0; i < n; i++)
		p[i].read();
	sort(p, p + n, cmp);
	for (int i = 0; i < n; i++)
		p[i].output();
}
