-- 通知模块
CREATE TABLE notifications (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    type ENUM('BROADCAST','PRIVATE','INTERACTION') NOT NULL,
    title VARCHAR(100),
    content TEXT NOT NULL,
    recipient_id BIGINT NOT NULL, -- 接收者ID
    sender_id BIGINT, -- 可为空（如系统广播）
    related_id BIGINT, -- 关联对象ID（如帖子ID）
    is_read BOOLEAN DEFAULT FALSE,
    expired_at DATETIME, -- 过期时间
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX (recipient_id, is_read),
    FOREIGN KEY (recipient_id) REFERENCES users(id)
);
