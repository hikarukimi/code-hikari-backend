-- 人工已经校验



-- 用户表（符合GORM模型规范），用于存储系统中用户的基本信息
CREATE TABLE users (
    -- 用户的唯一标识，使用自增序列生成
    id BIGSERIAL PRIMARY KEY,
    -- 用户的用户名，不能为空且必须唯一
    username VARCHAR(255) NOT NULL UNIQUE,
    -- 用户的手机号码，不能为空且必须唯一
    mobile VARCHAR(15) NOT NULL UNIQUE,
    -- 用户的电子邮箱地址，必须唯一，可以为空
    email VARCHAR(100) UNIQUE,
    -- 用户的头像地址，可以为空
    avatar VARCHAR(255),
    -- 用户的密码，不能为空
    password VARCHAR(255) NOT NULL,
    -- 用户的角色，使用小整数存储，不能为空
    -- 1: 普通用户，2: 管理员，3: 超级管理员
    role SMALLINT not null,
    -- 用户的经验值
    exp INTEGER,
    -- 用户的等级
    level INTEGER,
    -- 用户最后一次登录的时间，包含时区信息
    last_sign_in_at TIMESTAMP WITH TIME ZONE,
    -- 用户记录创建的时间，包含时区信息，不能为空
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    -- 用户记录更新的时间，包含时区信息，不能为空
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    -- 用户记录删除的时间，包含时区信息，用于软删除
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- 用户关系表（添加复合索引），用于存储用户之间的关系信息
-- relation_type 取值含义：
-- 1: 关注
-- 2: 黑名单
-- 3: 好友
CREATE TABLE user_relations (
    -- 关系记录的唯一标识，使用自增序列生成
    id BIGSERIAL PRIMARY KEY,
    -- 发起关系的用户ID
    user_id BIGINT NOT NULL,
    -- 关系目标的用户ID
    target_id BIGINT NOT NULL,
    -- 关系类型，使用小整数存储
    -- 1: 关注，2: 黑名单，3: 好友
    relation_type SMALLINT NOT NULL,
    -- 关系记录创建的时间，包含时区信息，不能为空
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    -- 关系记录更新的时间，包含时区信息，不能为空
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    -- 关系记录删除的时间，包含时区信息，用于软删除
    deleted_at TIMESTAMP WITH TIME ZONE,
    -- 确保 user_id、target_id 和 relation_type 的组合唯一
    UNIQUE(user_id, target_id, relation_type)
);

-- 用户签到记录表（添加时间索引），用于记录用户的签到信息
CREATE TABLE user_sign_ins (
    -- 签到记录的唯一标识，使用自增序列生成
    id BIGSERIAL PRIMARY KEY,
    -- 签到用户的ID
    user_id BIGINT NOT NULL,
    -- 用户签到的日期
    sign_in_date DATE NOT NULL,
    -- 签到记录创建的时间，包含时区信息，不能为空
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    -- 签到记录更新的时间，包含时区信息，不能为空
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    -- 签到记录删除的时间，包含时区信息，用于软删除
    deleted_at TIMESTAMP WITH TIME ZONE,
    -- 确保 user_id 和 sign_in_date 的组合唯一
    UNIQUE(user_id, sign_in_date)
);

-- 用户收藏表（添加类型索引），用于存储用户的收藏信息
-- target_type 取值含义：
-- 1: 文章
-- 2: 书籍
-- 3: 视频
CREATE TABLE user_favorites (
    -- 收藏记录的唯一标识，使用自增序列生成
    id BIGSERIAL PRIMARY KEY,
    -- 收藏用户的ID
    user_id BIGINT NOT NULL,
    -- 收藏目标的ID
    target_id BIGINT NOT NULL,
    -- 收藏目标的类型，使用小整数存储
    -- 1: 文章，2: 书籍，3: 视频
    target_type SMALLINT NOT NULL,
    -- 收藏记录创建的时间，包含时区信息，不能为空
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    -- 收藏记录更新的时间，包含时区信息，不能为空
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    -- 收藏记录删除的时间，包含时区信息，用于软删除
    deleted_at TIMESTAMP WITH TIME ZONE,
    -- 确保 user_id、target_id 和 target_type 的组合唯一
    UNIQUE(user_id, target_id, target_type)
);