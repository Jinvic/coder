// https://ac.nowcoder.com/acm/problem/16638
// https://blog.nowcoder.net/n/63ee5deaaf0e41d9b357ec2a74de0d01
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef unsigned long long ull;
typedef pair<int, int> pii;
typedef pair<ll, int> pli;
typedef pair<ll, ll> pll;
const int N = 1000006;
const int mod = 998244353;

int nex[N];
int v[N], val[N];
int n, m;
string A[N];
map<int, int> mp;
deque<int> dq;
void getnext1(int idx)
{
	int j = 0, k = -1;
	nex[0] = -1;
	while (j < m)
	{
		if (k == -1 || A[idx][j] == A[idx][k])
		{
			j++, k++;
			nex[j] = k;
		}
		else
			k = nex[k];
	}
}
void getnext2(int idx)
{
	int j = 0, k = -1;
	nex[0] = -1;
	while (j < n)
	{
		if (k == -1 || A[j][idx] == A[k][idx])
		{
			j++, k++;
			nex[j] = k;
		}
		else
			k = nex[k];
	}
}

int main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(0), cout.tie(0);
	cin >> n >> m;
	for (int i = 0; i < n; i++)
		cin >> A[i];
	for (int i = 0; i < n; i++)
		for (int j = 0; j < m; j++)
			cin >> v[i * m + j];

	int ln = n, lm = m;
	mp.clear();
	for (int i = 0; i < n; i++)
	{ //求每行next数组
		getnext1(i);
		int l = nex[m];
		while (l != -1)
		{
			mp[m - l]++;
			if (mp[m - l] == n)		 //所有行共同的循环节
				lm = min(lm, m - l); //更新最小循环节
			l = nex[l];
		}
	}
	mp.clear();
	for (int i = 0; i < m; i++)
	{ //求每列next数组
		getnext2(i);
		int l = nex[n];
		while (l != -1)
		{
			mp[n - l]++;
			if (mp[n - l] == m)		 //所有列共同的循环节
				ln = min(ln, n - l); //更新最小循环节
			l = nex[l];
		}
	}
	// cout << lm << " " << ln << endl;
	int ans = 1e9;
	for (int i = 0; i < n; i++)
	{
		while (dq.size())
			dq.pop_back();
		for (int j = 0; j < lm - 1; j++)
		{
			while (dq.size() && v[i * m + dq.back()] <= v[i * m + j])
				dq.pop_back();
			dq.push_back(j);
		}
		//维护以j为结尾的循环串最大值
		for (int j = lm - 1; j < m; j++)
		{
			while (dq.size() && dq.front() + lm <= j) //删去lm以外的部分
				dq.pop_front();
			while (dq.size() && v[i * m + dq.back()] <= v[i * m + j])
				dq.pop_back();
			dq.push_back(j);
			val[i * m + j] = v[i * m + dq.front()];
		}
	}
	for (int i = lm - 1; i < m; i++)
	{
		while (dq.size())
			dq.pop_back();
		//维护以ij为右下角的循环子矩阵最大值
		for (int j = 0; j < ln - 1; j++)
		{
			while (dq.size() && val[dq.back() * m + i] <= val[j * m + i])
				dq.pop_back();
			dq.push_back(j);
		}
		for (int j = ln - 1; j < n; j++)
		{
			while (dq.size() && dq.front() + ln <= j) //删去ln以外的部分
				dq.pop_front();
			while (dq.size() && val[dq.back() * m + i] <= val[j * m + i])
				dq.pop_back();
			dq.push_back(j);
			ans = min(ans, val[dq.front() * m + i]);
		}
	}
	ll res = (ll)ans * (ll)(lm + 1) * (ll)(ln + 1);
	cout << res << endl;
	return 0;
}
