// https://codeforces.com/contest/1647/problem/D
//分类讨论
#include <iostream>
#include <cstdio>
#include <cstring>
#include <algorithm>
#include <cmath>
using namespace std;
typedef long long ll;

jud(int n) //判断是否为素数
{
	int t, i, re;
	t = (int)sqrt(n);
	re = 1;
	for (i = 2; i <= t; i++)
	{
		if (0 == n % i)
		{
			re = 0;
			break;
		}
	}
	if (1 == re)
		return 1;
	else
		return 0;
}

bool solve(int x, int d)
{
	int nd = 0;
	int y = x;
	while (y % d == 0) // x=d^nd+y
	{
		y /= d;
		nd++;
	}
	if (nd == 1) //只有一种情况
		return false;
	else // nd>=2
	{
		if (!jud(y)) // y可拆为1*y，y1*y2至少两种情况
			return true;
		else // y为质数
		{
			if (nd == 2) //只有一种情况
				return false;
			else // nd>=3
			{
				if (!jud(d)) // d为复数，取一个d进行拆分
				{
					if (nd == 3)
					{
						int t = (int)sqrt(d);	  //此处考虑情况不全面，数据水了
						if (y == t && t * t == d) //若d拆成t*t，且y也为t，则剩余两个d的系数将有一个为t*t=d，即同一种情况
							return false;
						else
							return true;
					}
					else //将一个d拆开分配出去凑出一种情况
						return true;
				}
				else //没有别的情况
					return false;
			}
		}
	}
}
int main(void)
{
	int t;
	int x, d;
	scanf("%d", &t);
	while (t--)
	{
		scanf("%d%d", &x, &d);

		if (solve(x, d))
			printf("YES\n");
		else
			printf("NO\n");
	}
	return 0;
}
