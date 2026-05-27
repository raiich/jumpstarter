# 動作環境別カスタマイズ

## Claude Code

- **並列探索**: `Agent` ツール（`subagent_type=Explore`）を `breadth` パラメータで範囲指定して並列に

## GitHub Copilot (CLI / VSCode)

- **並列探索**: 組み込みエージェント `Explore` を `agent` ツールで並列起動。シンプルな範囲では `grep` / `glob` で逐次調査。広範な探索になりそうな場合は対象範囲をユーザーへ早めに確認
