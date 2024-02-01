# golang-test-practice

## 練習に使ったサイト
- [twihike様](https://www.twihike.dev/docs/golang-primer/testing)
- https://andmorefine.gitbook.io/learn-go-with-tests/

### gitコマンド IssueからPRまで
ブランチ作成
```
git checkout -b feature/issue-1
```

コミット
```
git add .
git commit -m "add: issue-1"
```

プッシュ
```
git push origin feature/issue-1
```

PR作成
- GitHubのWebインターフェースを使用するか、GitHub CLIの`gh pr create`コマンドを使用してください。

PRマージ
```
git checkout main
git pull origin main
git branch -d feature/issue-1
```

ブランチ削除
```
git branch -d feature/issue-1
```

リモートブランチ削除
```
git push origin --delete feature/issue-1
```

ブランチ削除確認
```
git fetch --prune
git branch -a
```