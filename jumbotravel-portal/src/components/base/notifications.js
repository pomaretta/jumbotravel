import { Component } from 'react';
import { Link } from 'react-router-dom';

class Notification extends Component {

    constructor(props) {
        super(props);
        this.state = {
            show: true,
            timer: 10
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
            }
        }, 1000);
    }

    render() {
        // TODO: Change icons with props
        // TODO: When notification is out, remove it from the list
        return (
            <Link
                className="relative flex items-center w-full max-w-xs p-4 space-x-4 text-gray-500 bg-white divide-x divide-gray-200 rounded-lg shadow dark:text-gray-400 dark:divide-gray-700 space-x dark:bg-gray-800 overflow-hidden"
                style={{
                    display: this.state.show ? 'flex' : 'none'
                }}
                to={this.props.to ? this.props.to : ''}
            >
                <svg className="w-5 h-5 text-blue-600 dark:text-blue-500" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="paper-plane" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path fill="currentColor" d="M511.6 36.86l-64 415.1c-1.5 9.734-7.375 18.22-15.97 23.05c-4.844 2.719-10.27 4.097-15.68 4.097c-4.188 0-8.319-.8154-12.29-2.472l-122.6-51.1l-50.86 76.29C226.3 508.5 219.8 512 212.8 512C201.3 512 192 502.7 192 491.2v-96.18c0-7.115 2.372-14.03 6.742-19.64L416 96l-293.7 264.3L19.69 317.5C8.438 312.8 .8125 302.2 .0625 289.1s5.469-23.72 16.06-29.77l448-255.1c10.69-6.109 23.88-5.547 34 1.406S513.5 24.72 511.6 36.86z"></path></svg>
                <div className="pl-4 text-sm font-normal">
                    {this.props.message}
                </div>
                <div className="absolute w-full bottom-0 -left-4">
                    <div
                        className="h-1 bg-jt-primary transition-all duration-200 ease-in-out"
                        style={{
                            width: `${this.state.timer * 100 / 10}%`
                        }}
                    ></div>
                </div>
            </Link>
        )
    }

}

class Notifications extends Component {

    render() {
        return (
            <div>
                {/* Desktop */}
                <div className="hidden absolute sm:flex flex-col space-y-4 items-center w-full max-w-xs top-20 right-5 space-x">
                    <Notification message="This is a notification" to="/planes" />
                </div>
                {/* Mobile */}
                <div className="fixed sm:hidden w-screen flex flex-col space-y-4 items-center bottom-0 space-x pb-10">
                </div>
            </div>
        )
    }

}

export default Notifications;