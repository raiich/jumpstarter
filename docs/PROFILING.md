# プロファイリング

ベンチマーク結果を詳細に分析し、ボトルネックを特定する方法を説明します。

## CPUプロファイル

```bash
# プロファイル取得
go test -bench=. -cpuprofile=cpu.prof

# コマンドラインで確認
go tool pprof cpu.prof
(pprof) top10       # 上位10個の関数を表示
(pprof) list FuncName  # 特定の関数の詳細

# Webブラウザで可視化
go tool pprof -http=:8080 cpu.prof
```

ブラウザで以下の情報が確認できます：
- **Graph**: 呼び出しグラフ
- **Flame Graph**: フレームグラフ
- **Top**: 時間を消費している関数のランキング
- **Source**: ソースコード別の時間

## メモリプロファイル

```bash
# プロファイル取得
go test -bench=. -memprofile=mem.prof

# 割り当て量で分析
go tool pprof -alloc_space mem.prof

# 割り当て回数で分析
go tool pprof -alloc_objects mem.prof

# 使用中のメモリで分析
go tool pprof -inuse_space mem.prof

# Webブラウザで可視化
go tool pprof -http=:8080 mem.prof
```

## ブロックプロファイル

並行処理の競合を分析します。

```bash
go test -bench=. -blockprofile=block.prof
go tool pprof -http=:8080 block.prof
```

チャネル操作やmutexでのブロック時間を確認できます。

## ミューテックスプロファイル

```bash
go test -bench=. -mutexprofile=mutex.prof
go tool pprof -http=:8080 mutex.prof
```

mutexの競合状況を分析できます。

## トレース取得

```bash
# 実行トレースを取得
go test -bench=. -trace=trace.out

# トレースを可視化
go tool trace trace.out
```

ブラウザでgoroutineのスケジューリング、システムコール、GCイベントなどを確認できます。

## 参考資料

- [pprof ドキュメント](https://pkg.go.dev/net/http/pprof)
- [Go execution tracer](https://pkg.go.dev/runtime/trace)
