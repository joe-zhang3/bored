import React, { lazy, Suspense } from "react";
import App from "./components/App"
const Header = lazy(() => import("header_module/header"));

export default function Main(){

    return (
        <div>
            <Header />
            <div className="container py-4 px-3 mx-auto">
                <App />
            </div>
        </div>
    )

}