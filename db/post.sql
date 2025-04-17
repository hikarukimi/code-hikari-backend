--已经人工检验
-- 帖子分类表
CREATE TABLE post_categories (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    parent_id BIGINT, -- 移除 REFERENCES post_categories(id) 外键
    created_at TIMESTAMP WITH TIME ZONE
);

-- 帖子主表
CREATE TABLE posts (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL, -- 移除 REFERENCES users(id) 外键
    category_id BIGINT, -- 移除 REFERENCES post_categories(id) 外键
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    status SMALLINT,
    view_count INTEGER,
    like_count INTEGER,
    collect_count INTEGER,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE -- 新增字段，符合 GORM 软删除规范
);

