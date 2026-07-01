package main

import (
	"log"
	"net/http"

	"code.riskrancher.com/RiskRancher/core/pkg/datastore"
	"code.riskrancher.com/RiskRancher/core/pkg/server"
	"code.riskrancher.com/RiskRancher/core/ui"
)

var (
	BuildVersion = "dev"
	BuildCommit  = "none"
)

func main() {
	ui.SetVersionInfo(BuildVersion, BuildCommit)

	db := datastore.InitDB("./data/RiskRancher.db")

	defer db.Close()

	store := datastore.NewSQLiteStore(db)

	app := server.NewApp(store)

	server.RegisterRoutes(app)

	log.Println("🤠 RiskRancher Core Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
