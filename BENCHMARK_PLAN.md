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
  - 小さい構造体（8バイト以下）
  - 中程度の構造体（32バイト）
  - 大きい構造体（256バイト以上）
- 型埋め込み（embedding）のメソッド呼び出し
- 昇格されたメソッドの呼び出しコスト
- defer付き関数 vs deferなし関数
- クロージャ呼び出し
- 再帰関数
- 可変長引数の関数
- インライン化可能な関数 vs 不可能な関数
- 関数ポインタ経由の呼び出し
- メソッド値 (method value) の呼び出し

### 2. メモリ割り当て (allocation_bench_test.go)
- スタック割り当て vs ヒープ割り当て
- 小さい構造体の割り当て vs 大きい構造体の割り当て
- スライスの事前確保 vs 動的拡張
  - `make([]int, 0, cap)` vs `make([]int, 0)` + append
- スライスの作成（make vs リテラル）
- スライスの append（容量あり vs 容量なし）
- マップの事前容量指定 vs 指定なし
  - `make(map[K]V, cap)` vs `make(map[K]V)`
- マップの操作（追加、検索、削除）
- 文字列の連結（+ vs strings.Builder vs bytes.Buffer）
- 文字列とバイトスライスの変換コスト
- インターフェース変換のコスト
- ポインタ vs 値のコピー
- 構造体のゼロ値初期化 vs フィールド指定

### 3. 並行処理 (concurrency_bench_test.go)
- goroutineの起動オーバーヘッド
- channelの送受信
- unbuffered channel vs buffered channel
- 異なるバッファサイズのchannel（1, 10, 100, 1000）
- mutexのロック/アンロック
- RWMutexの読み込みロック vs 書き込みロック
- sync.Once
- sync.Pool
- sync.WaitGroup のオーバーヘッド
- atomic操作 vs mutex
- select文のオーバーヘッド（2-way, 4-way, 8-way）
- context によるキャンセル伝播

### 4. 型変換とフォーマット (conversion_bench_test.go)
- 数値と文字列の変換
  - `strconv.Itoa` vs `fmt.Sprintf`
  - `strconv.ParseInt` vs `fmt.Sscanf`
  - `strconv.FormatInt` の各種ベース（2, 10, 16）
- 型変換のコスト
  - `int` → `int64` → `float64`
  - `[]byte` ↔ `string`
- フォーマット操作
  - `fmt.Sprintf` vs `fmt.Sprint`
  - `strings.Join` vs 手動連結

### 5. データ構造 (data_structure_bench_test.go)
- 配列 vs スライス（アクセス速度）
- スライスのアクセス vs マップのアクセス
- スライスのコピー: `copy()` vs ループ vs `append([]T{}, slice...)`
- マップの削除: `delete()` のコスト
- range によるイテレーション（インデックス vs 値）
- for文 vs range文
- 構造体の比較: `==` vs `reflect.DeepEqual`

### 6. エンコーディング (encoding_bench_test.go)
- JSON
  - `json.Marshal` / `json.Unmarshal`
  - `json.Encoder` / `json.Decoder` (ストリーム処理)
  - 構造体タグの有無による影響
  - 小さい構造体 vs 大きい構造体
- その他のエンコーディング
  - `gob.Encode` / `gob.Decode`
  - `base64.StdEncoding`
  - `hex.EncodeToString`

### 7. 時刻操作 (time_bench_test.go)
- `time.Now()` の呼び出しコスト
- `time.Since()` vs 手動計算
- タイムゾーン変換のオーバーヘッド
- `time.Format` vs `time.String`
- タイマーとティッカーの作成/破棄

### 8. コンテキスト (context_bench_test.go)
- `context.Background()` vs `context.TODO()`
- `context.WithValue` のオーバーヘッド
- `context.WithCancel` / `WithTimeout` / `WithDeadline` の作成コスト
- コンテキスト値の取得コスト
- 深くネストしたコンテキストからの値取得

### 9. その他の機能 (misc_bench_test.go)
- 型アサーション（成功 vs 失敗）
- 型switch（2-way, 5-way, 10-way）
- panic/recover のオーバーヘッド
- reflectionの使用（TypeOf, ValueOf, フィールドアクセス）
- エラーハンドリング（error返却 vs panic）
- 空インターフェースの使用
- ジェネリクス vs インターフェース（Go 1.18+）
- 正規表現: コンパイル vs プリコンパイル vs 文字列操作
- エラーラッピング: `fmt.Errorf` vs `errors.Join` (Go 1.20+)
- sort パッケージ: `sort.Slice` vs カスタム実装

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
