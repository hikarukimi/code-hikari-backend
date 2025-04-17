--已经人工检验
-- 图书模块
CREATE TABLE books (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(100) NOT NULL,
    cover_url VARCHAR(255),
    preview_url VARCHAR(255),      -- 在线预览链接
    file_size BIGINT,              -- 文件大小（字节）
    cloud_url VARCHAR(255) NOT NULL, -- 云盘下载链接
    view_count INTEGER,
    collect_count INTEGER,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);

