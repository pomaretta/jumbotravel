import DefaultNotification from '../../components/utils/notification';
import PopupNotification from '../../components/utils/popup';

class Notification {

    constructor({
        notification_id,
        scope,
        resource_id,
        resource_uuid,
        title,
        message,
        link,
        extra,
        type,
        popup,
        expires_at,
        created_at,
        seen,
        active,
        signature,
        local = false,
    }) {
        this.notification_id = notification_id;
        this.scope = scope;
        this.resource_id = resource_id;
        this.resource_uuid = resource_uuid;
        this.title = title;
        this.message = message;
        this.link = link;
        this.extra = extra;
        // Parse JSON extra
        try {
            this.extra = JSON.parse(extra);
        } catch (e) {
            this.extra = {};
        }
        this.type = type;
        this.popup = popup;
        this.expires_at = expires_at;
        this.created_at = created_at;
        this.seen = seen;
        this.active = active;
        this.signature = signature;
        this.local = local;
    }

    getId() {
        return this.notification_id;
    }

    isPopup() {
        return this.popup;
    }

    isActive() {
        return this.active;
    }

    isSeen() {
        return this.seen;
    }

    isExpired() {
        // Check if notification is expired with UTC time
        var now = new Date();
        var utc_timestamp = Date.UTC(now.getUTCFullYear(), now.getUTCMonth(), now.getUTCDate(),
            now.getUTCHours(), now.getUTCMinutes(), now.getUTCSeconds(), now.getUTCMilliseconds());
        return Date.parse(this.expires_at) < utc_timestamp;
    }

    getActive() {
        return this.active;
    }

    setActive(active) {
        this.active = active;
    }

    getSeen() {
        return this.seen;
    }

    setSeen(seen) {
        this.seen = seen;
    }

    isLocal() {
        return this.local;
    }

    getNotification() {
        return <DefaultNotification {...this} />;
    }

    getPopup() {
        return <PopupNotification {...this} />;
    }
    
    update(notification) {

        // Title has changed
        if (this.title !== notification.title) {
            this.title = notification.title;
        }

        // Message has changed
        if (this.message !== notification.message) {
            this.message = notification.message;
        }

        // Link has changed
        if (this.link !== notification.link) {
            this.link = notification.link;
        }

        // Extra has changed
        if (this.extra !== notification.extra) {
            this.extra = notification.extra;
        }

        // Type has changed
        if (this.type !== notification.type) {
            this.type = notification.type;
        }

        // Popup has changed
        if (this.popup !== notification.popup) {
            this.popup = notification.popup;
        }
        
        // Seen has changed
        if (this.seen !== notification.seen) {
            this.seen = notification.seen;
        }

        // Active has changed
        if (this.active !== notification.active) {
            this.active = notification.active;
        }
        
        // Signature has changed
        if (this.signature !== notification.signature) {
            this.signature = notification.signature;
        }

        return;
    }

}

export default Notification;