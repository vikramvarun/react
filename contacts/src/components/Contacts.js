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
            <img src={props.item.imgUrl} />
            <h3>{props.item.name}</h3>
            <p>Phone: {props.item.phone}</p>
            <p>Email: {props.item.email}</p>
            <br/>
            <hr />
        </div>)
}

export default Contacts