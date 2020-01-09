import React from "react"

import Header from "./components/Header"
import MainContent from "./components/MainContent"
import Footer from "./components/Footer"
import TodoItems from "./components/TodoItems"
import Contacts from "./components/Contacts"

function App() {
    const listStyle = {
        "backgroundColor": "white",
        "margin": "auto",
        "width": "50%",
        "display": "flex",
        "flexDirection": "column",
        "alignItems": "center",
        "border": "1px solid #efefef",
        "boxShadow": "\n        0 1px 1px rgba(0,0,0,0.15),\n            \n        0 10px 0 -5px #eee,\n            \n        0 10px 1px -4px rgba(0,0,0,0.15),\n            \n        0 20px 0 -10px #eee,\n            \n        0 20px 1px -9px rgba(0,0,0,0.15)",
        "padding": "30px"
    }
    return (
        <div>
            <Header />
            <Contacts />
            <MainContent />
            <div style={listStyle}>
                <TodoItems />
                <TodoItems />
                <TodoItems />
            </div>
            <Footer />
        </div>
    )
}

export default App