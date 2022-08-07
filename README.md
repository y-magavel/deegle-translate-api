# Deegle Translate API

## コンテナの起動方法

### 前提
- Docker Composeを使用してGoとMySQLのコンテナを起動する

### 準備
- infraディレクトリに.envファイルを作成する
- .envファイルに以下の設定を追加する
```
MYSQL_DATABASE=データベース名
MYSQL_USER=接続するユーザー名
MYSQL_PASSWORD=接続するユーザーのパスワード
MYSQL_ROOT_PASSWORD=ルートユーザーのパスワード
```

### 手順
- infraディレクトリに移動する
- `docker-compose up -d`を実行する