//https://leetcode-cn.com/problems/pacific-atlantic-water-flow/
#include <bits/stdc++.h>
using namespace std;
typedef pair<int, int> pii;
int n, m;
pii offset[4] = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};
// private:
pii pii_plus(pii x, pii y)
{
    return {x.first + y.first, x.second + y.second};
}

// private:
bool valiable(pii x)
{
    if (0 <= x.first && x.first < n && 0 <= x.second && x.second < m)
        return true;
    return false;
}

// public:
vector<vector<int>> pacificAtlantic(vector<vector<int>> &heights)
{
    n = heights.size(), m = heights[0].size();
    vector<vector<int>> vis(n, vector<int>(m, 0));
    vector<vector<int>> res;
    queue<pii> q;
    pii p, tp;
    // Pacific Ocean
    for (int i = 0; i < n; i++)
    {
        q.push({i, 0});
        vis[i][0] = 1;
    }
    for (int j = 1; j < m; j++)
    {
        q.push({0, j});
        vis[0][j] = 1;
    }
    while (q.size())
    {
        p = q.front();
        q.pop();
        for (int i = 0; i < 4; i++)
        {
            tp = pii_plus(p, offset[i]);
            if (valiable(tp) && !vis[tp.first][tp.second] && heights[tp.first][tp.second] >= heights[p.first][p.second])
            {
                q.push(tp);
                vis[tp.first][tp.second] = 1;
            }
        }
    }
    // Atlantic Ocean
    for (int i = 0; i < n; i++)
    {
        q.push({i, m - 1});
        vis[i][m - 1] += 2;
        if (vis[i][m - 1] == 3)
            res.push_back({i, m - 1});
    }
    for (int j = m - 2; j >= 0; j--)
    {
        q.push({n - 1, j});
        vis[n - 1][j] += 2;
        if (vis[n - 1][j] == 3)
            res.push_back({n - 1, j});
    }
    while (q.size())
    {
        p = q.front();
        q.pop();
        for (int i = 0; i < 4; i++)
        {
            tp = pii_plus(p, offset[i]);
            if (valiable(tp) && vis[tp.first][tp.second] < 2 && heights[tp.first][tp.second] >= heights[p.first][p.second])
            {
                q.push(tp);
                vis[tp.first][tp.second] += 2;
                if (vis[tp.first][tp.second] == 3)
                    res.push_back({tp.first, tp.second});
            }
        }
    }
    return res;
}

int main(void)
{
    ios::sync_with_stdio(false);
    cin.tie(0), cout.tie(0);
    int n, m, x;
    cin >> n >> m;
    vector<vector<int>> res, h(n, vector<int>(m));
    for (int i = 0; i < n; i++)
        for (int j = 0; j < m; j++)
        {
            cin >> h[i][j];
        }
    res = pacificAtlantic(h);
    cout << endl;
    for (auto x : res)
    {
        for (auto y : x)
            cout << y << " ";
        cout << endl;
    }

    return 0;
}
