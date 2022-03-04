// 2020ICPC上海站 M. Gitignore
// https://codeforces.com/gym/102900/problem/M
#include <iostream>
#include <cstdio>
#include <cstring>
#include <algorithm>
#include <vector>
using namespace std;
const int maxn = 1003;
struct node //树节点
{
    string val;
    vector<int> nex;
    int res;   //该目录下需要忽略的文件数，即n
    bool flag; //该目录下是否存在不应忽略的文件，即m
} tr[1003];
void tr_clear(int p) //清空节点值
{
    tr[p].val = "";
    tr[p].res = 0;
    tr[p].nex.clear();
    tr[p].flag = false;
}

int dfs(int p) //深度优先搜索
{
    if (tr[p].nex.size() == 0) //叶子节点返回res值
    {
        return tr[p].res;
    }
    int res = 0;
    for (auto x : tr[p].nex)
    {
        res += dfs(x);           //累加子节点的res值
        if (tr[x].flag == false) //若子节点目录下存在不应忽略的文件
            tr[p].flag = false;  //则父节点目录下也存在不应忽略的文件
        tr_clear(x);             //遍历完成后清空子节点
    }
    if (tr[p].flag) //若改节点目录下不存在不应忽略的文件，直接忽略该文件夹
        return 1;
    else
        return res;
}

int main(void)
{
    int t, n, m;
    string s, subs;
    cin >> t;
    while (t--)
    {
        tr_clear(0); //清空根节点
        cin >> n >> m;
        int tot = 0; //总节点数（不含根节点）
        for (int i = 0; i < n; i++)
        { //建树
            cin >> s;
            s += "/";  //方便读取最后文件名
            int p = 0; //当前所在节点，初始化为根结点
            for (int j = 0, k = 0; j < s.length(); j++)
            {
                if (s[j] == '/')
                {
                    subs = s.substr(k, j - k); //截取文件夹名|文件名
                    k = j + 1;                 //标记上一个'/'所在位置
                    bool flag = false;         //假设不在已有子节点中
                    for (auto x : tr[p].nex)   //遍历子节点
                        if (tr[x].val == subs)
                        {
                            flag = true; //在子节点中
                            p = x;       //移动到该节点
                            break;       //跳出循环
                        }
                    if (!flag) //不在子节点中，添加新节点
                    {
                        tr[++tot].val = subs;
                        tr[tot].flag = true; //对叶子节点，该目录下只有他自身且需忽略，对非叶子节点无影响
                        tr[tot].res = 1;     //对叶子节点，标记为1；对非叶子节点无影响
                        tr[p].nex.push_back(tot);
                        p = tot;
                    }
                }
            }
        }
        for (int i = 0; i < m; i++)
        {
            cin >> s;
            s += "/";
            int p = 0;
            for (int j = 0, k = 0; j < s.length(); j++)
            {
                if (s[j] == '/')
                {
                    subs = s.substr(k, j - k);
                    k = j + 1;
                    bool flag = false;
                    for (auto x : tr[p].nex)
                        if (tr[x].val == subs)
                        {
                            flag = true;
                            p = x;
                            break;
                        }
                    if (!flag)
                    {
                        tr[++tot].val = subs;
                        tr[tot].flag = false; //对叶子节点，该目录下只有他自身且不应忽略，对非叶子节点无影响
                        tr[tot].res = 0;      //对叶子节点，标记为0；对非叶子节点无影响
                        tr[p].nex.push_back(tot);
                        p = tot;
                    }
                }
            }
        }
        printf("%d\n", dfs(0));
    }
    return 0;
}
