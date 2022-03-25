import React from 'react'
import { Helmet } from 'react-helmet';

import Sidebar from '../utils/sidebar';
import Navbar from '../utils/navbar';

class Home extends React.Component {

    render() {
        return (
            <div className=''>
                <Helmet>
                    <title>Home - JumboTravel</title>
                </Helmet>
                <Navbar app={this.props.app} config={this.props.config} />
                <Sidebar app={this.props.app} config={this.props.config} />
            </div>
        );
    }

}

export default Home;