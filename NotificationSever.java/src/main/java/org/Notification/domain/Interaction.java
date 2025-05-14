package org.Notification.domain;

@Entity
public class Interaction {
    public enum Type { LIKE, COMMENT, FOLLOW }

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private Long sourceUserId;  // 触发者
    private Long targetUserId;   // 接收者
    private Type interactionType;
    private Long relatedEntityId; // 关联的文章/评论ID
    private LocalDateTime timestamp;
    // Getters & Setters
}
