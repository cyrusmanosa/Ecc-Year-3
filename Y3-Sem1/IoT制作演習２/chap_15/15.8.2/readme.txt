workspaceディレクトリを環境変数GOPATHに設定してください。
下記コマンドでgomobileのインストールとソースのコンパイルができます。

// Windowsの場合
go get golang.org/x/mobile/cmd/gomobile
%GOPATH%\bin\gomobile init
%GOPATH%\bin\gomobile build mypkg

// Mac OS X、Linuxの場合
go get golang.org/x/mobile/cmd/gomobile
$GOPATH/bin/gomobile init
$GOPATH/bin/gomobile build mypkg

