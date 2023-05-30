import React from 'react';
import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import * as bootstrap from 'bootstrap/dist/js/bootstrap.bundle';

import "./scss/styles.scss"

import Main from "./Main";

const rootElement = document.getElementById("container");
const root = createRoot(rootElement);

root.render(
  <StrictMode>
    <Main />
  </StrictMode>
);

