package repository

import (
	"base-project-api/db"
	"base-project-api/models"
	"github.com/google/uuid"
)

func CreateSession(session models.CreateSessionParams) (models.CreateSessionParams, error) {
	_, err := db.Conn.NamedExec("INSERT INTO sessions (id,email,refresh_token,user_agent,client_ip,is_blocked,expires_at) VALUES (:ID, :Email, :RefreshToken, :UserAgente, :ClientIp, :IsBlocked, :ExpireAt)", session)
	if err != nil {
		return session, err
	}
	return session, nil
}

func GetSessionByID(sessionID uuid.UUID) (models.CreateSessionParams, error) {
	var sessions models.CreateSessionParams
	err := db.Conn.Get(&sessions, "SELECT * FROM sessions WHERE id = $1", sessionID)
	if err != nil {
		return sessions, err
	}
	return sessions, nil
}
