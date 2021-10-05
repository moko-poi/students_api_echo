package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	// handlerパッケージのインポート
	"./handler"
)

func main() {
	// インスタンスの作成
	echo := echo.New()

	// ログ
	//HTTPのリクエストをログとってくれて、コンソールに出力
	echo.Use(middleware.Logger())
	//panicを起こしてしまっても、サーバは落とさずにエラーレスポンスを返せるようにリカバリー
	echo.Use(middleware.Recover())

	// ルーティング
	echo.GET("/student/:name", handler.SelectStudent())
	echo.GET("/student/all", handler.SelectStudents())

	//サーバの起動
	echo.Start(":1323")

}