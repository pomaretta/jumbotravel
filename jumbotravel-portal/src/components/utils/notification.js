import { Component } from "react";
import AppContext from "../context/app";

function classNames(...classes) {
    return classes.filter(Boolean).join(' ')
}

class Notification extends Component {

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

    getOutlineBackground(severity) {
        // Return SVG based on severity
        switch (severity) {
            case "INFO":
                return (
                    "outline-blue-300"
                );
            case "WARNING":
                return (
                    "outline-orange-300"
                );
            case "ERROR":
                return (
                    "outline-red-300"
                );
            case "SUCCESS":
                return (
                    "outline-green-300"
                );
            default:
                return (
                    "outline-blue-300"
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

    render() {
        return (
            <div
                className={classNames(
                    this.getOutlineBackground(this.props.type),
                    "flex | hover:shadow-md transition-all duration-200 ease-in-out | items-center w-full p-2 mb-2 text-gray-500 bg-white rounded-lg outline outline-2"
                )}
            >
                <a
                    id="toast-success"
                    className="flex | items-center | cursor-pointer"
                    role="alert"
                    href={
                        this.props.link ? this.props.link : ''
                    }
                    target={this.props.extra && this.props.extra["target"] ? this.props.extra["target"] : ''}
                >
                    <div
                        className={classNames(
                            this.getBackground(this.props.type),
                            "inline-flex items-center justify-center flex-shrink-0 w-8 h-8 text-green-500 rounded-lg"
                        )}
                    >
                        {
                            this.getSVG(this.props.type)
                        }
                    </div>
                    <div className="ml-3 text-sm font-normal">
                        {this.props.title}
                    </div>

                </a>
                <button
                    type="button"
                    className="ml-auto -my-1.5 bg-white text-gray-400 hover:text-gray-900 rounded-lg focus:ring-2 focus:ring-gray-300 p-1.5 hover:bg-gray-100 inline-flex h-8 w-8" data-dismiss-target="#toast-success" aria-label="Close"
                    onClick={() => {
                        this.context.markNotificationsRead(this.props.notification_id);
                    }}
                >
                    <span className="sr-only">Close</span>
                    <svg className="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fillRule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clipRule="evenodd"></path></svg>
                </button>
            </div>
        )
    }

}

Notification.contextType = AppContext;

export default Notification;