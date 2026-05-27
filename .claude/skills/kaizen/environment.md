# 動作環境別カスタマイズ

## Claude Code

- **構造化ヒアリング**: `AskUserQuestion`（1ラウンド最大 4 問）
- **対象設定の範囲**: `.claude/` 配下
- **置き場所候補**:

  | 置き場所 | 適する用途 | リスク |
  |---|---|---|
  | `rules/` | 必ず守らせたい禁止/必須事項 | コンテキスト圧迫 |
  | `skills/` | 繰り返し作業の標準化、skill 固有の観点・テンプレ | 起動忘れ |
  | `agents/` | 大規模探索、専門レビュー | 定義コスト |
  | `hooks/` | 検証、permission 削減、整形 | デバッグ困難 |
  
- **常時参照型の手段の扱い**: `rules/` は常時ロードでコンテキストを圧迫する。スキル化・フック化など必要時参照型で代替できないか先に検討
- **最新仕様の把握**: 下記キャッシュを `Read` し、`fetched_at` が7日以上前 or 未取得なら `curl -sL <URL>` で更新（`---\nfetched_at: YYYY-MM-DD\n---` フロントマター付きで `Write`、要約・リネーム禁止）
  - `.local/claude/cache/claude_code_docs_map.md` ← `https://code.claude.com/docs/en/claude_code_docs_map.md`
  - `.local/claude/cache/changelog.md` ← `https://code.claude.com/docs/en/changelog.md`

## GitHub Copilot (CLI / VSCode)

- **構造化ヒアリング**: 通常応答で選択肢を列挙する形に置き換え（1ラウンド最大 4 問のルールは維持）
- **対象設定の範囲**: 環境固有の設定構造を Glob/Read で確認しユーザーへ提示してから着手
- **置き場所候補**: 環境固有の設定ファイルから、強制力（常時参照 / オンデマンド）と起動条件で分類して提示
- **常時参照型の手段の扱い**: 常時参照される指示書はコンテキストを圧迫するため、オンデマンド参照可能な手段を優先検討。特に `.github/instructions/*.instructions.md` の `applyTo:` 範囲は最小化を検討する
- **最新仕様の把握**: GitHub Copilot (CLI / VSCode) の公式ドキュメント（`docs.github.com` の copilot セクション等）の該当ページを Web 取得しキャッシュする。
  鮮度7日（未取得 or `fetched_at` が7日以上前なら再取得）。
  `---\nfetched_at: YYYY-MM-DD\n---` フロントマター付与、要約・リネーム禁止（差分検出が壊れる）。
  キャッシュパスと取得すべき具体ページが不明ならユーザーに確認
