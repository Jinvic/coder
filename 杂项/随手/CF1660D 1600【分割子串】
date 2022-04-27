//https://codeforces.com/contest/1660/problem/D
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef pair<int, int> pii;
const int N = 200005;
int pos0[N], num[N], num2[N]; // 0的位置 负数数量 2数量
int cnt0;                     // 0的个数
int a[N];
int main()
{
    ios::sync_with_stdio(false);
    int t;
    cin >> t;
    int n;
    while (t--)
    {
        cin >> n;
        cnt0 = 0;
        pos0[cnt0++] = 0; //最前放一个0
        num[0] = 0;
        num2[0] = 0;
        for (int i = 1; i <= n; i++)
        {
            cin >> a[i];
            if (a[i] < 0) //负数计数
                num[i] = num[i - 1] + 1;
            else
                num[i] = num[i - 1];
            if (abs(a[i]) == 2) // 2计数
                num2[i] = num2[i - 1] + 1;
            else
                num2[i] = num2[i - 1];
            if (a[i] == 0) //存0位置
                pos0[cnt0++] = i;
        }
        pos0[cnt0++] = n + 1; //最后放一个0
        int max2 = 0, maxl = 1, maxr = 0;
        for (int i = 1; i < cnt0; i++) //遍历0分割子序列
        {
            int l = pos0[i - 1] + 1, r = pos0[i] - 1;
            if (l > r)
                continue;
            int cnt = num[r] - num[l - 1]; //负数数量

            if (cnt % 2) //为奇，删一个负数
            {
                int first = 0, last = 0;
                for (int i = l; i <= r; i++) //找第一个负数
                    if (a[i] < 0)
                    {
                        first = i;
                        break;
                    }
                for (int i = r; i >= l; i--) //找最后一个负数
                    if (a[i] < 0)
                    {
                        last = i;
                        break;
                    }
                //数2
                int cnt21 = num2[r] - num2[first];
                int cnt22 = num2[last - 1] - num2[l - 1];
                if (cnt21 >= cnt22 && cnt21 > max2)
                {
                    max2 = cnt21;
                    maxl = first + 1, maxr = r;
                }
                else if (cnt22 > max2)
                {
                    max2 = cnt22;
                    maxl = l, maxr = last - 1;
                }
            }
            else //为偶，直接数2
            {
                int cnt2 = num2[r] - num2[l - 1];
                if (cnt2 > max2)
                {
                    max2 = cnt2;
                    maxl = l, maxr = r;
                }
            }
        }
        cout << maxl - 1 << " " << n - maxr << endl;
    }
    return 0;
}
