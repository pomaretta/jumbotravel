import React from "react";
import {
    BrowserRouter as Router,
    Routes,
    Route
} from "react-router-dom"

class AppRouter extends React.Component {

    defineRoute({
        route,
        config
    }) {
        return <Route
            element={<route.component app={this.props.app} config={config} />}
            {...route}
        />
    }

    defineRoutes({
        routes,
        config
    }) {
        return routes.map(route => this.defineRoute({
            route: route,
            config: config
        }));
    }

    render() {
        return (
            <Router>
                <Routes>
                    {this.defineRoutes({
                        routes: this.props.app.routes,
                        config: this.props.app.config
                    })}
                </Routes>
            </Router>
        );
    }

}

export default AppRouter;
