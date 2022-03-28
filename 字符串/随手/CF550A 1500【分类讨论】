//https://codeforces.com/problemset/problem/550/A
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef pair<string, string> pii;
#define mp make_pair
const int N = 200005;

int main(void)
{
    ios::sync_with_stdio(false);
    string s;
    cin >> s;
    bool f1 = false, f2 = false;
    for (int i = 0; i < s.size() - 1; i++)
        if (!f1 && s[i] == 'A' && s[i + 1] == 'B') // ABABAB
        {
            i++;
            f1 = true;
        }
        else if (!f2 && s[i] == 'B' && s[i + 1] == 'A')
        {
            i++;
            f2 = true;
        }
    if (f1 && f2)
        cout << "YES";
    else
    { // ABAXXAB
        reverse(s.begin(), s.end());
        f1 = false, f2 = false;
        for (int i = 0; i < s.size() - 1; i++)
            if (!f1 && s[i] == 'A' && s[i + 1] == 'B')
            {
                i++;
                f1 = true;
            }
            else if (!f2 && s[i] == 'B' && s[i + 1] == 'A')
            {
                i++;
                f2 = true;
            }
        if (f1 && f2)
            cout << "YES";
        else
            cout << "NO";
    }
    return 0;
}
