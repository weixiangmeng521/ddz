import React, {useState} from "react"
import PropTypes from 'prop-types';
import Card from "../../../components/card"
import { sortCards, sortDealCards } from "../../../pocker/sort"

const noiseCardStyle = {
    width: "15px",
    display: "inline-block",
}

const fightBoard = {
    height:"200px", 
}

const lordStyle = {
    backgroundColor: "red",
    color: "white",
}



const GenNoiseCards = (num) => {
    const cards = [];
    for (let i = 0; i < num; i++) {
        cards.push(
        <div style={noiseCardStyle} key={ i }>
            <Card type={ "*" } value={ "*" }/>
        </div>
        )
    }
    return cards;
}

// const handleSelectedCards = () => {

// }


const Render = ({ cards, children, buttons, onSelectedCards }) => {    

    // state
    const [state ,setState] = useState({
        selected: [],
    })

    // 选牌
    const selectCards = (card) => {
        // if(!mode || mode === "call")return;
        const list = cards.myCards
        const selected = state.selected
        for(const i in list) {
            const obj = list[i];
            if(obj.value === card.value && String(obj.type) === card.type){
                let arr = list.splice(i,1)
                selected.push(...arr)
            }
        }
        sortCards(list);
        setState({ list: list, selected: selected })

        onSelectedCards && onSelectedCards(state.selected, setSelectedCards)
    }

    // 设置选中的牌
    const setSelectedCards = (list) => {
        setState({ selected: list });
        onSelectedCards && onSelectedCards(state.selected, setSelectedCards)
    }


    // 取消选牌
    const cancelSeletect = (card) => {
        const list = cards.myCards
        const selected = state.selected
        for(const i in selected) {
            const obj = selected[i];
            if(obj.value === card.value && String(obj.type) === card.type){
                let arr = selected.splice(i,1)
                list.push(...arr)
            }
        }
        setState({ list: list, selected: selected })       
        
        onSelectedCards && onSelectedCards(state.selected, setSelectedCards)
    }


    // 是不是地主
    const isLord = (name) => {
        // 如果name不存在
        if(name === "" && cards.rolesMap && cards.myId){
            return cards.rolesMap[cards.myId] === 1
        }
        if(cards.rolesMap){
            return cards.rolesMap[name] === 1
        }
        return false;
    }

    return (
        <>  
            <div className="container">
                <div className="row justify-content-center">
                    <div className="col-4">
                        <p>Lordcards:</p>
                        { 
                        cards.lordCards
                        .map((card,i) => 
                            <Card key={ i } type={ String(card.type) } value={ String(card.value) } 
                        />) 
                        }
                    </div>
                </div>
                <div className="row">
                    {
                        Object
                        .keys(cards.hideCards)
                        .map((k, i) => {
                            return <div className="col-6" key={ i }>
                                <p style={ isLord(k) ? lordStyle : {} }>{ k } : { cards.hideCards[k] }</p>
                                { GenNoiseCards(cards.hideCards[k]) }
                            </div>
                        })
                    }
                </div>
                <div className="row" style={ fightBoard }>
                    { children }
                </div>
                
                <div className="row justify-content-center">
                    <div className="col-12">
                        { buttons() }
                    </div>
                </div>
                <div className="row">
                    <p><span style={ isLord("") ? lordStyle : {} }>My Cards:</span></p>
                    { 
                    sortCards(cards.myCards)
                    .map((card,i) => 
                        <Card 
                        key={ i } 
                        type={ String(card.type) } 
                        value={ String(card.value) }
                        onClickHandle={ selectCards }
                    />) 
                    }
                </div>
                <div className="row">
                    <p>my selected cards:</p>
                    { 
                    sortDealCards(state.selected)
                    .map((card,i) => 
                        <Card 
                        key={ i } 
                        type={ String(card.type) } 
                        value={ String(card.value) }
                        onClickHandle={ cancelSeletect }
                    />) 
                    }                    
                </div>
            </div>

            
            
        </>
    );
}
Render.propTypes = {
    cards: PropTypes.object,
    myId: PropTypes.string,
    rolesMap: PropTypes.object,
}

export default Render;