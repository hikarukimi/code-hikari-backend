--已经人工检验
-- 评论系统表
CREATE TABLE comments (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,  -- 移除 REFERENCES users(id) 外键
    target_id BIGINT NOT NULL,      -- 被评论对象ID
    target_type SMALLINT NOT NULL,  -- 1:帖子 2:图书 3:面试题 4:评论
    content TEXT NOT NULL,
    parent_id BIGINT, -- 移除 REFERENCES comments(id) 外键
    like_count INTEGER,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE -- 新增 deleted_at 字段用于 GORM 软删除
);

-- 评论点赞记录表
CREATE TABLE comment_likes (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,  -- 移除 REFERENCES users(id) 外键
    comment_id BIGINT NOT NULL,  -- 移除 REFERENCES comments(id) 外键
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE, -- 新增 deleted_at 字段用于 GORM 软删除
    UNIQUE(user_id, comment_id)
);

-- @提及系统表
CREATE TABLE comment_mentions (
    id BIGSERIAL PRIMARY KEY,
    comment_id BIGINT NOT NULL,  -- 移除 REFERENCES comments(id) 外键
    mentioned_user_id BIGINT NOT NULL,  -- 移除 REFERENCES users(id) 外键
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE -- 新增 deleted_at 字段用于 GORM 软删除
);