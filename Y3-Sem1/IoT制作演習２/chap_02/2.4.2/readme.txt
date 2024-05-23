workspaceディレクトリを環境変数GOPATHに設定してください。
下記コマンドでコンパイルできます。（pkg4はエラーになります）

go install pkg1
go install pkg1/pkg2
go install pkg4

