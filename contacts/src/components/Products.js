import React from "react"
function Products(props) {
    return (
        <div>
            <h3> Name:{props.name}</h3>
            <p> Price: {props.price}</p>
            <p> Desciption: {props.description}</p>
            <br />
            <hr />
        </div>
    )
}

export default Products