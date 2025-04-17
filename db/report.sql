--已经人工检验
-- 举报系统表
CREATE TABLE reports (
    id BIGSERIAL PRIMARY KEY,
    reporter_id BIGINT NOT NULL, -- 移除 REFERENCES users(id) 外键
    target_id BIGINT NOT NULL,      -- 被举报对象ID
    target_type SMALLINT NOT NULL,  -- 1:用户 2:帖子 3:评论 4:图书
    reason TEXT NOT NULL,
    status SMALLINT, -- 移除默认值 1
    handler_id BIGINT, -- 移除 REFERENCES users(id) 外键
    created_at TIMESTAMP WITH TIME ZONE, -- 移除默认值
    updated_at TIMESTAMP WITH TIME ZONE, -- 移除默认值
    deleted_at TIMESTAMP WITH TIME ZONE -- 新增字段，用于 GORM 软删除
);