-- AI聊天模块
CREATE TABLE ai_chat_histories (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL, -- 移除 REFERENCES users(id) 外键约束
    title VARCHAR(255), -- 新增标题字段
    message TEXT NOT NULL,          -- 消息内容
    file_url VARCHAR(255),          -- 文件存储路径
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);