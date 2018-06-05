#ifndef CLIBRARY_H
#define CLIBRARY_H
typedef int (*callbackSum_fcn)(int);
void sum_c_func(callbackSum_fcn, int value, int value2);
typedef int (*callbackCubic_fcn)(int);
void cubic_c_func(callbackCubic_fcn, int value);
#endif