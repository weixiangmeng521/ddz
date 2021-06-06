import React from "react"
import { Socket } from "../socket/socket";
import { Link } from "react-router-dom"

class Rooms extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            list: []
        };

        this.getRoomList()
    }

    getRoomList = () => {
        const evt = "room:list";
        Socket.emit(evt, {})
        Socket.off(evt).on(evt, res => {
            this.setState({ list: res.data })
        })
    }

    render = props => {
        return(
            <>  
                {
                    this.state.list.map((v, i) => {
                        return <Link key={ i } to={{ pathname: "/game/" + v }}> { v } </Link>
                    })
                    
                }
            </>
        )
    }
}

export default Rooms