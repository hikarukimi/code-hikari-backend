-- AI聊天模块
DROP TABLE IF EXISTS ai_chat_histories;

CREATE TABLE ai_chat_histories (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL, -- 移除 REFERENCES users(id) 外键约束
    title VARCHAR(255), -- 新增标题字段
    problem TEXT NOT NULL,          -- 消息内容
    answer TEXT NOT NULL,          -- 消息内容
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);