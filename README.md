# sample-go-testing

## 参考
https://swet.dena.com/entry/2018/01/16/211035

## 実行コマンド

### go test
* `_test` である場合
```
$ go test xxx_test.go
```

* 同一パッケージである場合
```
$ go test xxx.go xxx_test.go
```

* 相対パス指定
```
$ go test ./xxx
```

* 詳細確認
```
$ go test -v xxx_test.go
```

* 特定のテスト実行
`--run`で指定した文字列を含むテストを実行。

```
$ go test -v {dir}/file_test.go --run hoge
``` 
## gomock
```
$ mockgen -source={file}.go -destination {dis}/{file}.go
```
`-source`: インタフェースが含まれるソースを指定  
`-destination`: 出力先
