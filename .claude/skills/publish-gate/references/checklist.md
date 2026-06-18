# 機微情報チェックリスト

追跡対象ファイルに対する走査カテゴリ・コマンド・許可基準。各パターンはノイズ前提で、ヒットは許可対象と照合してトリアージする。

## 検査カテゴリ

| カテゴリ | 検出対象 | 許可の例 |
|---|---|---|
| 認証情報 | API キー・トークン・パスワード・秘密鍵・OAuth クライアントシークレット | 語そのものをセキュリティ概念の用語として使う場合 |
| 個人識別子 | 実名・実在メールアドレス・電話番号・アカウントハンドル | リポジトリ自身の公開 URL・規約上のプレースホルダ名 |
| ローカル状態 | `/Users/<名前>/` や `/home/<名前>/` 直下の絶対パス・ホスト名・ローカル IDE 設定 | `/Users/username` 等の明示プレースホルダ |
| 内部参照 | プライベートホスト名・社内コードネーム・顧客/雇用主名 | 分析用プロンプト中の一般的なビジネス語 |
| 実 IP | 公開 IP・RFC 1918 レンジ（`10.x` / `172.16-31.x` / `192.168.x`） | RFC 5737 / RFC 3849 の文書用レンジ |
| 実 URL | 社内 wiki・プライベートリポジトリ・未公開エンドポイント | 公開ドキュメント（ベンダ公式・RFC 割当ドメイン） |
| プレースホルダ衛生 | スラング・実在リソースの紛れ込み・他プロジェクトからのコピペ残り | 規約で定めたプレースホルダ集合 |
| dev マーカー | 未完成の文脈や名前を漏らす `TODO` / `FIXME` / `XXX` | マーカー自体を説明する記述 |

## 許可対象（誤検知になりやすい安全な値）

- IP: `192.0.2.x` / `198.51.100.x` / `203.0.113.x`（RFC 5737）、`2001:db8::/32`（RFC 3849）
- ドメイン: `example.com` / `example.org` / `example.net`（RFC 2606）
- メール: `user@example.com` / `noreply@*`
- パス: 規約上のコンテナユーザー（例 `/home/vscode`）・明示プレースホルダ
- バージョン番号や日付が IP の正規表現に当たるヒット（例 `"version": "10.0"`）

プロジェクト固有の既知安全な参照（自リポジトリの公開 URL・公式ドキュメントのドメイン等）は、プロジェクト側で許可リストを 1 箇所に保守し、ここに追記して再判定を省く。

## 走査手順

対象ファイル集合を `FILES` に取り、各パターンを走査する。範囲に応じて `FILES` の取り方を選ぶ。

```bash
# 既定: 公開対象の差分（origin/main 以降）
FILES=$(git diff --name-only --diff-filter=ACMR origin/main..HEAD)

# コミット/プッシュ直前: staged 済み変更
FILES=$(git diff --name-only --cached --diff-filter=ACMR)

# 全体監査: 追跡対象すべて
FILES=$(git ls-files)
```

```bash
# 1. 認証情報・鍵
echo "$FILES" | xargs -r grep -nEi \
  '(secret|token|api[_-]?key|password|passwd|credential|bearer|BEGIN (RSA|OPENSSH|EC|PRIVATE) (PRIVATE )?KEY)' 2>/dev/null

# 2. 個人識別子・実メール
echo "$FILES" | xargs -r grep -nE \
  '@[A-Za-z0-9.-]+\.(com|jp|net|org|io|dev|ai)\b' 2>/dev/null \
  | grep -vE '(example\.(com|org|net)|user@example|noreply@)'

# 3. ローカル絶対パス
echo "$FILES" | xargs -r grep -nE '(/Users/|/home/[a-zA-Z]+)' 2>/dev/null \
  | grep -vE '(/Users/username|/home/vscode)'

# 4. プライベート / 内部 IP
echo "$FILES" | xargs -r grep -nE \
  '\b(10\.|172\.(1[6-9]|2[0-9]|3[01])\.|192\.168\.|127\.0\.0\.|0\.0\.0\.0)' 2>/dev/null

# 5. アウトバウンド URL（ユニーク集合を許可リストと照合）
echo "$FILES" | xargs -r grep -nhoE 'https?://[a-zA-Z0-9./_-]+' 2>/dev/null | sort -u

# 6. dev マーカー
echo "$FILES" | xargs -r grep -nE '\b(TODO|FIXME|XXX|HACK|TEMP)\b' 2>/dev/null

# 7. 除外パスが実際に追跡外か確認
git check-ignore .env .env.local secrets.key config.local.json 2>/dev/null
```

## 漏出しやすい箇所

- 設定・環境定義ファイル（`name` / マウントパス / IDE バックエンド等に作者のローカル既定が残る）
- ログ出力や transcript を引用したドキュメント（ローカルパスを埋め込みやすい）
- セキュリティ概念の例示（実値へ流れやすい）

## 検出時の対処

1. 規約のプレースホルダ集合へ置換する
2. 既にコミット済みの真の秘密は、追跡から消す（`git rm` + 新コミット）だけでは不十分。source 側で当該クレデンシャルをローテーションする
3. 新たな公開参照（公式ドキュメントのドメイン等）と判明した場合は、プロジェクトの許可リストへ追記し、次回以降の走査で省く
