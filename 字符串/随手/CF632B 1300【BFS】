//https://codeforces.com/problemset/problem/653/B
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef pair<string, string> pii;
#define mp make_pair
pii p[40];
int n, q;
set<string> st;
int bfs(void)
{
    queue<string> que;
    string s;
    que.push("a");
    while (!que.empty() && que.front().size() < n)
    {
        s = que.front();
        que.pop();
        for (int i = 0; i < q; i++)
            if (p[i].first[0] == s[0])
                que.push(p[i].second + s.substr(1, s.size() - 1));
    }
    st.clear();
    while (!que.empty())
    {
        st.insert(que.front());
        que.pop();
    }
    return st.size();
}
int main(void)
{
    ios::sync_with_stdio(false);
    cin >> n >> q;
    for (int i = 0; i < q; i++)
        cin >> p[i].second >> p[i].first;
    sort(p, p + q);
    cout << bfs();
    return 0;
}
