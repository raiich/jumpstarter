# 動作環境別カスタマイズ

## Claude Code

- **構造化ヒアリング**: `AskUserQuestion`（1ラウンド最大 4 問）
- **Web 取得**: `WebFetch`
- **対象設定の範囲（取り込み先候補）**: `.claude/` 配下で Claude Code が扱う全カテゴリ。ディレクトリ未存在のものは新規追加で対応
- **除外**: `CLAUDE.md` / `CLAUDE.local.md` への変更（本スキルは、これ以外の仕組みで代替を目指す）
- **ルール配置先**: `.claude/rules/` に直接
- **公式ドキュメント**: 下記キャッシュを `Read` し、`fetched_at` が7日以上前 or 未取得なら `WebFetch` で更新（`---\nfetched_at: YYYY-MM-DD\n---` フロントマター付きで `Write`、要約・リネーム禁止）。目次から関連個別ページも同様にキャッシュ
  - `.local/claude/cache/claude_code_docs_map.md` ← `https://code.claude.com/docs/en/claude_code_docs_map.md`

## GitHub Copilot (CLI / VSCode)

- **構造化ヒアリング**: 通常応答で選択肢を列挙する形に置き換え（1ラウンド最大 4 問のルールは維持）
- **Web 取得**: `fetch` 系の組み込み手段を使う。利用不可の場合はユーザーに URL の内容を貼ってもらう
- **対象設定の範囲（取り込み先候補）**: 環境固有の設定構造を Glob/Read で確認しユーザーへ提示してから着手
- **除外**: 環境固有のグローバル指示書は事前にユーザー確認
- **ルール配置先**: 環境の設定構造に応じて検討。判断に迷う場合はユーザーに確認
- **公式ドキュメント**: GitHub Copilot (CLI / VSCode) の公式ドキュメント（`docs.github.com` の copilot セクション / VS Code Copilot ドキュメント等）の該当ページを Web 取得しキャッシュする。鮮度7日（未取得 or `fetched_at` が7日以上前なら再取得）、`---\nfetched_at: YYYY-MM-DD\n---` フロントマター付与、要約・リネーム禁止（差分検出が壊れる）。キャッシュパスと取得すべき具体ページが不明ならユーザーに確認
