import React from "react"

function Contacts(props) {
    const card = {
        "flexBasis": "250px",
        "margin": "20px"
    }
    const cardImg = {
        "width": "100%",
        "height": "auto"
    }
    const cardH = {
        "textAlign": "center"
    }
    const cardP = {
        "fontSize": "12px"
    }

    return (
        <div style={card}>
            <img src={props.contact.imgUrl} />
            <h3>{props.contact.name}</h3>
            <p>Phone: {props.contact.phone}</p>
            <p>Email: {props.contact.email}</p>
        </div>)
}

export default Contacts