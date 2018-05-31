
# 準備

- 開啟 Android Studio進行 安裝 ndk-build
SDK Manager/Appearance & Behavior/System Settings/Android SDK/SDK Tools

- 安裝 ndk-build
勾選 NDK > OK

- 設定 android sdk環境變數(使用者變數)
PATH : ANDROID_HOME | C:\Users\yuan\AppData\Local\Android\Sdk\

- 設定 javac環境變數(使用者變數)
PATH : %ANDROID_HOME%\tools\;%ANDROID_HOME%\platform-tools\; (變數值放置開頭)
PATH : C:\Program Files\Java\jdk1.8.0_172\bin; (放置結尾即可)

- 取得 gomobile
```$ go get -v golang.org/x/mobile/cmd/gomobile```

- 設定GOLANG環境變數
linux
```$ export PATH=$PATH:$GOPATH/bin```

- 初始化工具
linux
```$ gomobile init -v -ndk ~/android-ndk/ndk-bundle/```
or
windows
```$ gomobile init -v -ndk C:\Users\yuan\AppData\Local\Android\Sdk\ndk-bundle```
# 測試包裝aar
- 切換到要被包裝的套件目錄
```$ cd /home/ubuntu/workspace/go/src/golang.org/x/mobile/example/bind/hello```

- 包裝aar
```$ gomobile bind -v -target=android```
>  **Note:** 命令完成後會再目錄內產生一個 .aar檔案
# 包裝 .so檔案與 .aar提供給 Android調用
1. 編寫 .h .c
2. 修改包裝設定檔案
- 修改 android.mk、application.mk
4. 進入FromNDK目錄編譯 .h .c, 會產生 libs目錄與 obj目錄並且底下也會產生相關檔案
```ndk-build NDK_PROJECT_PATH=. NDK_APPLICATION_MK=Application.mk```
> **Note:** 如果顯示 exec: "gcc": executable file not found in %PATH%等相關錯誤請參考下列網址除錯
https://blog.csdn.net/kingmax54212008/article/details/77188836
4. 複製 From/HelloTestFromNDKBuild.h與 libs/armeabi-v7a/libHelloTestFromNDKBuild.so, 複製至上層
5. 包裝 .aar
```gomobile bind -target=android/arm TestNdkBuildStatic```
6. 最後最上層的 .aar .so .h可供 Android使用