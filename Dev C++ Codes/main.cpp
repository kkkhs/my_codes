#include <iostream>
#include <ctime>
#include <cstdlib>
#include <graphics.h>	//EGE graphic library
using namespace std;

int main() {

	const int WIDTH = 800, HEIGHT = 600;
	initgraph(WIDTH, HEIGHT);	//initialize graphic

	PIMAGE pimage = newimage();
	getimage(pimage, "qq.png");

	char temp[50];
	int x = 0, y = 0;
	setfillcolor(EGERGB(0x0, 0xFF, 0x0));
	while (1) {
		cleardevice(); //«Â∆¡

		bar(0, 0, 100, 200);
		setcolor(EGERGB(0xFF, 0x00, 0x0));

		setfont(24, 0, "ÀŒÃÂ");

		sprintf(temp, "x:%d , y:%d", x, y);
		outtextxy(600, 200, temp);
		x ++;
		y ++;
		putimage(x * 20, y * 20, 20, 20, pimage, 0, 0, getwidth(pimage), getheight(pimage));


		Sleep(200);
	}







	return 0;
}
