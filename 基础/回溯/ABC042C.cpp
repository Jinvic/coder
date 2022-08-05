// #include <bits/stdc++.h>
#include <iostream>
#include <cstdio>
#include <cstdlib>
#include <algorithm>
#include <vector>
#include <stack>
using namespace std;
// typedef long long ll;
const int N = 5005;
vector<int> a;
vector<int> b;
int c = 0;
bool findequal(int x)
{
    for (auto y : b)
        if (x == y)
            return true;
    return false;
}
bool findbigger(int x, int &r)
{
    for (auto y : b)
        if (x < y)
        {
            r = y;
            return true;
        }
    return false;
}
int n;
int lft = 10;
vector<int> r;
vector<int> nd;
bool solve(int i)
{
    if (i == nd.size())
        return true;
    if (solve(i + 1))
    {
        int rt;
        if (findequal(nd[i]))
        {
            r[i] = nd[i];
            return true;
        }
        else if (findbigger(nd[i], rt))
        {
            r[i] = rt;
            lft = min(lft, i);
            return true;
        }
        else
            return false;
    }
    else
    {
        int rt;
        if (findbigger(nd[i], rt))
        {
            r[i] = rt;
            lft = min(lft, i);
            return true;
        }
        else
            return false;
    }
}
signed main(void)
{
    ios::sync_with_stdio(false);
    cin.tie(0), cout.tie(0);
    int n, k, x;
    cin >> n >> k;
    stack<int> s;
    while (n)
    {
        s.push(n % 10);
        n /= 10;
    }
    while (s.size())
    {
        nd.push_back(s.top());
        s.pop();
    }
    r.resize(nd.size());
    for (int i = 0; i < k; i++)
    {
        cin >> x;
        a.push_back(x);
    }
    for (int i = 0; i < 10; i++)
    {
        bool f = false;
        for (auto y : a)
        {
            if (i == y)
            {
                f = true;
                break;
            }
        }
        if (!f)
        {
            if (i != 0 && c == 0)
                c = i;
            b.push_back(i);
        }
    }
    // for (auto y : b)
    //     cout << y << " ";
    // cout << endl;
    if (solve(0))
    {
        for (int i = lft + 1; i < r.size(); i++)
            r[i] = b[0];
    }
    else
    {
        cout << c;
        for (int i = 0; i < r.size(); i++)
            r[i] = b[0];
    }
    for (int i = 0; i < r.size(); i++)
        cout << r[i];
    return 0;
}
