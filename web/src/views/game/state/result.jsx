import React  from "react"
import PropTypes from 'prop-types';

const Render = ({ data }) => { 
    return (
        <div className="container">
            <table className="table">
                <tbody>
                    {
                        Object
                        .keys(data)
                        .map((k, i) => 
                        <tr key={i} className={ data[k] ? "table-warning" : "" }>
                            <td>{k}</td>
                            <td>{data[k]}</td>
                        </tr>)
                    }
                </tbody>
            </table>
        </div>
    )
}
Render.propTypes = {
    data: PropTypes.object,
}
export default Render