import React from 'react';
import ReactDOM from 'react-dom';
import { CookiesProvider } from 'react-cookie';
import "./styles/_general.scss";

import App from './components/app';

const config = {
  schema: process.env.REACT_APP_SCHEMA || 'http',
  hostname: process.env.REACT_APP_HOSTNAME || 'localhost:3000',
  environment: process.env.REACT_APP_ENVIRONMENT || 'DEV',
};

const appEl = document.getElementById('root');

ReactDOM.render(
  <React.StrictMode>
    <CookiesProvider>
      <App config={config} />
    </CookiesProvider>
  </React.StrictMode>,
  appEl
);