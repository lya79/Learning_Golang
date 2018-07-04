# Golang筆記
 把學習 Golang遇到的問題進行紀錄。

# 常見問題
 - [資料排序問題(sort)](https://github.com/lya79/Learning_Golang/tree/master/SortData)
 - [struct、map、array、slice比較內容差異](https://github.com/lya79/Learning_Golang/tree/master/CompareData)
 - [複製 struct、map、array、slice](https://github.com/lya79/Learning_Golang/tree/master/CopyData)
 - 考慮到字串處理時的效率問題
 - 映射、型別判斷
 - 例外處理方式
 - [Goroutine相關問題](https://github.com/lya79/Learning_Golang/tree/master/Goroutine%E7%9B%B8%E9%97%9C%E5%95%8F%E9%A1%8C)
 - 單元測試
 - 程式運行 log紀錄
 - List、Heap、Map、Set
 - 檔案讀寫
 - JSON與 struct轉換
 - 程式運行的效率
 - [實現 callback](https://github.com/lya79/Learning_Golang/tree/master/golangCallback)
 - Golang撰寫規範

# 設計模式(Design Pattern)
 - [簡單工廠法(Simple Factory)](https://github.com/lya79/Learning_Golang/tree/master/designPattern/SimpleFactory)

# 通訊
 - NATS
 - CoAP
 - RPC
 - FTP

# 函式庫調用
 - [包裝 aar/Android調用 .aar](https://github.com/lya79/Learning_Golang/tree/master/TestNdkBuildStatic)
 - [包裝 .so/調用 .so](https://github.com/lya79/Learning_Golang/tree/master/TestNdkBuildStatic)
 - [c callback golang, golang在 callback android](https://github.com/lya79/Learning_Golang/tree/master/callbackAndroidGoCFunc)
 - [c與 golang型態轉換問題](http://colobu.com/2016/06/30/dive-into-go-10/)
 - java與 golang型態轉換問題

# 資料庫
 - SQLite使用

# 資料結構
 - 單向鏈結串列
 - 後序運算（Postfix expression)

# 經典案例
 - [生產者消費者問題(Producer-consumer problem, Bounded-buffer problem)](https://github.com/lya79/Learning_Golang/tree/master/other/ProducerConsumerProblem)
 - 哲學家就餐問題(Dining philosophers problem)
 - 理髮師問題(Sleeping-Barber Problem)
 - 八皇后問題(Eight Queens Puzzle)

# 其他
 - [字串檢查](https://github.com/lya79/Learning_Golang/tree/master/checkContent)
 <!-- 
defer與 return的順序
‍‍https://blog.csdn.net/samxx8/article/details/64442637
defer、return、返回值三者的执行顺序应该是：
1.return最先给返回值赋值；
2.接着defer开始执行一些收尾工作；
3.最后RET指令携带返回值退出函数。 
-->