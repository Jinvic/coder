//https://codeforces.com/problemset/problem/522/A
//不用hash
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef pair<int, int> pii;
const int N = 200005;
const int inf = 0x7f7f7f7f;
const int mod = 1000000007;
const double eps = 1e-8;

unordered_map<string, int> mp;
vector<int> G[405];
void change(string &s)
{
	for (auto &c : s)
		c = tolower(c);
}
int dfs(int u, int deep)
{
	int md = deep;
	for (auto v : G[u])
		md = max(md, dfs(v, deep + 1));
	return md;
}
signed main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(0), cout.tie(0);
	string s1, t, s2;
	int n;
	cin >> n;
	int root;
	string rs = "polycarp";
	int cnt = 0;
	while (n--)
	{
		cin >> s1 >> t >> s2;
		change(s1);
		change(s2);
		if (!mp[s1])
			mp[s1] = ++cnt;
		if (!mp[s2])
			mp[s2] = ++cnt;
		if (s1 == rs)
			root = mp[s1];
		if (s2 == rs)
			root = mp[s2];
		G[mp[s2]].push_back(mp[s1]);
	}
	cout << dfs(root, 1) << endl;
	return 0;
}

//用hash
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef unsigned long long ull;
typedef pair<int, int> pii;
const int N = 200005;
const int inf = 0x7f7f7f7f;
const int mod = 1000000007;
const double eps = 1e-8;

unordered_map<ull, int> mp;
vector<int> G[405];
void change(string &s)
{
	for (auto &c : s)
		c = tolower(c);
}
ull Hash(string s)
{
	ull r = 0;
	for (auto &c : s)
		r = r * 31 + c;
	return r;
}
int dfs(int u, int deep)
{
	int md = deep;
	for (auto v : G[u])
		md = max(md, dfs(v, deep + 1));
	return md;
}
signed main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(0), cout.tie(0);
	string s1, t, s2;
	int n;
	cin >> n;
	int root;
	string rs = "polycarp";
	int cnt = 0;
	while (n--)
	{
		cin >> s1 >> t >> s2;
		change(s1);
		change(s2);
		int u = Hash(s2), v = Hash(s1);
		if (!mp[v])
			mp[v] = ++cnt;
		if (!mp[u])
			mp[u] = ++cnt;
		if (s1 == rs)
			root = mp[v];
		if (s2 == rs)
			root = mp[u];
		G[mp[u]].push_back(mp[v]);
	}
	cout << dfs(root, 1) << endl;
	return 0;
}
