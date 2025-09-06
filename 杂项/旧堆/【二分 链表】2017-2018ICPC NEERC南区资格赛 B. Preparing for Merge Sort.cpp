// https://codeforces.com/contest/847/problem/B
//二分+链表
#include <iostream>
#include <cstdio>
#include <algorithm>
#include <cstring>
#include <vector>
#include <set>
using namespace std;
int n;
struct node //定义链表节点
{
    int val;
    int nex;
} lnk[200005];
int a, tot;
vector<int> head, tail; //存储链表头尾部
set<int> valt;          //存储链表尾的值
int main(void)
{
    ios::sync_with_stdio(false);
    cin >> n >> a;
    lnk[++tot].val = a;
    head.push_back(tot);
    tail.push_back(tot);
    valt.insert(a);
    for (int i = 1; i < n; i++)
    {
        cin >> a;
        if (a < *valt.begin()) // a小于所有链表尾
        {
            lnk[++tot].val = a;  //建新节点
            head.push_back(tot); //作为新链表放到数组尾
            tail.push_back(tot);
            valt.insert(a); //插入a
        }
        else if (a > lnk[tail[0]].val) // a大于所有链表
        {
            valt.erase(lnk[tail[0]].val); //删除原链表尾的值
            valt.insert(a);               //插入a
            lnk[++tot].val = a;           //新建节点
            lnk[tail[0]].nex = tot;       //将新节点添加到链表尾
            tail[0] = tot;
        }
        else
        {
            int l = 0, r = tail.size() - 1;
            int ans = 0;
            while (l <= r) //因为链表尾必为递减序列，使用二分查找
            {
                int mid = (l + r) >> 1;
                if (lnk[tail[mid]].val < a)
                {
                    ans = mid;
                    r = r - 1;
                }
                else
                    l = l + 1;
            }
            valt.erase(lnk[tail[ans]].val); //删除原链表尾的值
            valt.insert(a);                 //插入a
            lnk[++tot].val = a;             //新建节点
            lnk[tail[ans]].nex = tot;       //将新节点添加到链表尾
            tail[ans] = tot;
        }
    }
    for (auto x : head)
    {
        int y = x;
        while (y != 0)
        {
            cout << lnk[y].val << " ";
            y = lnk[y].nex;
        }
        cout << endl;
    }
    return 0;
}
