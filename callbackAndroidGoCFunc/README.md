android要求 golang執行動作, golang或將動作轉交給 c處理. c處理好後會 callback golang, golang在 callback android.

android code
```java
package com.example.ubuntu.myapplication13;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import callbackAndroidGoCFunc.CalculatorInterface;
import callbackAndroidGoCFunc.CallbackAndroidGoCFunc;
import android.util.*;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        long result =  CallbackAndroidGoCFunc.multiply(3,5);
        Log.i("testCallback","multiply, result="+result);// multiply, result=15

        Log.i("testCallback","-----------------------------------------------------------");


        boolean registered = false;
        registered = CallbackAndroidGoCFunc.isRegistered();
        Log.i("testCallback","isRegistered, result="+registered); // isRegistered, result=false

        long successful;
        successful = CallbackAndroidGoCFunc.registerCalculatorInterface(new CalculatorInterface(){
            public void squareResult(long var1){
                Log.i("testCallback","squareResult, result="+var1);// squareResult, result=9
            }

            public void sumResult(long var1){
                Log.i("testCallback","sumResult, result="+var1); // sumResult, result=14
            }

            public void cubicResult(long var1){
                Log.i("testCallback","cubicResult, result="+var1); // cubicResult, result=1000
            }
        });
        Log.i("testCallback","exec registerCalculatorInterface, result="+successful); // exec registerCalculatorInterface, result=0

        registered = CallbackAndroidGoCFunc.isRegistered();
        Log.i("testCallback","isRegistered, result="+registered); // isRegistered, result=true

        Log.i("testCallback","-----------------------------------------------------------");

        boolean successful2 = false;
        successful2 = CallbackAndroidGoCFunc.square(3);
        Log.i("testCallback","exec square, result="+successful2); // exec square, result=true

        Log.i("testCallback","-----------------------------------------------------------");

        boolean successful3 = false;
        successful3 =CallbackAndroidGoCFunc.sum(11,3);
        Log.i("testCallback","exec sum, result="+successful3); // exec sum, result=true

        Log.i("testCallback","-----------------------------------------------------------");

        boolean successful4 = false;
        successful4 =CallbackAndroidGoCFunc.cubic(10);
        Log.i("testCallback","exec cubic, result="+successful4); // exec cubic, result=true

        Log.i("testCallback","-----------------------------------------------------------");

        CallbackAndroidGoCFunc.unRegisterCalculatorInterface();
        Log.i("testCallback","unRegisterCalculatorInterface"); // unRegisterCalculatorInterface
        registered = CallbackAndroidGoCFunc.isRegistered();
        Log.i("testCallback","isRegistered, result="+registered); // isRegistered, result=false

        Log.i("testCallback","-----------------------------------------------------------");
    }
}
```