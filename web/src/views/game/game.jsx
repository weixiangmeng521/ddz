import React from "react"
import { Socket } from "../../socket/socket";
import ReadyState from "./state/ready"
import Board from "./state/board"
import Card from "../../components/card"
import { sortDealCards } from "../../pocker/sort"

// TODO: !!! 加心跳
class Game extends React.Component {
    constructor(props) {
        super(props);
        this.props = props
        this.state = {
            status: false,
            isReady: false,
            usersList: {},
            state: 0, // 游戏进行的场景 0：开局， 1：叫地主， 2：斗地主 3：结束游戏
            // cardsData: cardsData,
            cardsData: {},
            
            mode: "", // 当前是叫地主，还是出牌
            playerButtons: {},
            // 选中的卡牌
            selectedCards: [],
        }
    }


    componentDidMount(){
        const name = this.props.match.params.name;
        this.hasAllowed(name)
        this.reciverGameStatus()

    }

    hasAllowed = async (name) => {
        let evt = "name:set"
        Socket.emit(evt, {data: ""})
        await new Promise(r => Socket.on(evt, res => r(res)));
        Socket.off(evt)

        evt = "game:join"
        Socket.emit(evt, {data: name})
        const res = await new Promise(r => Socket.on(evt, res => r(res)));
        if(res.code === 1){
            this.setState({ status: true })
        }
        // 订阅卡牌
        this.reciverGameCards()
        this.reciverGameOptions()
    }

    // 等待游戏
    handleWaitGame = () => {
        let evt = "game:wait"
        Socket.emit(evt, {data: ""})
        this.setState({ isReady: false })
    }

    // 准备就绪
    handleReadyGame = () => {
        let evt = "game:ready"
        Socket.emit(evt, {data: ""})
        this.setState({ isReady: true });
    }

    // 获取游戏的准备信息
    reciverGameStatus = () => {
        const cb = (res) => {
            if(res.code !== 1)return
            this.setState({ usersList: res.data });
        }
        ["game:ready", "game:wait", "game:join"].map((v, i) => Socket.on(v, cb))
    }

    // 开始游戏后获取牌
    reciverGameCards = () => {
        const evt = "cards:changed";
        Socket.emit(evt, {})
        Socket.on(evt, res => {
            // debugger
            this.setState({ 
                state: 1, 
                cardsData: res.data,
            })
            // 清除之前的绑定
            const evts = ["game:ready", "game:wait", "game:join"];
            evts.map(v => Socket.off(v));       
            // 清空选中的牌
            this.setSelectedCards && this.setSelectedCards([]);
        })
    }

    // [订阅] 出牌权：选择
    reciverGameOptions = () => {
        const evt = "game:options";
        Socket.emit(evt, {})
        Socket.on(evt, async res => {
            if(this.state.mode === "call" && res.type === "play" && res.options.cannot_afford === "0"){
                this.setState({ playerButtons: { cannot_afford: "0", play_cards: "1" } })
                return
            }
            await new Promise(r => this.setState({ mode: res.type, playerButtons: res.options}, () => r()));
            res.type === "play" && this.confirmOptions()
        });
    }

    confirmOptions = async () => {
        const evt = "game:options[confirm]";
        Socket.emit(evt, {});
        const res = await new Promise(r => Socket.on(evt, async res => r(res)));
        this.setState({ playerButtons: res.options })
        Socket.off(evt);
    }


    // 命名转化
    convertStr = (s) => {
        return s.replace("_", " ");
    }


    // 叫地主，不叫
    callLoard = async (val) => {
        const evt = "game:deal";
        Socket.emit(evt, {
            data: {
                type: "call",
                option: val,
            }
        });
        const res = await new Promise(r => Socket.on(evt, res => r(res)));
        if(res.code !== 1){
            console.log(1);
            return
        }
        this.setState({ playerButtons: {} })
    }

    // 出牌，不出
    dealCards = async (val) => {
        const evt = "game:deal";
        const params = {
            data: {
                type: "play",
                option: val,
                cards: this.state.selectedCards,
            }
        }
        Socket.emit(evt, params);
        const res = await new Promise(r => Socket.on(evt, res => r(res)));
        if(res.code !== 1){
            console.log(res)
            return
        }
        this.setState({ playerButtons: {}, selectedCards: []})
    }




    // 渲染叫地主，抢地主，出牌
    renderCallLord = () => {
        const { playerButtons } = this.state;
        const list = []
        // console.log(playerButtons);
        Object.keys(playerButtons).map((k, i) =>             
            list.push(<button 
                key={i}
                type="button"
                className={ `me-2 btn ${playerButtons[k] === "1" ? "btn-primary" : "btn-secondary"}` }
                onClick={
                    () => this.state.mode === "call" ? 
                    this.callLoard(playerButtons[k]) : 
                    this.dealCards(playerButtons[k])                   
                }
            >{ this.convertStr(k) }
            </button>)
        )
        return (<>{list}</>);
    }






    // 选择中的卡牌
    selectedCards = (list, setSelectedCards) => {
        this.setState({ selectedCards: list });
        this.setSelectedCards = setSelectedCards;
    }





    
    render(){
        return (<>
            { this.state.state === 0 ? 
            <ReadyState 
                status={this.state.status}
                isReady={this.state.isReady}
                usersList={this.state.usersList}
                handleReady={this.handleReadyGame} 
                handleWait={this.handleWaitGame} 
             /> : "" }

            { this.state.state === 1 ? 
            <Board 
                cards={this.state.cardsData}
                buttons={this.renderCallLord}
                mode={this.state.mode}
                onSelectedCards={ this.selectedCards }
                >
                <div className="row justify-center-center">
                    <div>
                    {
                        sortDealCards(this.state.cardsData.playedCards)
                        .map((card,i) => 
                            <Card 
                            key={ i } 
                            type={ String(card.type) } 
                            value={ String(card.value) }
                        />) 
                    }
                    </div>  
                </div>
            </Board> : "" }



        </>)
    }
}

export default Game;