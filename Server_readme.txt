# API サーバー仕様

- クライアントと通信するメッセージの形式として JSON を使用する
  - `Content-Type` ヘッダに `application/json` が指定されていないリクエストは受け付けない
- API は `/api` 以下のパスで提供される
- すべての API は POST メソッドのみ受けつける
- すべての API は冪等性を持つ
- 詳細な通信仕様は [OpenAPI 定義](../swagger-ui/api.yaml) に従う
  - OpenAPI 定義は Swagger UI で確認できる

## ユーザー認証

- ユーザー認証 API が実行されるとサーバーは無期限の認証トークン ([JSON Web Token](https://www.rfc-editor.org/rfc/rfc7519)) を返す
- クライアントは `Authorization` ヘッダに `Bearer` スキームを指定してトークンを与えることでユーザー認証が必要な API を実行できる
  - `Authorization: Bearer ${token}`
- ユーザーは単一のセッションしか保持できない
  - 同一ユーザーで新たにユーザー認証 API が実行された場合、それ以前に発行された認証トークンは使用できなくなる

## 管理者認証

- 管理者認証 API が実行されるとサーバーは無期限の認証トークン ([JSON Web Token](https://www.rfc-editor.org/rfc/rfc7519)) を返す
- クライアントは `Authorization` ヘッダに `Bearer` スキームを指定してトークンを与えることで管理者認証が必要な API を実行できる
  - `Authorization: Bearer ${token}`
- 管理者アカウントは単一のセッションしか保持できない
  - 同一の管理者アカウントで新たに管理者認証 API が実行された場合、それ以前に発行された認証トークンは使用できなくなる
