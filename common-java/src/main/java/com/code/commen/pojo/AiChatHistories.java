package com.code.commen.pojo;


import java.util.Date;


import lombok.Data;


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