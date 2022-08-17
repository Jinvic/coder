//https://codeforces.com/problemset/problem/1474/B
//https://www.cnblogs.com/chantmee/p/14304349.html
#include <bits/stdc++.h>
using namespace std;
const int N = 200005;
bool isprime[N];
int prime[N];
int pcnt = 0;
void getprime(int size)
{
    memset(isprime, true, sizeof(bool) * size);
    for (int i = 2; i < size; i++)
        if (isprime[i])
        {
            prime[pcnt++] = i;
            for (int j = i + i; j < size; j += i)
                isprime[j] = false;
        }
}
int main(void)
{
    getprime(200000);
    int t, d;
    cin >> t;
    while (t--)
    {
        cin >> d;
        int a = *(lower_bound(prime, prime + pcnt, d + 1));
        int b = *(lower_bound(prime, prime + pcnt, d + a));
        cout << a * b << endl;
    }
    return 0;
}
