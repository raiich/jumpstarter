# セルフレビュー

ドキュメント・コードを記述または修正した後に必ず実施する。

## 観点

対象に応じて適用する。リンク先で具体的なチェック項目を確認する。

- **ドキュメント**: [documentation.md](../guidelines/perspectives/documentation.md) — 正確性・網羅性・一貫性・読みやすさ・信頼度マーク
- **コード設計**: [design-and-coding.md](../guidelines/perspectives/design-and-coding.md) — パッケージ構成・ファイル粒度・API 設計・プロジェクト固有パターン
- **テスト**: [testing.md](../guidelines/perspectives/testing.md) — 網羅性・テスト品質・False positive
- **深刻度分類**: [review-severity.md](../guidelines/perspectives/review-severity.md) — Critical/Major/Minor/Info

## 共通チェック

対象によらず常に確認する。

- **削除可能性**: 削除しても理解できるか？ 既知情報の説明か？ 重複していないか？ 半分に削れないか？
- **批判的視点**: 「良さそう」で済ませず弱点・リスクを探す。手段と目的の整合性を確認する
- **影響分析**: 変更した要素（関数名・ファイルパス・設定キー等）を Grep で検索し、参照箇所の整合性を確認する
