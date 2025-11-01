# Go言語 オーバーヘッド計測ベンチマーク計画

## 目的
Go言語の各種機能がどれだけのオーバーヘッドを持つかを計測し、パフォーマンスの最適化に役立てる。

## 計測対象の機能

### 1. 関数呼び出し (function_call_bench_test.go)
- ベースライン（何もしない処理）
- 単純な関数呼び出し
- 引数ありの関数呼び出し（1個、3個、10個）
- インターフェース経由の関数呼び出し
- 直接メソッド呼び出し
- 値レシーバ vs ポインタレシーバ
  - 小さい構造体（8バイト）
  - 中程度の構造体（32バイト）
  - 大きい構造体（256バイト）
- 型埋め込み（embedding）の直接フィールドアクセス vs 昇格メソッド呼び出し
- defer付き関数 vs deferなし関数
- クロージャ呼び出し
- 再帰関数
- 可変長引数の関数
- インライン化可能な関数 vs 不可能な関数
- 関数ポインタ経由の呼び出し
- メソッド値 (method value) の呼び出し

### 2. メモリ割り当て (allocation_bench_test.go)
- スタック割り当て vs ヒープ割り当て
- 構造体の割り当て（サイズ別）
  - 小さい構造体（8バイト）
  - 中程度の構造体（32バイト）
  - 大きい構造体（256バイト）
- スライスの事前確保 vs 動的拡張
  - `make([]int, 0, 1000)` vs `make([]int, 0)` + append 1000回
- スライスの作成（make vs リテラル）
- マップの事前容量指定 vs 指定なし
  - `make(map[K]V, 1000)` vs `make(map[K]V)` + 1000要素追加
- マップの操作（追加、検索、削除）
- 文字列の連結（+ vs strings.Builder vs bytes.Buffer）
- 文字列とバイトスライスの変換（`string([]byte)` vs `[]byte(string)`）
- interface{}への代入と型アサーション
- ポインタ vs 値のコピー（サイズ別: 8バイト、32バイト、256バイト）
- 構造体のゼロ値初期化 vs フィールド指定初期化

### 3. 並行処理 (concurrency_bench_test.go)
- goroutineの起動と終了
- channelの送受信（unbuffered）
- unbuffered channel vs buffered channel（バッファサイズ10）
- 異なるバッファサイズのchannel（1, 10, 100, 1000）
- mutexのロック/アンロック（競合なし vs 競合あり）
- RWMutexの読み込みロック vs 書き込みロック
- sync.Onceの初回実行 vs 2回目以降の呼び出し
- sync.PoolのGet/Put操作
- sync.WaitGroupのAdd/Done/Wait操作
- atomic操作（Load/Store/Add/CompareAndSwap）vs mutex
- select文のオーバーヘッド（2-way, 4-way, 8-way）
- contextによるキャンセル伝播（goroutine数: 1, 10, 100）

### 4. 型変換とフォーマット (conversion_bench_test.go)
- 数値と文字列の変換
  - `strconv.Itoa` vs `fmt.Sprintf("%d", n)`
  - `strconv.ParseInt` vs `fmt.Sscanf`
  - `strconv.FormatInt` の各種ベース（2, 10, 16）
- 数値型間の変換
  - `int` → `int64`
  - `int64` → `float64`
  - `float64` → `int` (切り捨て)
- 文字列とバイトスライスの相互変換
  - `[]byte(string)` のコスト
  - `string([]byte)` のコスト
- フォーマット操作
  - `fmt.Sprintf` vs `fmt.Sprint`
  - `strings.Join` vs 手動連結（+演算子）vs strings.Builder

### 5. データ構造 (data_structure_bench_test.go)
- 配列 vs スライスのインデックスアクセス（要素数: 100）
- スライスのインデックスアクセス vs マップのキーアクセス
- スライスのコピー: `copy()` vs forループ vs `append([]T{}, slice...)`
- マップの削除: `delete()` のコスト
- スライスのイテレーション
  - `for i := 0; i < len(s); i++`
  - `for i := range s` (インデックスのみ)
  - `for _, v := range s` (値のみ)
  - `for i, v := range s` (両方)
- 構造体の比較: `==` vs `reflect.DeepEqual`

### 6. エンコーディング (encoding_bench_test.go)
- JSON
  - `json.Marshal` / `json.Unmarshal`
  - `json.Encoder` / `json.Decoder` (ストリーム処理)
  - 構造体タグの影響（タグなし vs `json:"name"` vs `json:"name,omitempty"`）
  - 構造体のサイズ別（3フィールド vs 10フィールド vs 50フィールド）
- その他のエンコーディング
  - `gob.Encode` / `gob.Decode`
  - `base64.StdEncoding.EncodeToString` / `DecodeString`
  - `hex.EncodeToString` / `DecodeString`

### 7. 時刻操作 (time_bench_test.go)
- `time.Now()` の呼び出しコスト
- `time.Since(t)` vs `time.Now().Sub(t)`
- タイムゾーン変換（`time.In(loc)` のコスト）
- `time.Format(layout)` vs `time.String()`
- タイマー操作
  - `time.NewTimer` の作成と停止
  - `time.After` のチャネル受信
  - `time.AfterFunc` のコールバック実行
- ティッカー操作
  - `time.NewTicker` の作成と停止
  - `time.Tick` のチャネル受信

### 8. コンテキスト (context_bench_test.go)
- `context.Background()` vs `context.TODO()`
- `context.WithValue` のコンテキスト作成コスト
- `context.WithCancel` / `WithTimeout` / `WithDeadline` の作成コスト
- コンテキスト値の取得（`ctx.Value(key)`）
- ネストしたコンテキストからの値取得
  - ネスト深度: 1階層 vs 5階層 vs 10階層
- コンテキストキャンセルの検出（`ctx.Done()` のselect）

### 9. その他の機能 (misc_bench_test.go)
- 型アサーション
  - 成功ケース: `v := x.(T)`
  - 失敗ケース（panicなし）: `v, ok := x.(T)`
- 型switch（case数: 2, 5, 10）
- panic/recoverのオーバーヘッド
- reflection操作
  - `reflect.TypeOf` / `reflect.ValueOf` の呼び出し
  - リフレクション経由のフィールドアクセス vs 直接アクセス
  - リフレクション経由のメソッド呼び出し vs 直接呼び出し
- エラーハンドリング
  - error返却とチェック vs panic/recover
- interface{}の操作
  - 値の代入（基本型、構造体、ポインタ）
  - 型アサーションでの取り出し
- ジェネリクス vs interface{}（Go 1.18+）
- 正規表現
  - 毎回コンパイル: `regexp.Match`
  - プリコンパイル（グローバル変数）: `regexp.Compile` + `MatchString`
  - 文字列操作: `strings.Contains` / `strings.HasPrefix`
- エラーラッピング
  - `fmt.Errorf("%w", err)`
  - `errors.Join(err1, err2)` (Go 1.20+)
- ソート
  - `sort.Slice` vs カスタムソート実装
  - `sort.Ints` vs `sort.Slice`

## ファイル構成
```
jumpstarter/
├── BENCHMARK_PLAN.md              # この計画書
├── BENCHMARK_GUIDE.md             # 実行方法とプロファイリングガイド
├── BENCHMARK_RESULTS.md           # ベンチマーク結果（実行後に作成）
├── function_call_bench_test.go    # 関数呼び出しのベンチマーク
├── allocation_bench_test.go       # メモリ割り当てのベンチマーク
├── concurrency_bench_test.go      # 並行処理のベンチマーク
├── conversion_bench_test.go       # 型変換とフォーマットのベンチマーク
├── data_structure_bench_test.go   # データ構造のベンチマーク
├── encoding_bench_test.go         # エンコーディングのベンチマーク
├── time_bench_test.go             # 時刻操作のベンチマーク
├── context_bench_test.go          # コンテキストのベンチマーク
└── misc_bench_test.go             # その他の機能のベンチマーク
```
