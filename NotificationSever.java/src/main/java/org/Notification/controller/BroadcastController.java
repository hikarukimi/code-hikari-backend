package org.Notification.controller;

@Controller
public class BroadcastController {
    @Autowired
    private SimpMessagingTemplate messagingTemplate;

    @MessageMapping("/broadcast")
    @SendTo("/topic/broadcast")
    public String sendBroadcast(String message) {
        return "[全站通知] " + message;
    }

    // HTTP接口触发广播（可选）
    @PostMapping("/admin/broadcast")
    public ResponseEntity<?> sendBroadcastViaHttp(@RequestBody String message) {
        messagingTemplate.convertAndSend("/topic/broadcast", "[Admin] " + message);
        return ResponseEntity.ok().build();
    }
}
