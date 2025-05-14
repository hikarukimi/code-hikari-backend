package org.Notification.service;

@Service
public class PrivateMessageService {
    @Autowired
    private PrivateMessageRepository messageRepository;
    @Autowired
    private SimpMessagingTemplate messagingTemplate;

    public void sendMessage(Long senderId, Long recipientId, String content) {
        PrivateMessage message = new PrivateMessage();
        message.setSenderId(senderId);
        message.setRecipientId(recipId);
        message.setContent(content);
        message.setTimestamp(LocalDateTime.now());
        messageRepository.save(message);

        // 实时推送
        messagingTemplate.convertAndSendToUser(
                recipientId.toString(),
                "/queue/private",
                message
        );
    }
}
