#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef unsigned long long ull;
typedef pair<string, string> pii;
#define mp make_pair
const int N = 1000006;

int nex[N];
void getnext(string &p)
{
    int m = p.length();
    int j = 0, k = -1;
    nex[0] = -1;
    while (j < m)
    {
        if (k == -1 || p[j] == p[k])
        {
            j++;
            k++;
            nex[j] = k;
        }
        else
            k = nex[k];
    }
}

int kmp(string s, string p)
{
    int n = s.size();
    int m = p.size();
    int i = 0, j = 0;
    while (i < n)
    {
        if (j == -1 || s[i] == p[j])
        {
            i++;
            j++;
        }
        else
            j = nex[j];
        if (j == m)
            return i;
    }
    return -1;
}
int main(void)
{
    ios::sync_with_stdio(false);
    string s, p;
    cin >> s;
    getnext(s);
    p = s.substr(0, s.length());                                           //最大前后缀
    while (p.length() != 0 && (kmp(s.substr(1, s.length() - 2), p) == -1)) //未在非前后缀的子串中出现过
    {
        p = s.substr(0, nex[p.length()]); //缩小前后缀
    }
    if (p.length() == 0)
        cout << "Just a legend";
    else
        cout << p;
    return 
