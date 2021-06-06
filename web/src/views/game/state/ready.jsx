import React from "react"
import PropTypes from 'prop-types';
import UserList from "../../../components/userList";



// 游戏准备阶段
const Render = (props) => {
    const r = () => {
        return (
        <div className="container">
            <UserList data={ props.usersList }></UserList>
            <p> {props.isReady ? "already" : "waiting" }</p>

            <div className="row">
                <div className="col-6">
                    <button type="button" className="me-2 btn btn-secondary" onClick={ props.handleWait }>waiting</button>
                    <button type="button" className="btn btn-primary" onClick={ props.handleReady }>ready</button>
                </div>
            </div>
        </div>);
    }
    return (
        <>{ props.status ? r() : "you cannot play this game." }</>
    )
}

Render.propTypes = {
    status: PropTypes.bool.isRequired,
    isReady: PropTypes.bool.isRequired,
    userslist: PropTypes.object,
    hanleWait: PropTypes.func,
    handleReady: PropTypes.func,
};
export default Render;
