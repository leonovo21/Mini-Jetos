#include<stdio.h>
#include<math.h>
#include <strio.h>
void math(){
	int a = 20;  
	int b = 10;
	
	printf("Say the first number:\n ");
	
	scanf("%d", &a);
	
	printf("Now the second Number:\n ");
	
	scanf("%d", &b);

	printf("Your result is: %d\n", a * b);
}

void net(){
	int new = 20;
	int old = 20;
	int	now = 0;
	now = new * old;
	printf("%d\n",now);
}

int main(){
	math();
	net();
	return 0;
}


