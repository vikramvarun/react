import React from "react"

import Header from "./components/Header"
import MainContent from "./components/MainContent"
import TodoItems from "./components/TodoItems"
import Footer from "./components/Footer"

function App() {

    const header = {}
    const mainContent = {}
    const footer = {}

    const todoList = {
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
            <div style={header}>
                <Header />
            </div>

            <div style={mainContent}>
                <MainContent />
            </div>

            <div style={todoList}>
                <TodoItems />
                <TodoItems />
                <TodoItems />
            </div>

            <div style={footer}>
                <Footer />
            </div>

        </div>
    )
}

export default App