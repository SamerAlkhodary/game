import React, { useContext, useEffect, useState } from 'react';
import UserContext from '../../contexts/userContext';
import { ListGroup, Card } from 'react-bootstrap'
import { ScrollMenu, VisibilityContext } from "react-horizontal-scrolling-menu";

import './index.css';

import { Link } from 'react-router-dom';
const shopPage = () => {
    function LeftArrow() {
        const { isFirstItemVisible, scrollPrev } = React.useContext(VisibilityContext)

        return (
            <button className="arrowButton" disabled={isFirstItemVisible} onClick={() => scrollPrev()}>
                {"<"}
            </button>
        );
    }
    function RightArrow() {
        const { isLastItemVisible, scrollNext } = React.useContext(VisibilityContext)

        return (
            <button className="arrowButton" disabled={isLastItemVisible} onClick={() => scrollNext()}>
                {">"}
            </button>
        );
    }

    const data = [{ id: "0", name: "Nike", type: "Shoes" }, { id: "1", name: "Adidas", type: "Shoes" }, { id: "2", name: "Google", type: "Shoes" }, { id: "3", name: "Nike", type: "Shoes" }, { id: "4", name: "Adidas", type: "Shoes" }, { id: "5", name: "Google", type: "Shoes" }, { id: "6", name: "Nike", type: "Shoes" }, { id: "7", name: "Adidas", type: "Shoes" }, { id: "8", name: "Google", type: "Shoes" }];
    const listItems = data.map(
        item =>

            <Card className="menuItem">
                <Card.Body>
                    {item.name}
                </Card.Body>
            </Card>
    );

    return (
        <div style={{ padding: 10 }}>
            <h2 style={{ textAlign: 'center' }}>Shoes</h2>
            <ScrollMenu LeftArrow={LeftArrow} RightArrow={RightArrow}>
                {listItems}
            </ScrollMenu>
            <h2 style={{ textAlign: 'center' }}>Shirts</h2>
            <ScrollMenu LeftArrow={LeftArrow} RightArrow={RightArrow}>
                {listItems}
            </ScrollMenu>
        </div>

    )
}

export default shopPage;