#include <stdio.h>

#include "calculator.h"

void sum_c_func(callbackSum_fcn callback, int value, int value2)
{
	int response = callback(value + value2);
}

void cubic_c_func(callbackCubic_fcn callback, int value)
{
	int response = callback( value * value* value);
}