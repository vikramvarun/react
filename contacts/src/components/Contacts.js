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
            <img src={props.imgUrl} />
            <h3>{props.name}</h3>
            <p>Phone: {props.phone}</p>
            <p>Email: {props.email}</p>
            <br/>
            <hr />
        </div>)
}

export default Contacts