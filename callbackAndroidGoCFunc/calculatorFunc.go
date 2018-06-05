package callbackAndroidGoCFunc

/*

#include <stdio.h>

// The gateway function
int callbackSum_cgo(int in)
{
	int callbackSum(int);
	return callbackSum(in);
}

int callbackCubic_cgo(int in)
{
	int callbackCubic(int);
	return callbackCubic(in);
}
*/
import "C"
