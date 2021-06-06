import React from "react"
import { Socket } from "../../socket/socket";
import Room from "../../components/rooms"


class Home extends React.Component {
    constructor(props) {
        super(props);
        this.state = { name: 'kenny' };
    }

    // 点击设置用户名
    hanldeClick = async(e) => {
        e.preventDefault();
        const json = {
            code: 1,
            message: "succes",
            data: this.state.name,
        }
        Socket.emit("name:set", json)
        const res = await new Promise(r => Socket.on("name:set", res => r(res)));
        console.log(res);

        Socket.off("name:set")
    }



    onChange = event => {
        this.setState({name: event.target.value});
    }

    render(){
        return (<>
            <form>
                <input value={this.state.name} onChange={ this.onChange }/>
                <input type="submit" onClick={ this.hanldeClick } value="set name"/>
            </form>

            <Room />
        </>)
    }
}

export default Home;