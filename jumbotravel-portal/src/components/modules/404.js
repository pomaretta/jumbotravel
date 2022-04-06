import { Component } from "react";
import { Helmet } from "react-helmet";

class NotFound extends Component {
    render() {
        return (
            <div className="w-screen h-screen">
                <Helmet>
                    <title>Not Found - JumboTravel</title>
                </Helmet>
                <div className="relative | w-full h-full | flex flex-col items-center justify-center">
                    <h1 className="font-bold text-9xl mb-5">
                        <span className="text-brand-blue">4</span>
                        <span className="text-brand-green">0</span>
                        <span className="text-brand-blue">4</span>
                    </h1>
                    <h2 className="text-2xl font-bold text-gray-800 mb-3">Oops!</h2>
                    <p className="text-xl text-gray-600">
                        We can't find the page you're looking for.
                    </p>
                    <p className="mt-2">
                        Return to the 
                        <a
                            href="/"
                            className="text-brand-blue underline font-bold pl-2">
                            homepage
                        </a>.
                    </p>
                </div>
            </div>
        );
    }
}

export default NotFound;