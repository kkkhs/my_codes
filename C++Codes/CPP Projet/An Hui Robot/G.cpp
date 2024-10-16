#include <iostream>
#include <cmath>
#include<iomanip>
using namespace std;
const int PI = 3.14159;

int main() {
    double xa, ya, za, xb, yb,zb, xc, yc, zc, r;
    cin >> xa >> ya >> za >> xb >> yb >> zb >> xc >> yc >> zc >> r;

    double AB = sqrt(pow(xa - xb, 2.0) + pow(ya - yb, 2.0) + pow(za - zb, 2.0));
    double AC = sqrt(pow(xa - xc, 2.0) + pow(ya - yc, 2.0) + pow(za - zc, 2.0));
    double BC = sqrt(pow(xb - xc, 2.0) + pow(yb - yc, 2.0) + pow(zb - zc, 2.0));

    //利用余弦定理求出∠CAB
    double cosCAB = (pow(BC, 2.0) - pow(AB, 2.0) - pow(AC, 2.0))/(2.0 * AB * AC);
    double sinCAB = sqrt(1 - pow(cosCAB, 2.0));
    double d = AC * sinCAB;
    //cout << d << endl;
    
    if(d >= 10) cout << fixed <<setprecision(2)<< AB << endl;
    else{
        //利用余弦定理求出∠ACB
        double cosACB = -(pow(AB, 2.0) - pow(AC, 2.0) - pow(BC, 2.0))/(2.0 * AC * BC);
        //cout << cosACB <<endl;

        double ACB = acos(cosACB);
        //cout << ACB * (180.0 / PI) <<endl;
        double ACD = acos(r/AC);
        double BCE = acos(r/BC);
        double DCE = ACB - ACD - BCE;
        double D = DCE * (180.0 / PI);
        //cout << D << endl;

        double AD = r * tan(ACD);
        //cout << AD << endl;
        double BE = r * tan(BCE);
        //cout << BE << endl;
        double DE = 2.0 * PI * r * (D / 360.0);
        //cout << DE << endl;

        double min = AD + BE + DE;

        cout << fixed <<setprecision(2)<< min << endl;
    }
    

    return 0;
}
