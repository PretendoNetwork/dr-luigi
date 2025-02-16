package database

import (
	"time"
	"github.com/PretendoNetwork/dr-luigi/globals"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

func UploadCommonData(pid types.PID, uniqueID types.UInt64, commonData types.Buffer) error {
	now := time.Now().UnixNano()

	_, err := globals.Postgres.Exec(`
		INSERT INTO common_datas (
			owner_pid,
			unique_id,
			common_data,
			created_at
		)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (owner_pid, unique_id) DO UPDATE 
		SET common_data = excluded.common_data;`,
		pid,
		uniqueID,
		commonData,
		now,
	)

	return err
}