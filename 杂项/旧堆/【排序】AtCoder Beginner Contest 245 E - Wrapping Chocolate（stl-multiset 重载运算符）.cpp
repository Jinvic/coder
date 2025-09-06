#include <iostream>
#include <cstdio>
#include <algorithm>
#include <string>
#include <cstring>
#include <cmath>
#include <cstdlib>
#include <vector>
#include <set>
using namespace std;
typedef long long ll;
const int N = 200005;
int n, m;
struct rect
{
    int width, length;
    int cato;
} r[N << 1];
bool operator<(rect a, rect b)
{
    return (a.length == b.length) ? a.width < b.width : a.length < b.length;
}
bool cmp(rect a, rect b)
{
    return (a.width == b.width) ? ((a.cato != b.cato)
                                       ? ((a.cato == 0) ? true : false)
                                       : a.length < b.length)
                                : a.width < b.width;
}
multiset<rect> s;

int main()
{
    ios::sync_with_stdio(false);
    cin >> n >> m;
    for (int i = 0; i < n; i++)
        cin >> r[i].width;
    for (int i = 0; i < n; i++)
    {
        cin >> r[i].length;
        r[i].cato = 0;
    }
    for (int i = 0; i < m; i++)
        cin >> r[n + i].width;
    for (int i = 0; i < m; i++)
    {
        cin >> r[n + i].length;
        r[n + i].cato = 1;
    }
    int nm = n + m;
    sort(r, r + nm, cmp);
    bool flag = true;
    for (int i = nm - 1; i >= 0; i--)
    {
        if (r[i].cato == 1)
            s.insert(r[i]);
        else
        {
            auto it = s.lower_bound(r[i]);
            // cout << it->width << " " << it->length << " " << it->cato;
            if (it == s.end())
            {
                flag = false;
                break;
            }
            else
            {
                s.erase(it);
            }
        }
    }
    if (flag)
        cout << "Yes";
    else
        cout << "No";
    return 0;
}
