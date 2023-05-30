import React from 'react';
import ReactDOM from 'react-dom/client';
import { StrictMode } from "react";
import MainApp from './Main';
// Import our custom CSS
import './scss/styles.scss'

// Import all of Bootstrap's JS
import * as bootstrap from 'bootstrap'

const root = ReactDOM.createRoot(document.getElementById("container"));

root.render(
  <React.StrictMode>
    <MainApp />
  </React.StrictMode>
);

