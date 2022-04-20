import { Component } from 'react';
import AppContext from "../context/app";

function classNames(...classes) {
    return classes.filter(Boolean).join(' ')
}

class PopupNotification extends Component {

    constructor(props) {
        super(props);
        // 10 seconds
        this.state = {
            show: true,
            timer: 10
        }
    }

    getSVG(severity) {

        // Return default SVG if no severity is provided
        if (!severity) {
            return (
                <svg className="w-5 h-5 text-blue-300" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="paper-plane" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                    <path fill="currentColor" d="M511.6 36.86l-64 415.1c-1.5 9.734-7.375 18.22-15.97 23.05c-4.844 2.719-10.27 4.097-15.68 4.097c-4.188 0-8.319-.8154-12.29-2.472l-122.6-51.1l-50.86 76.29C226.3 508.5 219.8 512 212.8 512C201.3 512 192 502.7 192 491.2v-96.18c0-7.115 2.372-14.03 6.742-19.64L416 96l-293.7 264.3L19.69 317.5C8.438 312.8 .8125 302.2 .0625 289.1s5.469-23.72 16.06-29.77l448-255.1c10.69-6.109 23.88-5.547 34 1.406S513.5 24.72 511.6 36.86z"></path>
                </svg>
            );
        }

        // Return SVG based on severity
        switch (severity) {
            case "INFO":
                return (
                    <svg className="w-5 h-5 text-blue-300" viewBox="0 0 24 24">
                        <path fill="currentColor" d="M11,9H13V7H11M12,20C7.59,20 4,16.41 4,12C4,7.59 7.59,4 12,4C16.41,4 20,7.59 20,12C20,16.41 16.41,20 12,20M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2M11,17H13V11H11V17Z" />
                    </svg>
                );
            case "WARNING":
                return (
                    <svg className="w-5 h-5 text-orange-300" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="exclamation-triangle" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 576 512"><path fill="currentColor" d="M569.517 440.013C587.975 472.007 564.806 512 527.94 512H48.054c-36.937 0-59.999-40.055-41.577-71.987L246.423 23.985c18.467-32.009 64.72-31.951 83.154 0l239.94 416.028zM288 354c-25.405 0-46 20.595-46 46s20.595 46 46 46 46-20.595 46-46-20.595-46-46-46zm-43.673-165.346l7.418 136c.347 6.364 5.609 11.346 11.982 11.346h48.546c6.373 0 11.635-4.982 11.982-11.346l7.418-136c.375-6.874-5.098-12.654-11.982-12.654h-63.383c-6.884 0-12.356 5.78-11.981 12.654z"></path></svg>
                );
            case "ERROR":
                return (
                    <svg className="w-5 h-5 text-red-300" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="exclamation-triangle" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 576 512"><path fill="currentColor" d="M569.517 440.013C587.975 472.007 564.806 512 527.94 512H48.054c-36.937 0-59.999-40.055-41.577-71.987L246.423 23.985c18.467-32.009 64.72-31.951 83.154 0l239.94 416.028zM288 354c-25.405 0-46 20.595-46 46s20.595 46 46 46 46-20.595 46-46-20.595-46-46-46zm-43.673-165.346l7.418 136c.347 6.364 5.609 11.346 11.982 11.346h48.546c6.373 0 11.635-4.982 11.982-11.346l7.418-136c.375-6.874-5.098-12.654-11.982-12.654h-63.383c-6.884 0-12.356 5.78-11.981 12.654z"></path></svg>
                );
            case "SUCCESS":
                return (
                    <svg className="w-5 h-5 text-green-300" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="check-circle" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path fill="currentColor" d="M504 256c0 136.967-111.033 248-248 248S8 392.967 8 256 119.033 8 256 8s248 111.033 248 248zM227.314 387.314l184-184c6.248-6.248 6.248-16.379 0-22.627l-22.627-22.627c-6.248-6.249-16.379-6.249-22.628 0L216 308.118l-70.059-70.059c-6.248-6.248-16.379-6.248-22.628 0l-22.627 22.627c-6.248 6.248-6.248 16.379 0 22.627l104 104c6.249 6.249 16.379 6.249 22.628.001z"></path></svg>
                );
            default:
                return (
                    <svg className="w-5 h-5 text-white" viewBox="0 0 24 24">
                        <path fill="currentColor" d="M11,9H13V7H11M12,20C7.59,20 4,16.41 4,12C4,7.59 7.59,4 12,4C16.41,4 20,7.59 20,12C20,16.41 16.41,20 12,20M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2M11,17H13V11H11V17Z" />
                    </svg>
                );
        }
    }

    getBackground(severity) {
        // Return SVG based on severity
        switch (severity) {
            case "INFO":
                return (
                    "bg-blue-100"
                );
            case "WARNING":
                return (
                    "bg-orange-100"
                );
            case "ERROR":
                return (
                    "bg-red-100"
                );
            case "SUCCESS":
                return (
                    "bg-green-100"
                );
            default:
                return (
                    "bg-blue-100"
                );
        }
    }

    componentDidMount() {
        this.timer = setInterval(() => {
            this.setState({
                timer: this.state.timer - 1
            });
            if (this.state.timer <= 0) {
                this.setState({
                    show: false
                });
                clearInterval(this.timer);
                // Remove the notification from the queue
                if (!this.props.local) {
                    this.context.markNotificationsRead(this.props.notification_id);
                } else {
                    this.context.localNotifications.notifications.splice(this.context.localNotifications.notifications.indexOf(this.props.notification_id), 1);
                }
            }
        }, 1000);

        // Play sound
        let audio = new Audio("/resources/beep.mp3");
        audio.play();
    }

    render() {
        return (
            <a
                className="relative flex items-center w-full max-w-xs p-4 space-x-4 text-gray-500 bg-white divide-x divide-gray-200 rounded-lg shadow space-x overflow-hidden"
                style={{
                    display: this.state.show ? 'flex' : 'none'
                }}
                href={this.props.link ? this.props.link : ''}
                target={this.props.extra && this.props.extra["target"] ? this.props.extra["target"] : ''}
                onClick={() => {
                    // Remove the notification from the queue
                    this.context.markNotificationsRead(this.props.notification_id);
                }}
            >
                {
                    this.getSVG(this.props.type)
                }
                <div className={classNames(
                    "pl-4 text-sm font-normal",
                    this.props.message ? "flex flex-col justify-center items-start" : ""
                )}>
                    <span className={classNames(
                        this.props.message ? "font-bold" : ""
                    )}>
                        {this.props.title}
                    </span>
                    <span className='text-xs'>
                        {this.props.message}
                    </span>
                </div>
                <div className="absolute w-full bottom-0 -left-4">
                    <div
                        className={classNames(
                            this.getBackground(this.props.type),
                            "h-1 transition-all duration-200 ease-in-out"
                        )}
                        style={{
                            width: `${this.state.timer * 100 / 10}%`
                        }}
                    ></div>
                </div>
            </a>
        )
    }

}

PopupNotification.contextType = AppContext;

export default PopupNotification;