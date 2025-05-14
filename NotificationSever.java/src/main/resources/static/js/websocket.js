const socket = new SockJS('/ws');
const stompClient = Stomp.over(socket);

stompClient.connect({}, () => {
    stompClient.subscribe('/topic/broadcast', (message) => {
        showNotification(JSON.parse(message.body));
    });
});

function showNotification(content) {
    // 在页面显示通知逻辑
}
