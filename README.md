# seismicwave

地震波をファイルから読み込むライブラリパッケージです。

## Install

``` go get github.com/takatoh/seismicwave```

## Usage
読み込みたいファイルフォーマットによって関数を使い分けます。いずれの関数も seismicwave.Wave 構造体
（のポインタ）のスライスとエラー値を返します。

## Wave formats

CSV，気象庁形式，K-NET 形式，固定長の各フォーマットに対応しています。

### CSV

```go
waves, err := seismicwave.LoadCSV("filename.csv")
```

### 気象庁形式

```go
waves, err := seismicwave.LoadJMA("filename.txt")
```

### K-NET 形式

K-NET の地震波（強震動記録）のファイルは成分ごとに分かれているので、その1つを読み込む関数と、
3成分をまとめて読み込む関数があります。

```go
waves, err := seismicwave.LoadKNET("filename.NS")
```

3成分をまとめて読み込む関数では、ファイル名から拡張子を除いた部分を指定します。

```go
waves, err := seismicwave.LoadKNETSet("basename")
```

### 固定長フォーマット

固定長フォーマットではパラメータとして、ファイル名のほかに地震波名、各行のフォーマット、時間刻み、データ数、読み飛ばし行数
を指定する必要があります。この関数では、1つの地震波を読み込みます。

```go
waves, err := seismicwave.LoadFixedFormat("filename.dat", "wavename", "10F8.2", 0.01, 6000, 2)
```

または、各パラメータを TOML 式のファイルに記述して読み込むこともできます。この関数では複数の地震波を読み込むことができます。

```go
waves, err := seismicwave.LoadFixedFormatWithTOML("input.toml")
```

TOML 形式のインプットファイルは次のようになります。

```toml
[[wave]]
name   = "wave-1"
file   = "example.dat"
format = "10F8.2"
dt     = 0.01
ndata  = 6000
skip   = 2

[[wave]]
name   = "wave-2"
file   = "example.dat"
format = "10F8.2"
dt     = 0.01
ndata  = 6000
skip   = 604
```

この例では example.dat ファイルから2つの地震波を読み込んでいます。

## License

MIT License
