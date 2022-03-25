import { Component } from "react";
import { Helmet } from "react-helmet";

class NotFound extends Component {
    render() {
        return (
            <div>
                <Helmet>
                    <title>Not Found - JumboTravel</title>
                </Helmet>
                <h1>404 - Not Found</h1>
                <p>Path: {document.location.href}</p>
            </div>
        );
    }
}

export default NotFound;