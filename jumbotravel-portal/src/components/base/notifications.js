import { Component } from 'react';
import { Link } from 'react-router-dom';

import Context from '../context/app';


class Notifications extends Component {

    render() {
        return (
            <div className='z-10'>
                {/* Desktop */}
                <div className="hidden absolute sm:flex flex-col space-y-4 items-center w-full max-w-xs top-20 right-5 space-x">
                    {
                        this.context.hasNotifications && this.context.notifications ?
                            this.context.notifications.getPopup().map((notification, index) => {
                                return (
                                    <div
                                        key={index}
                                        className="flex flex-row items-center justify-center w-full"
                                    >
                                        {
                                            notification.getPopup()
                                        }
                                    </div>
                                )
                            })
                            :
                            ''
                    }
                    {
                        this.context.localNotifications.notifications.length > 0 ?
                        this.context.localNotifications.getPopup().map((notification, index) => {
                            return (
                                <div
                                    key={index}
                                    className="flex flex-row items-center justify-center w-full"
                                >
                                    {
                                        notification.getPopup()
                                    }
                                </div>
                            )
                        })
                        :
                        ''
                    }
                </div>
                {/* Mobile */}
                <div className="fixed sm:hidden w-screen flex flex-col space-y-4 items-center bottom-0 space-x pb-10">
                    {
                        this.context.hasNotifications && this.context.notifications ?
                            this.context.notifications.getPopup().map((notification, index) => {
                                return (
                                    <div
                                        key={index}
                                        className="flex flex-row items-center justify-center w-full"
                                    >
                                        {
                                            notification.getPopup()
                                        }
                                    </div>
                                )
                            })
                            :
                        ''
                    }
                    {
                        this.context.localNotifications.notifications.length > 0 ?
                        this.context.localNotifications.getPopup().map((notification, index) => {
                            return (
                                <div
                                    key={index}
                                    className="flex flex-row items-center justify-center w-full"
                                >
                                    {
                                        notification.getPopup()
                                    }
                                </div>
                            )
                        })
                        :
                        ''
                    }
                </div>
            </div>
        )
    }

}

Notifications.contextType = Context;

export default Notifications;