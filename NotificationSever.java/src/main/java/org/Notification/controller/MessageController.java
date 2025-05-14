package org.Notification.controller;

@RestController
@RequestMapping("/api/messages")
public class MessageController {
    @Autowired
    private PrivateMessageService messageService;

    @PostMapping("/send")
    public ResponseEntity<?> sendMessage(@RequestBody MessageRequest request) {
        messageService.sendMessage(
                request.getSenderId(),
                request.getRecipientId(),
                request.getContent()
        );
        return ResponseEntity.ok().build();
    }
}
