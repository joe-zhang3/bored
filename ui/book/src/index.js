import React from 'react';
import ReactDOM from 'react-dom/client';
import Main from './Main';
// Import our custom CSS
import './scss/styles.scss'
// Import all of Bootstrap's JS
import 'bootstrap/js/dist/popover';
import 'bootstrap/dist/js/bootstrap.bundle';

const root = ReactDOM.createRoot(document.getElementById("container"));

root.render(
  <React.StrictMode>
    <Main />
  </React.StrictMode>
);

