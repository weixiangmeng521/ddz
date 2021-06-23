import React from "react"
import PropTypes from 'prop-types';

const getCard = (name) => {
    return require("../assets/img/cards/" + name + ".png")
}
// 扑克映射表
const cardsMap = {}

// heart 红桃
// spade 黑桃
// club 梅花
// diamond 方块
const cardsType = [
    "spade", "heart", "diamond", "club", 
];
const jackTypes = [
    "freak", "real",
]
const unorderCards = [
    "A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "Jack"
];

let counter = 1;
unorderCards.forEach((v, i) => {
    if(v === "Jack"){
        jackTypes.forEach((t, j) => {
            cardsMap[t + v] = getCard(counter)
            counter++
        })
        return
    }
    ["spade", "heart", "club", "diamond"].forEach((t, j) => {
        cardsMap[t + v] = getCard(counter)
        counter++
    })
})

const getCardSource = (value, type) => {
    if(value === "*" || type === "*"){
        return getCard("noise").default;
    }
    const k = cardsType.concat(jackTypes)[Number(type)];
    return cardsMap[k + value].default;
}


const cardStyle = {
    width: "50px",
    display: "inline-block",
    padding: "0px",
} 

const Render = ({ value, type, onClickHandle }) => {
    return (
        <>
            <img 
                src={ getCardSource(value, type) } 
                alt="cards" 
                style={cardStyle} 
                onClick={ () => {onClickHandle && onClickHandle({value, type})} } 
            />
        </>
    )
}

Render.propTypes = {
    value: PropTypes.string,
    type: PropTypes.string,
    onClickHandle: PropTypes.func,
};
export default Render