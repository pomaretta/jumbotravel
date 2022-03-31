import DefaultNotification from '../../components/utils/notification';
import PopupNotification from '../../components/utils/popup';

class Notification {

    constructor({
        notification_id,
        scope,
        resource_id,
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
    }) {
        this.notification_id = notification_id;
        this.scope = scope;
        this.resource_id = resource_id;
        this.title = title;
        this.message = message;
        this.link = link;
        this.extra = extra;
        this.type = type;
        this.popup = popup;
        this.expires_at = expires_at;
        this.created_at = created_at;
        this.seen = seen;
        this.active = active;
        this.signature = signature;
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

    getNotification() {
        return <DefaultNotification {...this} />;
    }

    getPopup() {
        return <PopupNotification {...this} />;
    }

}

export default Notification;