//随机选取链表节点 O(1)空间复杂度
//https://leetcode-cn.com/problems/linked-list-random-node/solution/lian-biao-sui-ji-jie-dian-by-leetcode-so-x6it/
#include <bits/stdc++.h>
using namespace std;

//   Definition for singly-linked list.
struct ListNode
{
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

class Solution
{
    ListNode *head;

public:
    Solution(ListNode *head)
    {
        this->head = head;
    }
    int getRandom()
    {
        int i = 0, res = 0;
        ListNode *p = head;
        while (p != NULL)
        {
            if (rand() % (i + 1) == 0)
                res = p->val;
            p = p->next;
            i++;
        }
        return res;
    }
};
int main(void)
{
    ios::sync_with_stdio(false);
    cin.tie(0), cout.tie(0);
    int n, m, x;
    // cin >> n >> m;
    cin >> n;
    ListNode *head, *p, *pre;
    for (int i = 0; i < n; i++)
    {
        cin >> x;
        if (i == 0)
        {
            head = (ListNode *)malloc(sizeof(ListNode));
            head->val = x;
            head->next = NULL;
            pre = head;
        }
        else
        {
            p = (ListNode *)malloc(sizeof(ListNode));
            p->val = x;
            p->next = NULL;
            pre->next = p;
        }
    }
    Solution solu(head);
    cout << solu.getRandom();
    return 0;
}
