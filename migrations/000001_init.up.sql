-- 用户表
CREATE TABLE users (
                       id BIGSERIAL PRIMARY KEY,
                       username VARCHAR(128) NOT NULL UNIQUE,
                       email VARCHAR(256) UNIQUE,
                       nickname VARCHAR(128) NOT NULL,
                       logo VARCHAR(512),
                       password_hash VARCHAR(256) NOT NULL,
                       first_login BOOLEAN NOT NULL,
                       is_admin BOOLEAN NOT NULL,
                       created_at TIMESTAMP WITH TIME ZONE,
                       updated_at TIMESTAMP WITH TIME ZONE
);

-- 角色表
CREATE TABLE roles (
                       id BIGSERIAL PRIMARY KEY,
                       name VARCHAR(128) NOT NULL UNIQUE,
                       type INTEGER NOT NULL,
                       "desc" VARCHAR(512)
);
COMMENT ON COLUMN roles.type IS '角色类型: 1:系统角色, 2:自定义角色';

-- 团队表
CREATE TABLE teams (
                       id BIGSERIAL PRIMARY KEY,
                       name VARCHAR(128) NOT NULL UNIQUE,
                       description VARCHAR(512),
                       leader_id BIGINT DEFAULT 0 REFERENCES users(id) ON DELETE SET DEFAULT,
                       created_at TIMESTAMP WITH TIME ZONE,
                       updated_at TIMESTAMP WITH TIME ZONE
);

-- 项目表
CREATE TABLE projects (
                          id BIGSERIAL PRIMARY KEY,
                          team_id BIGINT REFERENCES teams(id) ON DELETE CASCADE,
                          name VARCHAR(128) NOT NULL,
                          description VARCHAR(512),
                          "status" VARCHAR(32) NOT NULL,
                          created_at TIMESTAMP WITH TIME ZONE,
                          updated_at TIMESTAMP WITH TIME ZONE,
                          UNIQUE(team_id, name)
);

-- 审计日志表
CREATE TABLE audits (
                        id BIGSERIAL PRIMARY KEY,
                        user_id BIGINT,
                        content VARCHAR(1024) NOT NULL,
                        created_at TIMESTAMP WITH TIME ZONE
);

-- 用户角色关联表
CREATE TABLE user_roles (
                            user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
                            role_id BIGINT,
                            PRIMARY KEY (user_id, role_id)
);

-- 团队用户关联表
CREATE TABLE team_users (
                            team_id BIGINT,
                            user_id BIGINT,
                            created_at TIMESTAMP WITH TIME ZONE,
                            PRIMARY KEY (team_id, user_id)
);

-- 项目用户关联表
CREATE TABLE project_users (
                               project_id BIGINT,
                               user_id BIGINT,
                               created_at TIMESTAMP WITH TIME ZONE,
                               PRIMARY KEY (project_id, user_id)
);

-- 创建索引
CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_teams_name ON teams(name);
CREATE INDEX idx_roles_name ON roles(name);
CREATE INDEX idx_projects_team_id ON projects(team_id);
CREATE INDEX idx_projects_team_name ON projects(team_id, name);
CREATE INDEX idx_audits_user_id ON audits(user_id);
CREATE INDEX idx_audits_created_at ON audits(created_at);