// 订阅个人通知
stompClient.subscribe('/user/queue/private', (message) => {
    handlePrivateMessage(JSON.parse(message.body));
});

stompClient.subscribe('/user/queue/notifications', (message) => {
    handleInteraction(JSON.parse(message.body));
});

function handlePrivateMessage(message) {
    // 更新私信列表
}

function handleInteraction(notification) {
    const messages = {
        LIKE: `${user} 赞了你的内容`,
        COMMENT: `${user} 评论了你的帖子`,
        FOLLOW: `${user} 关注了你`
    };
    showNotification(messages[notification.type]);
}
