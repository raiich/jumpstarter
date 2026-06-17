<!-- wording-check: skip -->
# 単語チェックリスト（要検討語）

`.claude/hooks/check-wording.sh` が PostToolUse で照合する要検討語と置換案。各行 `パターン ::: 指摘文`。`exact` は固定文字列一致、`regex` は正規表現。語幹で登録する（活用形を拾う）。方針の詳細は [単語選択](writing-wording.md)。

```exact
毀損 ::: 品質が下がる／成立しなくなる 等
破綻 ::: 成り立たない
破壊 ::: 対応関係が崩れる
台無し ::: 成り立たなくなる
崩壊 ::: 整合性が失われる
xxx ::: サンプル値は foo / 192.0.2.x / 2001:db8::/32 / example.com 等
XXX ::: サンプル値は foo / 192.0.2.x / 2001:db8::/32 / example.com 等
```

```regex
という技術 ::: 「技術」はツールと区別（例: コンテナ技術（実装: Docker））
```
