# ベンチマークカタログ

このドキュメントでは、実装済みのベンチマーク一覧と各ベンチマークが計測する内容を説明します。

## ファイル構成

```
jumpstarter/
├── function_call_bench_test.go    # 関数呼び出し
├── allocation_bench_test.go       # メモリ割り当て
├── concurrency_bench_test.go      # 並行処理
├── conversion_bench_test.go       # 型変換とフォーマット
├── data_structure_bench_test.go   # データ構造
├── encoding_bench_test.go         # エンコーディング
├── time_bench_test.go             # 時刻操作
├── context_bench_test.go          # コンテキスト
└── misc_bench_test.go             # その他の機能
```

## 1. 関数呼び出し (function_call_bench_test.go)

関数やメソッドの呼び出しに伴うオーバーヘッドを計測します。

### 基本的な呼び出し

- **Simple**: 引数なし関数の呼び出し
- **Args1/Args3/Args10**: 引数の数による影響（1個、3個、10個）
- **Variadic**: 可変長引数の関数
- **Closure**: クロージャの呼び出し
- **Recursive**: 再帰関数（深さ10）
- **FunctionPointer**: 関数ポインタ経由の呼び出し

### インターフェースとメソッド

- **DirectMethod**: 直接メソッド呼び出し
- **InterfaceMethod**: インターフェース経由のメソッド呼び出し
- **MethodValue**: メソッド値（method value）の呼び出し

### レシーバーの種類

構造体サイズ別に値レシーバーとポインタレシーバーを比較：

- **Small（8バイト）**: ValueReceiver vs PointerReceiver
- **Medium（32バイト）**: ValueReceiver vs PointerReceiver
- **Large（256バイト）**: ValueReceiver vs PointerReceiver

### 型埋め込み

- **DirectFieldAccess**: 埋め込み型の直接フィールドアクセス
- **PromotedMethodCall**: 昇格メソッドの呼び出し

### defer とインライン化

- **WithDefer/WithoutDefer**: defer使用時のオーバーヘッド
- **Inlinable/NonInlinable**: インライン化可能/不可能な関数の比較

## 2. メモリ割り当て (allocation_bench_test.go)

メモリ割り当てのコストとパターンを計測します。

### スタックとヒープ

- **Stack**: スタック割り当て（エスケープしない変数）
- **Heap**: ヒープ割り当て（エスケープする変数）

### 構造体の割り当て

構造体サイズ別の割り当てコスト：

- **Small（8バイト）**
- **Medium（32バイト）**
- **Large（256バイト）**

### スライスの事前確保

- **Preallocated**: `make([]int, 0, 1000)` で容量を指定
- **Dynamic**: `make([]int, 0)` + append 1000回

### スライスの作成

- **MakeWithCapacity**: `make([]int, 0, 10)` での作成
- **Literal**: `[]int{}` リテラルでの作成
- **LiteralWithValues**: `[]int{1, 2, 3, ...}` 値付きリテラル

### マップの事前確保

- **Preallocated**: `make(map[int]int, 100)` で容量を指定
- **Dynamic**: `make(map[int]int)` + 100要素追加

### マップ操作

- **Insert**: 要素の挿入
- **Lookup**: 要素の検索
- **Delete**: 要素の削除

### ポインタと値のコピー

構造体サイズ別にポインタ渡しと値渡しを比較：

- **Small/Medium/Large**: Pointer vs Value

### 構造体の初期化

- **ZeroValue**: ゼロ値初期化 `SmallStruct{}`
- **FieldInit**: フィールド指定初期化 `SmallStruct{A: 1, B: 2}`

## 3. 並行処理 (concurrency_bench_test.go)

goroutine、channel、同期プリミティブのオーバーヘッドを計測します。

### Goroutine

- **Launch**: goroutineの起動と終了
- **LaunchMany**: 100個のgoroutineを起動

### Channel

- **Unbuffered**: バッファなしチャネルの送受信
- **Buffered10**: バッファサイズ10のチャネル
- **BufferSize1/10/100/1000**: バッファサイズ別の比較

### 同期プリミティブ

- **Mutex**: sync.Mutexのロック/アンロック
- **RWMutex**: 読み込みロック vs 書き込みロック
- **Once**: sync.Onceの初回実行 vs 2回目以降
- **WaitGroup**: sync.WaitGroupのAdd/Done/Wait

### Atomic操作

- **Load/Store/Add/CompareAndSwap**: atomic操作の各種命令
- **AtomicVsMutex**: atomic vs mutexの比較

### Select

- **Select2/4/8**: case数別のselectのオーバーヘッド

### その他

- **SyncPool**: sync.PoolのGet/Put操作

## 4. 型変換とフォーマット (conversion_bench_test.go)

型変換とフォーマット操作のコストを計測します。

### 数値と文字列の変換

- **Itoa**: `strconv.Itoa(n)`
- **Sprintf**: `fmt.Sprintf("%d", n)`
- **ParseInt**: `strconv.ParseInt(s, 10, 64)`
- **Sscanf**: `fmt.Sscanf(s, "%d", &result)`

### 基数変換

- **FormatInt/Base2/10/16**: 2進数、10進数、16進数への変換

### 数値型間の変換

- **IntToInt64**: `int` → `int64`
- **Int64ToFloat64**: `int64` → `float64`
- **Float64ToInt**: `float64` → `int`（切り捨て）

### 文字列とバイトスライスの変換

- **StringToBytes**: `[]byte(string)`
- **BytesToString**: `string([]byte)`

### フォーマット操作

- **Sprintf**: `fmt.Sprintf("Hello %s %d", s, n)`
- **Sprint**: `fmt.Sprint("Hello ", s, " ", n)`
- **Join**: `strings.Join(parts, " ")`
- **PlusOperator**: `+` 演算子での連結
- **Builder**: `strings.Builder` での構築

## 5. データ構造 (data_structure_bench_test.go)

配列、スライス、マップなどのデータ構造の操作コストを計測します。

### アクセス

- **ArrayAccess**: 配列のインデックスアクセス
- **SliceAccess**: スライスのインデックスアクセス
- **MapAccess**: マップのキーアクセス

### スライスのコピー

- **BuiltinCopy**: `copy()` 組み込み関数
- **ForLoop**: forループでの要素コピー
- **AppendVariadic**: `append([]int{}, src...)` での複製

### マップ操作

- **Delete**: `delete()` 関数での削除

### イテレーション

- **IndexLoop**: `for i := 0; i < len(s); i++`
- **RangeIndex**: `for i := range s`
- **RangeValue**: `for _, v := range s`
- **RangeBoth**: `for i, v := range s`

### 構造体の比較

- **EqualOperator**: `==` 演算子
- **DeepEqual**: `reflect.DeepEqual()`

## 6. エンコーディング (encoding_bench_test.go)

各種エンコーディング形式の性能を計測します。

### JSON

#### 構造体サイズ別

- **Small（3フィールド）**: Marshal/Unmarshal
- **Medium（10フィールド）**: Marshal
- **Large（50フィールド）**: Marshal

#### ストリーム処理

- **Encoder**: `json.NewEncoder().Encode()`
- **Decoder**: `json.NewDecoder().Decode()`

#### 構造体タグの影響

- **NoTags**: タグなし
- **WithTags**: `json:"name"` タグあり
- **WithOmitEmpty**: `json:"name,omitempty"` タグ

### その他のエンコーディング

- **Gob**: Encode/Decode
- **Base64**: Encode/Decode
- **Hex**: Encode/Decode

## 7. 時刻操作 (time_bench_test.go)

時刻関連の操作コストを計測します。

### 時刻取得

- **Now**: `time.Now()` の呼び出し
- **Since**: `time.Since(start)`
- **NowSub**: `time.Now().Sub(start)`

### タイムゾーン変換

- **In**: `time.In(loc)` での変換

### フォーマット

- **Format**: `time.Format(time.RFC3339)`
- **String**: `time.String()`

### タイマー

- **NewTimer**: `time.NewTimer()` の作成と停止
- **After**: `time.After()` のチャネル受信（スキップ）
- **AfterFunc**: `time.AfterFunc()` のコールバック実行

### ティッカー

- **NewTicker**: `time.NewTicker()` のティック受信
- **Tick**: `time.Tick()` のチャネル受信（スキップ）

## 8. コンテキスト (context_bench_test.go)

context パッケージの操作コストを計測します。

### コンテキストの作成

- **Background**: `context.Background()`
- **TODO**: `context.TODO()`

### WithValue

- **WithValue**: コンテキストへの値の追加
- **Value**: コンテキストからの値の取得

### キャンセル機能

- **WithCancel**: キャンセル可能なコンテキスト
- **WithTimeout**: タイムアウト付きコンテキスト
- **WithDeadline**: デッドライン付きコンテキスト

### ネストしたコンテキスト

- **Depth1/5/10**: ネスト深度別の値取得コスト

### キャンセル検出

- **DoneCheck**: `ctx.Done()` のselectチェック
- **CancelledDoneCheck**: キャンセル済みコンテキストのチェック

## 9. その他の機能 (misc_bench_test.go)

その他の言語機能のオーバーヘッドを計測します。

### 型アサーション

- **Success**: 成功する型アサーション（パニックで失敗）
- **SuccessWithCheck**: `v, ok := x.(T)` 形式（成功）
- **FailureWithCheck**: `v, ok := x.(T)` 形式（失敗）

### 型switch

- **TypeSwitch2/5/10**: case数別の型switch

### Reflection

- **TypeOf**: `reflect.TypeOf()` の呼び出し
- **ValueOf**: `reflect.ValueOf()` の呼び出し
- **FieldAccess**: リフレクション経由 vs 直接アクセス
- **MethodCall**: リフレクション経由 vs 直接呼び出し

### エラーハンドリング

- **ErrorReturn**: error返却とチェック（BenchmarkErrorHandling）
- **PanicRecover**: panic/recoverのオーバーヘッド（BenchmarkErrorHandling）
- **Errorf**: `fmt.Errorf("%w", err)` でのラップ（BenchmarkErrorWrapping）
- **Join**: `errors.Join(err1, err2)` での結合（BenchmarkErrorWrapping）

### Interface{}

- **AssignInt**: `int` → `interface{}` への代入
- **AssignStruct**: 構造体 → `interface{}` への代入
- **AssignPointer**: ポインタ → `interface{}` への代入
- **ExtractInt**: `interface{}` からの値取得

### ジェネリクス

- **Generic**: ジェネリック関数
- **Interface**: `interface{}` 経由

### 正規表現

- **CompileEachTime**: 毎回コンパイル
- **Precompiled**: プリコンパイル済み
- **StringContains**: `strings.Contains()` での代替
- **StringHasPrefix**: `strings.HasPrefix()` での代替

### ソート

- **Ints**: `sort.Ints()` 組み込み
- **Slice**: `sort.Slice()` でのソート
- **CustomSort**: カスタムソート実装

---

実行方法は [実行ガイド](RUNNING_BENCHMARKS.md) を参照してください。
