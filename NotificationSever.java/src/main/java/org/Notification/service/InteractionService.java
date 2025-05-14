package org.Notification.service;

@Service
public class InteractionService {
    @Autowired
    private InteractionRepository interactionRepo;
    @Autowired
    private SimpMessagingTemplate messagingTemplate;

    public void handleInteraction(Type type, Long sourceUser, Long targetUser, Long entityId) {
        Interaction interaction = new Interaction();
        interaction.setInteractionType(type);
        interaction.setSourceUserId(sourceUser);
        interaction.setTargetUserId(targetUser);
        interaction.setRelatedEntityId(entityId);
        interaction.setTimestamp(LocalDateTime.now());
        interactionRepo.save(interaction);

        // 推送提醒
        NotificationDTO dto = new NotificationDTO(type, sourceUser, entityId);
        messagingTemplate.convertAndSendToUser(
                targetUser.toString(),
                "/queue/notifications",
                dto
        );
    }
}
