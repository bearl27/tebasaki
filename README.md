# tebasaki
バックエンドのビルドと起動手順

**前提**
- Docker Desktop がインストールされ起動していること
- リポジトリルートに `.env`  を置く

## 1. `.env` の例
リポジトリルートに `.env` を作成してください。
（長張に聞いてください）

```env
# Postgres settings for app
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=examplepassword
DB_NAME=todos_db

APP_PORT=8080
```

## 2. 起動手順
ターミナルでリポジトリルートに移動して実行します。

```bash
# Docker Desktopを起動（GUI）
docker-compose up --build
```

## 3. 起動確認
- コンテナログでマイグレーションやDB接続エラーがないことを確認してください。
- エンドポイント確認:

```bash
curl http://localhost:8080/todos
```
