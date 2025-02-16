package database

import (
	"github.com/PretendoNetwork/dr-luigi/globals"
)

func GetTotalRankingsByCategory(category uint32) (error, uint32) {
	var total uint32

	err := globals.Postgres.QueryRow(`SELECT COUNT(*) FROM rankings WHERE category=$1`, category).Scan(&total)
	if err != nil {
		return err, 0
	}

	return nil, total
}