import React from "react"


function TodoItems() {
    const itemStyle = {
        "display": "flex",
        "justifyContent": "flex-start",
        "alignItems": "center",
        "padding": "30px 20px 0",
        "width": "70%",
        "borderBottom": "1px solid #cecece",
        "fontFamily": "Roboto, sans-serif",
        "fontWeight": "100",
        "fontSize": "15px",
        "color": "#333333"
    }
    return (
        <div style = {itemStyle}>
            <input type="checkbox" />
            <p>Select this option</p>
        </div>
    )

}

export default TodoItems