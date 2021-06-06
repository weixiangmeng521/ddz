import React from "react";
import cardsData from "../../api/cards.json";
import Card from "../../components/card"
import { Socket } from "../../socket/socket";
import { sortCards, sortDealCards } from "../../pocker/sort"

class Cards extends React.Component {
    constructor(props) {
        super(props);
        this.state = { 
            list: cardsData.myCards,
            selected: [],
        }
    }

    componentDidMount = () => {
        Socket.disconnect()
    }


    selectCards = (card) => {
        const list = this.state.list
        const selected = this.state.selected
        for(const i in list) {
            const obj = list[i];
            if(obj.value === card.value && String(obj.type) === card.type){
                let arr = list.splice(i,1)
                selected.push(...arr)
            }
        }
        sortCards(list);
        this.setState({ list: list, selected: sortDealCards(selected) })
    }

    cancelSeletect = (card) => {
        const list = this.state.list
        const selected = this.state.selected
        for(const i in selected) {
            const obj = selected[i];
            if(obj.value === card.value && String(obj.type) === card.type){
                let arr = selected.splice(i,1)
                list.push(...arr)
            }
        }
        this.setState({ list: list, selected: selected })        
    }


    render(){
        return (<>  
            <div className="container">
                <div className="row">
                    { 
                    sortCards(this.state.list)
                    .map((card,i) => 
                        <Card 
                            key={ i } 
                            type={ String(card.type) } 
                            value={ String(card.value) }
                            onClickHandle={ this.selectCards }
                        />) 
                    }
                </div>

                <p>Selected:</p>
                <div className="row">
                { 
                    sortDealCards(this.state.selected)
                    .map((card,i) => 
                        <Card 
                            key={ i } 
                            type={ String(card.type) } 
                            value={ String(card.value) }
                            onClickHandle={ this.cancelSeletect }
                        />) 
                }                    
                </div>
            </div>
        </>)
    }
}

export default Cards;