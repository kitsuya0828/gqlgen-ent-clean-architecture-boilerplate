package ent

import (
	"context"
	"fmt"

	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/ent/schema/ulid"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/pkg/const/globalid"
)

var globalIDs = globalid.New()

// IDToType maps an ulid.ID to the underlying table.
func IDToType(_ context.Context, id ulid.ID) (string, error) {
	if len(id) < 3 {
		return "", fmt.Errorf("IDToType: id too short")
	}
	prefix := id[:3]
	t, err := globalIDs.FindTableByID(string(prefix))
	if err != nil {
		return "", fmt.Errorf("IDToType: failed to map prefix '%s' to a type", prefix)
	}
	return t, nil
}