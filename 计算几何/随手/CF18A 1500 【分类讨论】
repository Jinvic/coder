//https://codeforces.com/problemset/problem/18/A
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
inline int dcmp(double x)							//和0比较 > = < 1 0 -1
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
double Length(Vector A) { return msqrt(Dot(A, A)); }										//向量长
double Length2(Vector A) { return Dot(A, A); }												//向量长平方
inline double Angle(Vector A, Vector B) { return acos(Dot(A, B) / Length(A) / Length(B)); } //向量夹角 a*b/(|a|*|b|)
inline double angle(Vector v) { return atan2(v.y, v.x); }									//向量与x轴夹角
inline double Cross(Vector A, Vector B) { return A.x * B.y - A.y * B.x; }					// A x B > 0表示A在B的顺时针方向上
inline Vector Unit(Vector x) { return x / Length(x); }										//单位向量
inline Vector Normal(Vector x) { return Point(-x.y, x.x) / Length(x); }						//垂直法向量,y反转
inline double Area2(Point A, Point B, Point C) { return Cross(B - A, C - A); }
double Cross3(Point p0, Point p1, Point p2) { return (p1.x - p0.x) * (p2.y - p0.y) - (p2.x - p0.x) * (p1.y - p0.y); } ///返回结果为正说明p0p1在p0p2的顺时针方向，正说明p2在向量p0p1的左边（三点构成逆时针方向）

inline Vector Rotate(Vector A, double rad)
{
	return Vector(A.x * cos(rad) - A.y * sin(rad), A.x * sin(rad) + A.y * cos(rad));
}

bool judge(Vector a, Vector b)
{
	a = Unit(a);
	b = Normal(b);
	Vector c = {-b.x, -b.y};
	if (a == b || a == c)
		return true;
	else
		return false;
}
bool judge1(Point a, Point b, Point c)
{
	return judge(a - b, b - c) || judge(b - c, c - a) || judge(c - a, a - b);
}
Vector offset[4] = {
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0}};
bool judge2(Point a, Point b, Point c)
{
	for (int i = 0; i < 4; i++)
		if (judge1(a + offset[i], b, c))
			return true;
	for (int i = 0; i < 4; i++)
		if (judge1(b + offset[i], a, c))
			return true;
	for (int i = 0; i < 4; i++)
		if (judge1(c + offset[i], a, b))
			return true;
	return false;
}
int main(void)
{
	Point a, b, c;
	a.read();
	b.read();
	c.read();
	if (judge1(a, b, c))
		printf("RIGHT");
	else if (judge2(a, b, c))
		printf("ALMOST");
	else
		printf("NEITHER");
	return 0;
}
