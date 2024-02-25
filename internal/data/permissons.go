package data

import (
	"context"
	"database/sql"
	"time"
)

type Permission []string

func (p Permission) Include(code string) bool {
	for i := range p {
		if code == p[i] {
			return true
		}
	}
	return false
}

type PermissionModel struct {
	DB *sql.DB
}

func (m *PermissionModel) GetAllForUser(userID int64) (Permission, error) {
	query := `
		SELECT permissions.code
		FROM permissions
		INNER JOIN users_permissions ON users_permissions.permission_id = permissions.id
		INNER JOIN users ON users_permissions.user_id = users.id
		WHERE users.id = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := m.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var permissions Permission
	for rows.Next() {
		var permisson string
		err := rows.Scan(&permisson)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permisson)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return permissions, nil
}
