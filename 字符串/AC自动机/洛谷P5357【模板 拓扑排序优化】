//https://www.luogu.com.cn/problem/P5357
//https://www.luogu.com.cn/blog/juruohyfhaha/solution-p5357
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef unsigned long long ull;
typedef pair<int, int> pii;
typedef pair<ll, int> pli;
typedef pair<ll, ll> pll;
const int N = 2000006;
const int mod = 998244353;

int tr[N][26], tot;
int vis[N]; //访问次数
int pos[N];
int in[N]; //入度
int num[N];
int fail[N];
void insert(char *s, int idx)
{
	int u = 0, len = strlen(s);
	for (int i = 0; i < len; i++)
	{
		if (!tr[u][s[i] - 'a'])
		{
			// memset(tr[tot], 0, sizeof(tr[tot]));
			tr[u][s[i] - 'a'] = ++tot; //如果没有则插入新节点
		}
		u = tr[u][s[i] - 'a']; //搜索下一个节点
	}
	pos[idx] = u; /* 标记每个模式串对应节点 */
}

void build()
{
	queue<int> q;
	for (int i = 0; i < 26; i++)
		if (tr[0][i])
			q.push(tr[0][i]);
	while (q.size())
	{
		int u = q.front();
		q.pop();
		// st.push(u);
		for (int i = 0; i < 26; i++)
		{
			if (tr[u][i])
			{
				fail[tr[u][i]] =
					tr[fail[u]][i];	  // fail数组：同一字符可以匹配的其他位置
				in[fail[tr[u][i]]]++; //统计入度
				q.push(tr[u][i]);
			}
			else
				tr[u][i] = tr[fail[u]][i];
		}
	}
}
void query(char *t)
{
	int u = 0, len = strlen(t);
	for (int i = 0; i < len; i++)
	{
		u = tr[u][t[i] - 'a']; // 转移
		vis[u]++;			   //标记节点访问次数
	}
}
void topu() //拓扑排序
{
	queue<int> q;
	for (int i = 1; i <= tot; i++)
		if (in[i] == 0)
			q.push(i);
	while (q.size())
	{
		int u = q.front();
		q.pop();
		num[u] = vis[u];
		int v = fail[u];
		vis[v] += vis[u];
		in[v]--;
		if (in[v] == 0)
			q.push(v);
	}
}
char s[N];
int main(void)
{
	int n;
	scanf("%d", &n);
	for (int i = 0; i < n; i++)
	{
		scanf("%s", s);
		insert(s, i);
	}
	build();
	scanf("%s", s);
	query(s);
	topu();
	for (int i = 0; i < n; i++)
		printf("%d\n", num[pos[i]]);
}
