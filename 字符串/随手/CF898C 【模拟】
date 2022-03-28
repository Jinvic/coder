//https://codeforces.com/contest/898/problem/C
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef pair<string, string> pii;
#define mp make_pair
const int N = 200005;
struct info
{
    string name;
    vector<string> num;
};
vector<info> fri;
bool judge(string s, string p)//后缀判断
{
    if (s.size() < p.size())
        swap(s, p);
    int l = min(s.size(), p.size());
    for (int i = 1; i <= l; i++)
        if (s[s.size() - i] != p[p.size() - i])
            return false;
    return true;
}
int main(void)
{
    ios::sync_with_stdio(false);
    int n, m;
    string name, s;
    vector<string> num;
    bool flag, flag2;
    cin >> n;
    while (n--)
    {
        cin >> name >> m;
        num.clear();
        while (m--)
        {
            cin >> s;
            flag = false;
            for (auto &p : num)//去重
                if (judge(s, p))
                {
                    flag = true;
                    if (s.size() > p.size())//更新更长串
                        p = s;
                    break;
                }
            if (!flag)
                num.push_back(s);
        }
        flag2 = false;
        for (auto &v : fri)
            if (name == v.name)
            {
                flag2 = true;
                for (auto &s : num)
                {
                    flag = false;
                    for (auto &p : v.num)//同上
                        if (judge(s, p))
                        {
                            flag = true;
                            if (s.size() > p.size())
                                p = s;
                            break;
                        }
                    if (!flag)
                        v.num.push_back(s);
                }
                break;
            }
        if (!flag2)
            fri.push_back({name, num});
    }
    cout << fri.size() << endl;
    for (auto &v : fri)
    {
        cout << v.name << " " << v.num.size();
        for (auto &p : v.num)
            cout << " " << p;
        cout << endl;
    }
    return 0;
}
