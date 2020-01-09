import React from "react"

function Header() {
    const firstName = "Fname"
    const lastName = "Lname"
    const navbar = {
        "height": "100px",
        "backgroundColor": "#333",
        "color": "whitesmoke",
        "marginBottom": "15px",
        "textAlign": "center",
        "fontSize": "30px",
        "display": "flex",
        "justifyContent": "center",
        "alignItems": "center"
      }
    return (
        <header style={navbar}>
            This is the header for {`${firstName} ${lastName}`}
        </header>
    )
}

export default Header