
  

# 準備
<br>
- 開啟 Android Studio進行 安裝 ndk-build
<br>
SDK Manager/Appearance & Behavior/System Settings/Android SDK/SDK Tools
<br>
- 安裝 ndk-build
<br>
勾選 NDK > OK
<br>
- 設定 android sdk環境變數(使用者變數)
<br>
PATH : ANDROID_HOME | C:\Users\yuan\AppData\Local\Android\Sdk\
<br>
- 設定 javac環境變數(使用者變數)
<br>
PATH : %ANDROID_HOME%\tools\;%ANDROID_HOME%\platform-tools\; (變數值放置開頭)
<br>
PATH : C:\Program Files\Java\jdk1.8.0_172\bin; (放置結尾即可)
<br>
- 取得 gomobile
<br>
```$ go get -v golang.org/x/mobile/cmd/gomobile```
<br>
- 設定GOLANG環境變數
<br>
linux
<br>
```$ export PATH=$PATH:$GOPATH/bin```
<br>
- 初始化工具
<br>
linux
<br>
```$ gomobile init -v -ndk ~/android-ndk/ndk-bundle/```
<br>
or
<br>
windows
<br>
```$ gomobile init -v -ndk C:\Users\yuan\AppData\Local\Android\Sdk\ndk-bundle```
<br>
# 測試包裝aar
<br>
- 切換到要被包裝的套件目錄
<br>
```$ cd /home/ubuntu/workspace/go/src/golang.org/x/mobile/example/bind/hello```
<br>
- 包裝aar
<br>
```$ gomobile bind -v -target=android```
<br>
>  **Note:** 命令完成後會再目錄內產生一個 .aar檔案
<br>
# 包裝 .so檔案與 .aar提供給 Android調用
<br>
1. 編寫 .h .c
<br>
2. 修改包裝設定檔案(android.mk、application.mk)
<br>
3. 進入FromNDK目錄編譯 .h .c, 會產生 libs目錄與 obj目錄並且底下也會產生相關檔案
<br>
```ndk-build NDK_PROJECT_PATH=. NDK_APPLICATION_MK=Application.mk```
<br>
>  **Note:** 如果顯示 exec: "gcc": executable file not found in %PATH%等相關錯誤請參考下列網址除錯
<br>
https://blog.csdn.net/kingmax54212008/article/details/77188836
<br>
4. 複製 From/HelloTestFromNDKBuild.h與 libs/armeabi-v7a/libHelloTestFromNDKBuild.so, 複製至上層
<br>
5. 包裝 .aar
<br>
```gomobile bind -target=android/arm TestNdkBuildStatic```
<br>
6. 最後最上層的 .aar .so .h可供 Android使用