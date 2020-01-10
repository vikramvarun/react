import React from "react"

function MainContent() {
    const itemStyle = {
        "display": "flex",
        "justifyContent": "flex-start",
        "alignItems": "left",
        "padding": "10px 10px 0",
        "width": "40px",
        "borderBottom": "1px solid #212121",
        "fontFamily": "Roboto, sans-serif",
        "fontWeight": "100",
        "fontSize": "12px",
        "color": "#339966"
    }
    return (
        <main>
            <h2>Hello, everyone!</h2>
            <ul>
                <li style={itemStyle}>Item 1</li>
                <li style={itemStyle}>Item 2</li>
                <li style={itemStyle}>Item 3</li>
            </ul>
        </main>
    )
}

export default MainContent