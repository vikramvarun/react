import React from "react"
import Contacts from "./components/Contacts"
import ContactsData from "./components/ContactsData"
import Jokes from "./components/Jokes"
import JokesData from "./components/JokesData"
import Products from "./components/Products"
import ProductsData from "./components/ProductsData"

function App() {

    const contactList = {
        "display": "flex",
        "flexWrap": "wrap"
    }

    const jokeComponents = JokesData.map(joke =>
        <Jokes
            key={joke.id}
            question={joke.question}
            punchLine={joke.punchLine}
        />
    )

    const contactComponents = ContactsData.map(contact =>
        <Contacts
            key={contact.id}
            name={contact.name}
            imgUrl={contact.imgUrl}
            phone={contact.phone}
            email={contact.email}
        />
    )

    const productsCompomemts = ProductsData.map(products =>
        <Products
            key={products.id}
            name={products.name}
            price={products.price}
            description={products.description}
        />
    )

    return (
        <div>
            <div style={contactList}>
                {contactComponents}
            </div>
            <div>
                <h2 >Jokes:</h2>
                <br />
                {jokeComponents}
            </div>
            <div>
                <h2 >Jokes:</h2>
                <br />
                {productsCompomemts}
            </div>
        </div>



    )
}

export default App