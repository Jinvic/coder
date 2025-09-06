#include <bits/stdc++.h>
using namespace std;
// typedef long long ll;
typedef pair<int, int> pii;
const int N=25;
int n;
int a[N], b[N];
int lson[N], rson[N];
int ll[N], rr[N];
int build(int la, int ra, int lb, int rb)
{
	if (la > ra || lb > rb)
		return -1;

	int fa = b[rb];
	int p;
	for (int i = 1; i <= n; i++)
		if (a[i] == fa)
		{
			p = i;
			break;
		}
	// cout << p << " " << la << " " << ra << " " << lb << " " << rb << endl;
	int llen = p - la, rlen = ra - p;
	lson[fa] = build(la, p - 1, lb, rb - rlen - 1);
	rson[fa] = build(p + 1, ra, rb - rlen, rb - 1);

	return fa;
}

int main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(0), cout.tie(0);
	cin >> n;
	for (int i = 1; i <= n; i++)
	{
		cin >> a[i];
	}
	for (int i = 1; i <= n; i++)
	{
		cin >> b[i];
	}
	build(1, n, 1, n);
	// for (int i = 1; i <= n; i++)
	// 	cout << lson[i] << " " << rson[i] << endl;
	int root = b[n];
	// ll[1] = rr[1] = root;
	int now = 0;
	queue<int> q[2];
	q[0].push(root);
	int deep = 0;
	while (1)
	{
		ll[++deep] = q[now].front();
		while (!q[now].empty())
		{
			int p = q[now].front();
			q[now].pop();
			if (q[now].empty())
				rr[deep] = p;
			if (lson[p] != -1)
				q[now ^ 1].push(lson[p]);
			if (rson[p] != -1)
				q[now ^ 1].push(rson[p]);
		}
		if (q[now ^ 1].empty())
			break;
		now = now ^ 1;
	}
	cout << "R:";
	for (int i = 1; i <= deep; i++)
		cout << " " << rr[i];
	cout << endl;
	cout << "L:";
	for (int i = 1; i <= deep; i++)
		cout << " " << ll[i];
	cout << endl;
	return 0;
}
