package auth

import (
	"github.com/gin-gonic/gin"
	"item-manager-new/internal/errors/business"
	"item-manager-new/internal/pkg/global"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// Manager 管理会话生命周期。
type Manager struct {
	store      Store
	ttl        time.Duration
	cookieName string
}

// ManagerOption 会话管理器选项。
type ManagerOption func(*Manager)

// NewStoreManager 创建 Manager。
func NewStoreManager(store Store) *Manager {
	cfg := global.GetAuthConfig()
	m := &Manager{
		store:      store,
		ttl:        cfg.SessionTTL,
		cookieName: cfg.CookieName,
	}
	return m
}

// SetSession 创建新会话。
func (m *Manager) SetSession(c *gin.Context, userID int64) error {
	id := uuid.NewString()
	session := &Session{
		ID:        id,
		UserID:    userID,
		CreatedAt: time.Now().UTC(),
		ExpiresAt: time.Now().UTC().Add(m.ttl),
	}
	if err := m.store.Save(session); err != nil {
		return err
	}
	// 写入 Cookie
	m.writeSessionCookie(c, session)
	return nil
}

// GetSession 根据 ID 获取会话，过期则删除。
func (m *Manager) GetSession(id string) (*Session, error) {
	session, err := m.store.Get(id)
	if err != nil {
		return nil, err
	}
	if time.Now().UTC().After(session.ExpiresAt) {
		_ = m.store.Delete(id)
		return nil, business.ErrSessionNotFound
	}
	return session, nil
}

// DeleteSession 失效会话。
func (m *Manager) DeleteSession(id string) error {
	return m.store.Delete(id)
}

// CookieName 返回 Cookie 名称。
func (m *Manager) CookieName() string {
	return m.cookieName
}

func (m *Manager) writeSessionCookie(c *gin.Context, session *Session) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     m.cookieName,
		Value:    session.ID,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  session.ExpiresAt,
	})
	c.Set(session.UserID, session.ID)
}
