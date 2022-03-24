import React from 'react';
import ReactDOM from 'react-dom';
import { CookiesProvider } from 'react-cookie';
import "./styles/_general.scss";

// Configuration
import config from './config/config.json';

// Main Component
import App from './components/app';

const appEl = document.getElementById('root');

ReactDOM.render(
  <React.StrictMode>
    <CookiesProvider>
      <App config={config} />
    </CookiesProvider>
  </React.StrictMode>,
  appEl
);