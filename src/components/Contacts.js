import React from "react"

function Contacts() {
    const list = {
        "display": "flex",
        "flexWrap": "wrap"
    }

    const cards = {
        "flexBasis": "250px",
        "margin": "20px"
    }

    const cardImg = {
        "width": "100%",
        "height": "auto"
    },
    const cardH3 = {
        "textAlign": "center"
    },
    const cardP = {
        "fontSize": "12px"
    }

    return (
        <div style={list}>
            <div style={cards}>
                <img style={cardImg} src="http://placekitten.com/300/200" />
                <h3 style={cardH3} >Mr. Whiskerson</h3>
                <p style={cardP}>Phone: (212) 555-1234</p>
                <p style={cardP}>Email: mr.whiskaz@catnap.meow</p>
            </div>

            <div className="contact-card">
                <img style={cardImg} src="http://placekitten.com/400/200" />
                <h3 style={cardH3}>Fluffykins</h3>
                <p style={cardP}>Phone: (212) 555-2345</p>
                <p style={cardP}>Email: fluff@me.com</p>
            </div>

            <div className="contact-card">
                <img style={cardImg} src="http://placekitten.com/400/300" />
                <h3 style={cardH3}>Destroyer</h3>
                <p style={cardP}>Phone: (212) 555-3456</p>
                <p style={cardP}>Email: ofworlds@yahoo.com</p>
            </div>

            <div className="contact-card">
                <img style={cardImg} src="http://placekitten.com/200/100" />
                <h3 style={cardH3}>Felix</h3>
                <p style={cardP}>Phone: (212) 555-4567</p>
                <p style={cardP}>Email: thecat@hotmail.com</p>
            </div>
        </div>
    )
}

export default Contacts