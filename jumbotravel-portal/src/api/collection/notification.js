import Notification from "../domain/notification";

class NotificationCollection {

    constructor({
        notifications = [],
    }) {
        this.notifications = notifications;
    }

    addLocal(notification) {
        this.notifications.push(notification);
    }

    getAll() {
        return this.notifications;
    }

    getPopup() {
        return this.notifications.copyWithin().filter(notification => {
            if (notification.isPopup()) {
                return notification;
        }});
    }

    getNotPopup() {
        return this.notifications.copyWithin().filter(notification => {
            if (!notification.isPopup()) {
                return notification;
        }});
    }

    static parse(data) {
        return new NotificationCollection({
            notifications: data.map(notification => new Notification(notification))
        });
    }

}

export default NotificationCollection;