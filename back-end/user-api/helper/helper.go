package helper

import (
	"backend/user-api/global"
	"context"
	"encoding/json"
	"fmt"
)

func GetUserIDFromContext(ctx context.Context) (int64, error) {

	uid, ok := ctx.Value(global.CtxJwtUserIDKey).(json.Number)
	if !ok {
		return 0, fmt.Errorf("jwt has no userID")
	}

	return uid.Int64()
}
