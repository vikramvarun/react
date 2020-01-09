import React from "react"

function Contacts(props) {
    const list = {
        "display": "flex",
        "flexWrap": "wrap"
    }

    const cards = {
        "flexBasis": "250px",
        "margin": "20px"
    }

    const cardImg = {
        "flexBasis": "250px",
        "margin": "20px",
        "width": "100%",
        "height": "auto"
    }
    const cardH = {
        "flexBasis": "250px",
        "margin": "20px",
        "textAlign": "center"
    }
    const cardP = {
        "flexBasis": "250px",
        "margin": "20px",
        "fontSize": "12px"
    }

    return (
        <div style={list}>
            <img src={props.imgUrl}/>
            <h3>{props.name}</h3>
            <p>Phone: {props.phone}</p>
            <p>Email: {props.email}</p>
        </div>
    )
}

export default Contacts