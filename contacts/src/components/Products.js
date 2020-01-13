import React from "react"
function Products(props) {
    return (
        <div>
            <h3> Name: {props.item.name}</h3>
            <p> Price: {props.item.price.toLocaleString("en-US", {style: "currency", currency: "INR"})}</p>
            <p> Desciption: {props.item.description}</p>
            <br />
            <hr />
        </div>
    )
}

export default Products