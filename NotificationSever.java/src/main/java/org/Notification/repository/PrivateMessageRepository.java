package org.Notification.repository;

import org.Notification.domain.PrivateMessage;

public interface PrivateMessageRepository extends JpaRepository<PrivateMessage, Long> {
    List<PrivateMessage> findByRecipientIdOrderByTimestampDesc(Long recipientId);
}
