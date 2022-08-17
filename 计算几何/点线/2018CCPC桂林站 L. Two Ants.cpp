//2018CCPC桂林站 L. Two Ants
//https://codeforces.com/gym/102823/problem/L
#include <bits/stdc++.h>
using namespace std;
typedef long long int LL;
const int N = 100000 + 7;

#define abs(x) (((x) > 0) ? (x) : -(x))

/***************************************/

const double PI = acos(-1.0);
const double eps = 1e-8;
const double INF = 1e18;
#define pb push_back
#define mp make_pair
#define fi first
#define se second

///*************基础***********/
double torad(double deg) { return deg / 180 * PI; } //角度转换成弧度
inline int dcmp(double x)                           //和0比较 > = < 1 0 -1
{
    if (fabs(x) < eps)
        return 0;
    else
        return x < 0 ? -1 : 1;
}
struct Point
{
    double x, y;
    Point(double x = 0, double y = 0) : x(x), y(y) {}
    inline void read(void) { scanf("%lf%lf", &x, &y); };
    inline void set(double a, double b) { x = a, y = b; };
};
typedef vector<Point> Polygon;
typedef Point Vector;
//四则运算
inline Vector operator+(Vector A, Vector B) { return Vector(A.x + B.x, A.y + B.y); }
inline Vector operator-(Point A, Point B) { return Vector(A.x - B.x, A.y - B.y); }
inline Vector operator*(Vector A, double p) { return Vector(A.x * p, A.y * p); }
inline Vector operator/(Vector A, double p) { return Vector(A.x / p, A.y / p); }
//大小比较 先优先x再优先y
inline bool operator<(Vector A, Vector B) { return (A.x == B.x) ? A.y < B.y : A.x < B.x; }
inline bool operator==(Vector A, Vector B) { return dcmp(A.x - B.x) == 0 && dcmp(A.y - B.y) == 0; }
//数量积
inline double Dot(Vector A, Vector B) { return A.x * B.x + A.y * B.y; }
double msqrt(double x) //开方
{
    if (x < 0)
        return 0;
    else
        return sqrt(x);
}
double Length(Vector A) { return msqrt(Dot(A, A)); }                                        //向量长
double Length2(Vector A) { return Dot(A, A); }                                              //向量长平方
inline double Angle(Vector A, Vector B) { return acos(Dot(A, B) / Length(A) / Length(B)); } //向量夹角 a*b/(|a|*|b|)
inline double angle(Vector v) { return atan2(v.y, v.x); }                                   //向量与x轴夹角
inline double Cross(Vector A, Vector B) { return A.x * B.y - A.y * B.x; }                   // A x B > 0表示A在B的顺时针方向上
inline Vector Unit(Vector x) { return x / Length(x); }                                      //单位向量
inline Vector Normal(Vector x) { return Point(-x.y, x.x) / Length(x); }                     //垂直法向量,y反转
inline double Area2(Point A, Point B, Point C) { return Cross(B - A, C - A); }
double Cross3(Point p0, Point p1, Point p2) { return (p1.x - p0.x) * (p2.y - p0.y) - (p2.x - p0.x) * (p1.y - p0.y); } ///返回结果为正说明p0p1在p0p2的顺时针方向，正说明p2在向量p0p1的左边（三点构成逆时针方向）
                                                                                                                      ///也可以说明p0在向量p1p2的左边（三点构成逆时针方向）,为负则相反,为0则三点共线(叉积的性质很重要)
inline Vector Rotate(Vector A, double rad)
{
    return Vector(A.x * cos(rad) - A.y * sin(rad), A.x * sin(rad) + A.y * cos(rad));
}

//线段相交判定
inline int SegmentProperIntersection(Point a, Point b, Point c, Point d)
{ //快速排斥实验
    if (max(a.x, b.x) < min(c.x, d.x))
        return false;
    if (max(a.y, b.y) < min(c.y, d.y))
        return false;
    if (max(c.x, d.x) < min(a.x, b.x))
        return false;
    if (max(c.y, d.y) < min(a.y, b.y))
        return false;
    // c,d在ab两端，且a,b在cd两端
    double judge1 = dcmp(Cross(b - a, c - a)) * dcmp(Cross(b - a, d - a));
    double judge2 = dcmp(Cross(d - c, a - c)) * dcmp(Cross(d - c, b - c));
    /* return dcmp(Cross(b - a, c - a)) * dcmp(Cross(b - a, d - a)) <= 0 &&
           dcmp(Cross(d - c, a - c)) * dcmp(Cross(d - c, b - c)) <= 0; */
    if (judge1 < 0 && judge2 < 0) //相交
        return 2;
    else if (judge1 == 0 || judge2 == 0) //在端点上
        return 1;
    else //不相交
        return 0;
}

//直线ab与直线cd的交点
Point LineToLine(Point a, Point b, Point c, Point d)
{ ///已验证
    double x = Cross3(a, c, d), y = Cross3(b, c, d);
    int xx = dcmp(x), yy = dcmp(y);
    if (xx == yy && xx == 0)
        return {-1e10, 1e10}; //重合
    else if (dcmp(x - y) == 0)
        return {-1e10, -1e10}; //平行
    else
    {
        Vector u = a - c;
        Vector v = (b - a) / Length((b - a));
        Vector w = (d - c) / Length((d - c));
        return a + (Point)(v * (Cross(w, u) / Cross(v, w)));
    }
}

//点到直线距离
inline double DistanceToLine(Point P, Point A, Point B) ///已验证
{
    Vector v1 = B - A, v2 = P - A;
    return fabs(Cross(v1, v2)) / Length(v1); // 如果不取绝对值，得到的是有向距离
}

//点在p线段上(包括端点)
inline bool OnSegment(Point p, Point a1, Point a2)
{ // pa1和pa2共线，且a1a2在p两端
    return dcmp(Cross(a1 - p, a2 - p)) == 0 && dcmp(Dot(a1 - p, a2 - p)) <= 0;
}

//线段到直线的方位
int SegToLine(Point a, Point b, Point c, Point d)
{ /// Seg:ab  Line:cd  已验证
    if (c == d)
        return 3; // cd为同一点
    int x = dcmp(Cross3(a, c, d));
    int y = dcmp(Cross3(b, c, d));
    if (x == y && y == 0)
        return 0; //线上
    if (x + y == 2)
        return 1; //同侧 左边
    if (x + y == -2)
        return 2; //同侧 右边
    if (x + y == 0)
        return -1; //异侧
    if (x + y == 1)
        return -2; //一点在线上，另一点左边
    return -3;     //一点在线上，另一点右边
}

int main(void)
{
    Point w1, w2, b1, b2;
    int T;
    scanf("%d", &T);
    for (int t = 1; t <= T; t++)
    {
        w1.read();
        w2.read();
        b1.read();
        b2.read();
        printf("Case %d: ", t);
        Point wb1 = LineToLine(w1, w2, b1, b2); //两直线交点
        if (wb1.x == -1e10 && wb1.y == 1e10)    //重合
        {
            printf("0\n");
        }
        else if (wb1.x == -1e10 && wb1.y == -1e10) //平行
        {
            goto LABEL;
        }
        else //相交
        {
            if (OnSegment(wb1, w1, w2)) //交点在线段w上
            {
                if (SegToLine(b1, b2, w1, w2) == -1) // b线段端点在w直线异侧
                    printf("0\n");
                else //同侧
                    printf("inf\n");
            }
            else //交点不在线段w上
            {
                if (SegToLine(b1, b2, w1, w2) <= 0) //异测或一点在直线上
                    printf("0\n");
                else //线段b两端点在直线w同侧
                {
                LABEL:
                    if (SegmentProperIntersection(w1, b1, w2, b2)) //线段相交判定
                        swap(w1, w2);
                    Point wb2 = LineToLine(w1, b1, w2, b2);
                    if (wb2.x == -1e10 && wb2.y == -1e10) //平行
                        printf("inf\n");
                    else
                    {
                        if (SegToLine(b1, wb2, w1, w2) == -1) //线段b和交点wb2在直线w异侧
                        {
                            double d = DistanceToLine(wb2, w1, w2);
                            double l = Length(w2 - w1);
                            double s = l * d / 2.0;
                            printf("%.10f\n", s);
                        }
                        else //同侧
                        {
                            printf("inf\n");
                        }
                    }
                }
            }
        }
    }
}

/*
16
0 0 0 1 0 0 1 0
Case 1: inf
0 0 1 0 0 1 0 -1
Case 2: 0
1 1 1 2 0 0 0 3
Case 3: 0.2500000000
1 -2  0 -1  0 0 0 10
Case 4: inf
3 2 1 5 6 2 4 78
Case 5: 3.1849315068
9 1 9 6 5 6 9 4
Case 6: inf
1 4 0 5 5 0 6 0
Case 7: 0
3 2 4 2 5 3 5 8
Case 8: 0.7500000000
4 8 6 9 4 1 3 6
Case 9: 5.2500000000
4 4 8 4 3 3 7 4
Case 10: inf
2 4 3 4 8 5 6 4
Case 11: 0
1 9 9 1 5 5 1 4
Case 12: inf
4 3 4 4 4 6 2 9
Case 13: 0
3 3 2 8 2 9 1 2
Case 14: 0
3 3 2 8 2 9 2 1
Case 15: 0
1 4 4 3 4 4 8 3
Case 16: inf 
 */
