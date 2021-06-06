import React  from "react"
import PropTypes from 'prop-types';

const Render = ({ data }) => { 
    return (
        <table className="table">
            <tbody>
                {
                    Object
                    .keys(data)
                    .map((k, i) => 
                    <tr key={i} className={ data[k] === "already" ? "table-warning" : "" }>
                        <td>{k}</td>
                        <td>{data[k]}</td>
                    </tr>)
                }
            </tbody>
        </table>
    )
}
Render.propTypes = {
    data: PropTypes.object,
}
export default Render