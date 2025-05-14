package com.code.commen.pojo;


import java.util.Date;

import com.baomidou.mybatisplus.annotation.TableName;
import lombok.Data;

/**
 * @TableName ai_chat_histories
 */
@TableName(value ="ai_chat_histories")
@Data
public class AiChatHistories {
    private Long id;

    private Long userId;

    private String title;

    private String problem;

    private String answer;

    private Date createdAt;

    private Date updatedAt;

    private Date deletedAt;
}