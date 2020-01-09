import React from "react"
//used under the hood to interpret JSX

import ReactDOM from "react-dom"
import App from "./App"

//JSX -> sort of like javascript rendition/version of html
//React-Dom.render(<h1>Hello, World!</h1>, document.getElementById("root")) -> valid
//React-Dom.render(<h1>Hello, World!</h1><p>A paragraph.</p>, document.getElementById("root")) -> invalid
ReactDOM.render(<App />, document.getElementById("root"))


//without JSX, old method
//var myNewP = document.createElement("p")
//myNewP.innerHTML = "This is a paragraph."
