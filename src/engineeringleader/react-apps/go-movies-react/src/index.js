import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import ErrorPage from "./components/ErrorPage";
import Home from "./components/Home"
import Movies from "./components/Movies";
import Genres from "./components/Genres";
import EditMovie from "./components/EditMovie";
import Login from "./components/Login";
import ManageCatalogue from "./components/ManageCatalogue";
import GraphQL from "./components/GraphQL";

const router = createBrowserRouter([
    {
        path: "/",
        element: <App />,
        errorElement: <ErrorPage />,
        children: [
            {index: true, element: <Home/>},
            {
                path: "/movies",
                element: <Movies/>,
            },
            {
                path: "/genres",
                element: <Genres/>,
            },
            {
                path: "/manage-catalogue",
                element: <ManageCatalogue/>,
            },
            {
                path: "/login",
                element: <Login/>,
            },
            {
                path: "/admin/movie/0",
                element: <EditMovie/>,
            },
            {
                path: "/graphql",
                element: <GraphQL/>,
            },
        ]
    }
]);

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
    <React.StrictMode>
        <RouterProvider router={router} />
    </React.StrictMode>
);